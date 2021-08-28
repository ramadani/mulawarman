package balancer

import (
	"context"
	"time"
)

// Request to update balance
type Request struct {
	ID    string
	Value float64
}

// Result from update balance
type Result struct {
	Value float64
	Time  time.Time
}

// Balancer add or deduct balance
type Balancer interface {
	Add(context.Context, *Request) (*Result, error)
	Deduct(context.Context, *Request) (*Result, error)
}
