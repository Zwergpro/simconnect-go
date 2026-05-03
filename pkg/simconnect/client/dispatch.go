//go:build windows

package client

import "github.com/Zwergpro/simconnect-go/pkg/simconnect/core"

func (c *Sim) dispatch(msg core.Message) {
	switch m := msg.(type) {
	case core.ExceptionError:
		c.dispatchException(m)
	case interface{ DispatchRequestID() uint32 }:
		c.dispatchRequest(m.DispatchRequestID(), msg)
	default:
		c.dispatchToHandlers(msg)
	}
}

func (c *Sim) dispatchException(err core.ExceptionError) {
	matched := false
	c.mu.Lock()
	if reqID, ok := c.sendToReq[err.SendID]; ok {
		delete(c.sendToReq, err.SendID)
		if waiter, ok := c.waiters[reqID]; ok {
			delete(c.waiters, reqID)
			waiter <- core.RequestResult{Err: err}
			close(waiter)
			matched = true
		}
	}
	c.mu.Unlock()
	if !matched {
		c.ReportError(err)
	}
}

func (c *Sim) dispatchRequest(requestID uint32, msg core.Message) {
	c.mu.Lock()
	if waiter, ok := c.waiters[requestID]; ok {
		delete(c.waiters, requestID)
		waiter <- core.RequestResult{Msg: msg}
		close(waiter)
	}
	handlers := append([]func(core.Message){}, c.dataSubs[requestID]...)
	c.mu.Unlock()

	for _, handler := range handlers {
		handler(msg)
	}
}

func (c *Sim) dispatchToHandlers(msg core.Message) {
	c.mu.Lock()
	handlers := append([]func(core.Message){}, c.handlers[msg.RecvID()]...)
	c.mu.Unlock()
	for _, h := range handlers {
		h(msg)
	}
}
