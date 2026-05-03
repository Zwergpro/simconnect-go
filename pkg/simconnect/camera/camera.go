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

// Session is the subset of client.Client methods used by this package.
type Session interface {
	Bindings() *bindings.SimConnect
	RegisterHandler(core.RecvID, func(core.Message))
	RegisterCloseHook(func())
	ChannelBuffer() int
	Context() context.Context
}

// Camera exposes camera control functions and manages camera-specific subscriptions.
type Camera struct {
	c Session

	mu               sync.Mutex
	dataWaiter       chan core.RequestResult
	statusWaiter     chan core.RequestResult
	definitionWaiter chan core.RequestResult
	lockerWaiter     chan core.RequestResult
	cameraStatusSubs []chan core.CameraStatusMessage
	cameraLockerSubs []chan core.CameraWorldLockerMessage
}

// New creates a Camera client and registers message handlers on the core client.
func New(c Session) *Camera {
	s := &Camera{c: c}
	c.RegisterHandler(core.RecvIDCameraData, func(msg core.Message) {
		m, ok := msg.(core.CameraDataMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		if w := s.dataWaiter; w != nil {
			s.dataWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		s.mu.Unlock()
	})
	c.RegisterHandler(core.RecvIDCameraStatus, func(msg core.Message) {
		m, ok := msg.(core.CameraStatusMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		if w := s.statusWaiter; w != nil {
			s.statusWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		subs := append([]chan core.CameraStatusMessage(nil), s.cameraStatusSubs...)
		s.mu.Unlock()
		for _, ch := range subs {
			s.sendStatus(ch, m)
		}
	})
	c.RegisterHandler(core.RecvIDCameraDefinitionList, func(msg core.Message) {
		m, ok := msg.(core.CameraDefinitionListMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		if w := s.definitionWaiter; w != nil {
			w <- core.RequestResult{Msg: m}
			if m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf {
				s.definitionWaiter = nil
				close(w)
			}
		}
		s.mu.Unlock()
	})
	c.RegisterHandler(core.RecvIDCameraWorldLocker, func(msg core.Message) {
		m, ok := msg.(core.CameraWorldLockerMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		if w := s.lockerWaiter; w != nil {
			s.lockerWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		subs := append([]chan core.CameraWorldLockerMessage(nil), s.cameraLockerSubs...)
		s.mu.Unlock()
		for _, ch := range subs {
			s.sendWorldLocker(ch, m)
		}
	})
	c.RegisterCloseHook(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.dataWaiter != nil {
			s.dataWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(s.dataWaiter)
			s.dataWaiter = nil
		}
		if s.statusWaiter != nil {
			s.statusWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(s.statusWaiter)
			s.statusWaiter = nil
		}
		if s.definitionWaiter != nil {
			s.definitionWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(s.definitionWaiter)
			s.definitionWaiter = nil
		}
		if s.lockerWaiter != nil {
			s.lockerWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(s.lockerWaiter)
			s.lockerWaiter = nil
		}
		for _, ch := range s.cameraStatusSubs {
			close(ch)
		}
		s.cameraStatusSubs = nil
		for _, ch := range s.cameraLockerSubs {
			close(ch)
		}
		s.cameraLockerSubs = nil
	})
	return s
}

func (s *Camera) Acquire(clientID string) error {
	return s.c.Bindings().CameraAcquire(clientID)
}

func (s *Camera) Release(cameraDefinitionName string) error {
	return s.c.Bindings().CameraRelease(cameraDefinitionName)
}

func (s *Camera) Set(data core.CameraData, mask core.CameraDataMask) error {
	return s.c.Bindings().CameraSet(core.CameraDataToBinding(data), uint32(mask))
}

// Get requests the current camera data. SimConnect does not include a
// request ID in this response, so concurrent calls return core.ErrBusy.
func (s *Camera) Get(ctx context.Context, referential core.PositionReferential) (core.CameraDataMessage, error) {
	var zero core.CameraDataMessage
	waiter, err := s.addSingletonWaiter(&s.dataWaiter, "camera get")
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().CameraGet(uint32(referential)); err != nil {
		s.removeSingletonWaiter(&s.dataWaiter)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.removeSingletonWaiter(&s.dataWaiter)
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
func (s *Camera) GetStatus(ctx context.Context) (core.CameraStatusMessage, error) {
	var zero core.CameraStatusMessage
	waiter, err := s.addSingletonWaiter(&s.statusWaiter, "camera status")
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().CameraGetStatus(); err != nil {
		s.removeSingletonWaiter(&s.statusWaiter)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.removeSingletonWaiter(&s.statusWaiter)
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

func (s *Camera) EnableFlag(flag core.CameraFlag) error {
	return s.c.Bindings().CameraEnableFlag(uint32(flag))
}

func (s *Camera) DisableFlag(flag core.CameraFlag) error {
	return s.c.Bindings().CameraDisableFlag(uint32(flag))
}

func (s *Camera) SubscribeStatus(ctx context.Context) (<-chan core.CameraStatusMessage, error) {
	if err := s.c.Bindings().SubscribeToCameraStatusUpdate(); err != nil {
		return nil, err
	}
	ch := make(chan core.CameraStatusMessage, s.c.ChannelBuffer())
	s.mu.Lock()
	s.cameraStatusSubs = append(s.cameraStatusSubs, ch)
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		s.removeStatusSub(ch)
	}()
	return ch, nil
}

func (s *Camera) UnsubscribeStatus() error {
	return s.c.Bindings().UnsubscribeToCameraStatusUpdate()
}

// EnumerateDefinitions lists camera definitions. SimConnect does not
// include a request ID in this response, so concurrent calls return core.ErrBusy.
func (s *Camera) EnumerateDefinitions(ctx context.Context) (core.CameraDefinitionListMessage, error) {
	var zero core.CameraDefinitionListMessage
	waiter, err := s.addSingletonWaiter(&s.definitionWaiter, "camera definition enumeration")
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().EnumerateCameraDefinitions(); err != nil {
		s.removeSingletonWaiter(&s.definitionWaiter)
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
			s.removeSingletonWaiter(&s.definitionWaiter)
			return zero, ctx.Err()
		}
	}
}

func (s *Camera) SetUsingDefinition(cameraDefinitionName string) error {
	return s.c.Bindings().CameraSetUsingCameraDefinition(cameraDefinitionName)
}

// RequestWorldLocker requests a world-locker camera. SimConnect does not
// include a request ID in this response, so concurrent calls return core.ErrBusy.
func (s *Camera) RequestWorldLocker(ctx context.Context, position core.XYZ, referential core.PositionReferential, objectID core.ObjectID) (core.CameraWorldLockerMessage, error) {
	var zero core.CameraWorldLockerMessage
	waiter, err := s.addSingletonWaiter(&s.lockerWaiter, "camera world locker")
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().RequestCameraWorldLocker(core.XYZToBinding(position), bindings.SIMCONNECT_POSITION_REFERENTIAL(referential), uint32(objectID)); err != nil {
		s.removeSingletonWaiter(&s.lockerWaiter)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.removeSingletonWaiter(&s.lockerWaiter)
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

func (s *Camera) DeleteWorldLocker() error {
	return s.c.Bindings().DeleteCameraWorldLocker()
}

func (s *Camera) SubscribeWorldLockerStatus(ctx context.Context) (<-chan core.CameraWorldLockerMessage, error) {
	if err := s.c.Bindings().SubscribeToCameraWorldLockerStatusUpdate(); err != nil {
		return nil, err
	}
	ch := make(chan core.CameraWorldLockerMessage, s.c.ChannelBuffer())
	s.mu.Lock()
	s.cameraLockerSubs = append(s.cameraLockerSubs, ch)
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		s.removeLockerSub(ch)
	}()
	return ch, nil
}

func (s *Camera) UnsubscribeWorldLockerStatus() error {
	return s.c.Bindings().UnsubscribeToCameraWorldLockerStatusUpdate()
}

func (s *Camera) addSingletonWaiter(slot *chan core.RequestResult, name string) (<-chan core.RequestResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if *slot != nil {
		return nil, fmt.Errorf("%w: %s", core.ErrBusy, name)
	}
	ch := make(chan core.RequestResult, 16)
	*slot = ch
	return ch, nil
}

func (s *Camera) removeSingletonWaiter(slot *chan core.RequestResult) {
	s.mu.Lock()
	if *slot != nil {
		close(*slot)
		*slot = nil
	}
	s.mu.Unlock()
}

func (s *Camera) removeStatusSub(target chan core.CameraStatusMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ch := range s.cameraStatusSubs {
		if ch == target {
			close(ch)
			s.cameraStatusSubs = append(s.cameraStatusSubs[:i], s.cameraStatusSubs[i+1:]...)
			if len(s.cameraStatusSubs) == 0 {
				_ = s.c.Bindings().UnsubscribeToCameraStatusUpdate()
			}
			return
		}
	}
}

func (s *Camera) removeLockerSub(target chan core.CameraWorldLockerMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ch := range s.cameraLockerSubs {
		if ch == target {
			close(ch)
			s.cameraLockerSubs = append(s.cameraLockerSubs[:i], s.cameraLockerSubs[i+1:]...)
			if len(s.cameraLockerSubs) == 0 {
				_ = s.c.Bindings().UnsubscribeToCameraWorldLockerStatusUpdate()
			}
			return
		}
	}
}

func (s *Camera) sendStatus(ch chan core.CameraStatusMessage, msg core.CameraStatusMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-s.c.Context().Done():
	}
}

func (s *Camera) sendWorldLocker(ch chan core.CameraWorldLockerMessage, msg core.CameraWorldLockerMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-s.c.Context().Done():
	}
}
