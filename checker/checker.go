package checker

import (
	"context"
	"time"
)

// Checker performs a configured check
type Checker interface {
	Check(context.Context) (*Result, error)
}

// Config is a common configuration for any checker
type Config struct {
	Type     string        `json:"type"`
	Interval time.Duration `json:"interval"`
}

// Result from a check
type Result struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	RTT     time.Duration `json:"rtt"`
}
