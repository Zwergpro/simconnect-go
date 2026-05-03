//go:build windows

package client

import "github.com/Zwergpro/simconnect-go/pkg/simconnect/core"

func (s *Sim) dispatch(msg core.Message) {
	switch m := msg.(type) {
	case core.ExceptionError:
		s.dispatchException(m)
	case interface{ DispatchRequestID() uint32 }:
		s.dispatchRequest(m.DispatchRequestID(), msg)
	default:
		s.dispatchToHandlers(msg)
	}
}

func (s *Sim) dispatchException(err core.ExceptionError) {
	matched := false
	s.mu.Lock()
	if reqID, ok := s.sendToReq[err.SendID]; ok {
		delete(s.sendToReq, err.SendID)
		if waiter, ok := s.waiters[reqID]; ok {
			delete(s.waiters, reqID)
			waiter <- core.RequestResult{Err: err}
			close(waiter)
			matched = true
		}
	}
	s.mu.Unlock()
	if !matched {
		s.ReportError(err)
	}
}

func (s *Sim) dispatchRequest(requestID uint32, msg core.Message) {
	s.mu.Lock()
	if waiter, ok := s.waiters[requestID]; ok {
		delete(s.waiters, requestID)
		waiter <- core.RequestResult{Msg: msg}
		close(waiter)
	}
	handlers := append([]func(core.Message){}, s.dataSubs[requestID]...)
	s.mu.Unlock()

	for _, handler := range handlers {
		handler(msg)
	}
}

func (s *Sim) dispatchToHandlers(msg core.Message) {
	s.mu.Lock()
	handlers := append([]func(core.Message){}, s.handlers[msg.RecvID()]...)
	s.mu.Unlock()
	for _, h := range handlers {
		h(msg)
	}
}
