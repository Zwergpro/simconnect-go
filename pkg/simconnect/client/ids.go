//go:build windows

package client

import "sync/atomic"

type idAllocator struct {
	next atomic.Uint32
}

func newIDAllocator(start uint32) *idAllocator {
	a := &idAllocator{}
	a.next.Store(start)
	return a
}

func (a *idAllocator) Next() uint32 {
	return a.next.Add(1) - 1
}
