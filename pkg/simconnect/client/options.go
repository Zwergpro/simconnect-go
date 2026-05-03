//go:build windows

package client

import (
	"time"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

type Option func(*clientConfig)

type clientConfig struct {
	pollInterval   time.Duration
	channelBuffer  int
	manualDispatch bool
	hwnd           uintptr
	eventID        uint32
	eventHandle    uintptr
	configIndex    uint32
}

func defaultClientConfig() clientConfig {
	return clientConfig{
		pollInterval:  50 * time.Millisecond,
		channelBuffer: 16,
		configIndex:   core.ConfigIndexLocal,
	}
}

func WithPollInterval(d time.Duration) Option {
	return func(cfg *clientConfig) {
		if d > 0 {
			cfg.pollInterval = d
		}
	}
}

func WithChannelBuffer(n int) Option {
	return func(cfg *clientConfig) {
		if n >= 0 {
			cfg.channelBuffer = n
		}
	}
}

func WithManualDispatch() Option {
	return func(cfg *clientConfig) {
		cfg.manualDispatch = true
	}
}

func WithWindowHandle(hwnd uintptr) Option {
	return func(cfg *clientConfig) {
		cfg.hwnd = hwnd
	}
}

func WithEventID(eventID uint32) Option {
	return func(cfg *clientConfig) {
		cfg.eventID = eventID
	}
}

func WithEventHandle(eventHandle uintptr) Option {
	return func(cfg *clientConfig) {
		cfg.eventHandle = eventHandle
	}
}

func WithConfigIndex(configIndex uint32) Option {
	return func(cfg *clientConfig) {
		cfg.configIndex = configIndex
	}
}
