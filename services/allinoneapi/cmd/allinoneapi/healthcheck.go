package main

import (
	"time"
)

func NewHealthCheck() (*HealthChecker, error) {
	return &HealthChecker{}, nil
}

type HealthChecker struct {
}

func (h *HealthChecker) Run(d time.Duration, close <-chan struct{}) {
	for range time.Tick(d) {
		select {
		case <-close:
			return
		default:
		}
	}
}

func (h *HealthChecker) Healthy() bool {
	return true
}
