package mulawarman

import (
	"context"
	"time"
)

// UpdateBalanceRequest data request to update balance
type UpdateBalanceRequest struct {
	ID     string
	Amount float64
}

// UpdateBalanceResult return updated balance and time
type UpdateBalanceResult struct {
	Amount    float64
	Balance   float64
	UpdatedAt time.Time
}

// UpdateBalance command interface to update balance
type UpdateBalance interface {
	Do(ctx context.Context, req *UpdateBalanceRequest) (*UpdateBalanceResult, error)
}
