package balancer

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type redisBalancer struct {
	client redis.UniversalClient
}

// NewRedisBalancer .
func NewRedisBalancer(client redis.UniversalClient) Balancer {
	return &redisBalancer{client: client}
}

// Add .
func (p *redisBalancer) Add(ctx context.Context, req *Request) (*Result, error) {
	tx := p.client.TxPipeline()
	valueUpdater := tx.IncrByFloat(ctx, req.ID, req.Value)
	timeUpdater := tx.Time(ctx)

	if _, err := tx.Exec(ctx); err != nil {
		return nil, err
	}

	return &Result{
		Value: valueUpdater.Val(),
		Time:  timeUpdater.Val(),
	}, nil
}

// Deduct .
func (p *redisBalancer) Deduct(ctx context.Context, req *Request) (*Result, error) {
	req.Value = 0 - req.Value

	return p.Add(ctx, req)
}
