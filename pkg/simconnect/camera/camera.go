//go:build windows

// Package camera implements the Camera SimConnect API category.
package camera

import (
	"context"
	"fmt"
	"sync"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	Bindings() *bindings.SimConnect
	RegisterHandler(core.RecvID, func(core.Message))
	RegisterCloseHook(func())
	ChannelBuffer() int
	Context() context.Context
}

// Camera exposes camera control functions and manages camera-specific subscriptions.
type Camera struct {
	session Session

	mu               sync.Mutex
	dataWaiter       chan core.RequestResult
	statusWaiter     chan core.RequestResult
	definitionWaiter chan core.RequestResult
	lockerWaiter     chan core.RequestResult
	cameraStatusSubs []chan core.CameraStatusMessage
	cameraLockerSubs []chan core.CameraWorldLockerMessage
}

// New creates a Camera client and registers message handlers on the core client.
func New(s Session) *Camera {
	c := &Camera{session: s}
	s.RegisterHandler(core.RecvIDCameraData, func(msg core.Message) {
		m, ok := msg.(core.CameraDataMessage)
		if !ok {
			return
		}
		c.mu.Lock()
		if w := c.dataWaiter; w != nil {
			c.dataWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		c.mu.Unlock()
	})
	s.RegisterHandler(core.RecvIDCameraStatus, func(msg core.Message) {
		m, ok := msg.(core.CameraStatusMessage)
		if !ok {
			return
		}
		c.mu.Lock()
		if w := c.statusWaiter; w != nil {
			c.statusWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		subs := append([]chan core.CameraStatusMessage(nil), c.cameraStatusSubs...)
		c.mu.Unlock()
		for _, ch := range subs {
			c.sendStatus(ch, m)
		}
	})
	s.RegisterHandler(core.RecvIDCameraDefinitionList, func(msg core.Message) {
		m, ok := msg.(core.CameraDefinitionListMessage)
		if !ok {
			return
		}
		c.mu.Lock()
		if w := c.definitionWaiter; w != nil {
			w <- core.RequestResult{Msg: m}
			if m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf {
				c.definitionWaiter = nil
				close(w)
			}
		}
		c.mu.Unlock()
	})
	s.RegisterHandler(core.RecvIDCameraWorldLocker, func(msg core.Message) {
		m, ok := msg.(core.CameraWorldLockerMessage)
		if !ok {
			return
		}
		c.mu.Lock()
		if w := c.lockerWaiter; w != nil {
			c.lockerWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		subs := append([]chan core.CameraWorldLockerMessage(nil), c.cameraLockerSubs...)
		c.mu.Unlock()
		for _, ch := range subs {
			c.sendWorldLocker(ch, m)
		}
	})
	s.RegisterCloseHook(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		if c.dataWaiter != nil {
			c.dataWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(c.dataWaiter)
			c.dataWaiter = nil
		}
		if c.statusWaiter != nil {
			c.statusWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(c.statusWaiter)
			c.statusWaiter = nil
		}
		if c.definitionWaiter != nil {
			c.definitionWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(c.definitionWaiter)
			c.definitionWaiter = nil
		}
		if c.lockerWaiter != nil {
			c.lockerWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(c.lockerWaiter)
			c.lockerWaiter = nil
		}
		for _, ch := range c.cameraStatusSubs {
			close(ch)
		}
		c.cameraStatusSubs = nil
		for _, ch := range c.cameraLockerSubs {
			close(ch)
		}
		c.cameraLockerSubs = nil
	})
	return c
}

func (c *Camera) Acquire(clientID string) error {
	return c.session.Bindings().CameraAcquire(clientID)
}

func (c *Camera) Release(cameraDefinitionName string) error {
	return c.session.Bindings().CameraRelease(cameraDefinitionName)
}

func (c *Camera) Set(data core.CameraData, mask core.CameraDataMask) error {
	return c.session.Bindings().CameraSet(core.CameraDataToBinding(data), uint32(mask))
}

// Get requests the current camera data. SimConnect does not include a
// request ID in this response, so concurrent calls return core.ErrBusy.
func (c *Camera) Get(ctx context.Context, referential core.PositionReferential) (core.CameraDataMessage, error) {
	var zero core.CameraDataMessage
	waiter, err := c.addSingletonWaiter(&c.dataWaiter, "camera get")
	if err != nil {
		return zero, err
	}
	if err := c.session.Bindings().CameraGet(uint32(referential)); err != nil {
		c.removeSingletonWaiter(&c.dataWaiter)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		c.removeSingletonWaiter(&c.dataWaiter)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.CameraDataMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected camera data, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

// GetStatus requests the current camera status. SimConnect does not
// include a request ID in this response, so concurrent calls return core.ErrBusy.
func (c *Camera) GetStatus(ctx context.Context) (core.CameraStatusMessage, error) {
	var zero core.CameraStatusMessage
	waiter, err := c.addSingletonWaiter(&c.statusWaiter, "camera status")
	if err != nil {
		return zero, err
	}
	if err := c.session.Bindings().CameraGetStatus(); err != nil {
		c.removeSingletonWaiter(&c.statusWaiter)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		c.removeSingletonWaiter(&c.statusWaiter)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.CameraStatusMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected camera status, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

func (c *Camera) EnableFlag(flag core.CameraFlag) error {
	return c.session.Bindings().CameraEnableFlag(uint32(flag))
}

func (c *Camera) DisableFlag(flag core.CameraFlag) error {
	return c.session.Bindings().CameraDisableFlag(uint32(flag))
}

func (c *Camera) SubscribeStatus(ctx context.Context) (<-chan core.CameraStatusMessage, error) {
	if err := c.session.Bindings().SubscribeToCameraStatusUpdate(); err != nil {
		return nil, err
	}
	ch := make(chan core.CameraStatusMessage, c.session.ChannelBuffer())
	c.mu.Lock()
	c.cameraStatusSubs = append(c.cameraStatusSubs, ch)
	c.mu.Unlock()

	go func() {
		<-ctx.Done()
		c.removeStatusSub(ch)
	}()
	return ch, nil
}

func (c *Camera) UnsubscribeStatus() error {
	return c.session.Bindings().UnsubscribeToCameraStatusUpdate()
}

// EnumerateDefinitions lists camera definitions. SimConnect does not
// include a request ID in this response, so concurrent calls return core.ErrBusy.
func (c *Camera) EnumerateDefinitions(ctx context.Context) (core.CameraDefinitionListMessage, error) {
	var zero core.CameraDefinitionListMessage
	waiter, err := c.addSingletonWaiter(&c.definitionWaiter, "camera definition enumeration")
	if err != nil {
		return zero, err
	}
	if err := c.session.Bindings().EnumerateCameraDefinitions(); err != nil {
		c.removeSingletonWaiter(&c.definitionWaiter)
		return zero, err
	}
	for {
		select {
		case result, ok := <-waiter:
			if !ok {
				return zero, nil
			}
			if result.Err != nil {
				return zero, result.Err
			}
			msg, ok := result.Msg.(core.CameraDefinitionListMessage)
			if !ok {
				return zero, fmt.Errorf("%w: expected camera definition list, got %T", core.ErrDecode, result.Msg)
			}
			if zero.OutOf == 0 || msg.EntryNumber < zero.EntryNumber {
				zero.FacilityListMeta = msg.FacilityListMeta
			}
			zero.Definitions = append(zero.Definitions, msg.Definitions...)
			zero.ArraySize = uint32(len(zero.Definitions))
		case <-ctx.Done():
			c.removeSingletonWaiter(&c.definitionWaiter)
			return zero, ctx.Err()
		}
	}
}

func (c *Camera) SetUsingDefinition(cameraDefinitionName string) error {
	return c.session.Bindings().CameraSetUsingCameraDefinition(cameraDefinitionName)
}

// RequestWorldLocker requests a world-locker camera. SimConnect does not
// include a request ID in this response, so concurrent calls return core.ErrBusy.
func (c *Camera) RequestWorldLocker(ctx context.Context, position core.XYZ, referential core.PositionReferential, objectID core.ObjectID) (core.CameraWorldLockerMessage, error) {
	var zero core.CameraWorldLockerMessage
	waiter, err := c.addSingletonWaiter(&c.lockerWaiter, "camera world locker")
	if err != nil {
		return zero, err
	}
	if err := c.session.Bindings().RequestCameraWorldLocker(core.XYZToBinding(position), bindings.SIMCONNECT_POSITION_REFERENTIAL(referential), uint32(objectID)); err != nil {
		c.removeSingletonWaiter(&c.lockerWaiter)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		c.removeSingletonWaiter(&c.lockerWaiter)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.CameraWorldLockerMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected camera world locker, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

func (c *Camera) DeleteWorldLocker() error {
	return c.session.Bindings().DeleteCameraWorldLocker()
}

func (c *Camera) SubscribeWorldLockerStatus(ctx context.Context) (<-chan core.CameraWorldLockerMessage, error) {
	if err := c.session.Bindings().SubscribeToCameraWorldLockerStatusUpdate(); err != nil {
		return nil, err
	}
	ch := make(chan core.CameraWorldLockerMessage, c.session.ChannelBuffer())
	c.mu.Lock()
	c.cameraLockerSubs = append(c.cameraLockerSubs, ch)
	c.mu.Unlock()

	go func() {
		<-ctx.Done()
		c.removeLockerSub(ch)
	}()
	return ch, nil
}

func (c *Camera) UnsubscribeWorldLockerStatus() error {
	return c.session.Bindings().UnsubscribeToCameraWorldLockerStatusUpdate()
}

func (c *Camera) addSingletonWaiter(slot *chan core.RequestResult, name string) (<-chan core.RequestResult, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if *slot != nil {
		return nil, fmt.Errorf("%w: %s", core.ErrBusy, name)
	}
	ch := make(chan core.RequestResult, 16)
	*slot = ch
	return ch, nil
}

func (c *Camera) removeSingletonWaiter(slot *chan core.RequestResult) {
	c.mu.Lock()
	if *slot != nil {
		close(*slot)
		*slot = nil
	}
	c.mu.Unlock()
}

func (c *Camera) removeStatusSub(target chan core.CameraStatusMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, ch := range c.cameraStatusSubs {
		if ch == target {
			close(ch)
			c.cameraStatusSubs = append(c.cameraStatusSubs[:i], c.cameraStatusSubs[i+1:]...)
			if len(c.cameraStatusSubs) == 0 {
				_ = c.session.Bindings().UnsubscribeToCameraStatusUpdate()
			}
			return
		}
	}
}

func (c *Camera) removeLockerSub(target chan core.CameraWorldLockerMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, ch := range c.cameraLockerSubs {
		if ch == target {
			close(ch)
			c.cameraLockerSubs = append(c.cameraLockerSubs[:i], c.cameraLockerSubs[i+1:]...)
			if len(c.cameraLockerSubs) == 0 {
				_ = c.session.Bindings().UnsubscribeToCameraWorldLockerStatusUpdate()
			}
			return
		}
	}
}

func (c *Camera) sendStatus(ch chan core.CameraStatusMessage, msg core.CameraStatusMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-c.session.Context().Done():
	}
}

func (c *Camera) sendWorldLocker(ch chan core.CameraWorldLockerMessage, msg core.CameraWorldLockerMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-c.session.Context().Done():
	}
}
