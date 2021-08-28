package balancer

import (
	"context"
	"time"
)

// Request .
type Request struct {
	ID    string
	Value float64
}

// Result .
type Result struct {
	Value float64
	Time  time.Time
}

type Balancer interface {
	Add(context.Context, *Request) (*Result, error)
	Deduct(context.Context, *Request) (*Result, error)
}
