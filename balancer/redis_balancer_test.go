package balancer_test

import (
	"context"
	"errors"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/ramadani/mulawarman/balancer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisBalancer(t *testing.T) {
	miniRedis, err := miniredis.Run()
	assert.Nil(t, err)

	ctx := context.Background()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{miniRedis.Addr()},
	})
	redisBalancer := balancer.NewRedisBalancer(redisClient)

	defer redisClient.Close()

	t.Run("Add", func(t *testing.T) {
		mockReq := &balancer.Request{
			ID:    "qwerty-1",
			Value: 1000,
		}
		mockResVal := float64(1500)

		redisClient.Set(ctx, mockReq.ID, "500", time.Minute*1)

		res, err := redisBalancer.Add(ctx, mockReq)

		assert.Nil(t, err)
		assert.Equal(t, mockResVal, res.Value)
	})

	t.Run("Deduct", func(t *testing.T) {
		mockReq := &balancer.Request{
			ID:    "qwerty-2",
			Value: 1000,
		}
		mockResVal := float64(200)

		redisClient.Set(ctx, mockReq.ID, "1200", time.Minute*1)

		res, err := redisBalancer.Deduct(ctx, mockReq)

		assert.Nil(t, err)
		assert.Equal(t, mockResVal, res.Value)
	})

	t.Run("ExecError", func(t *testing.T) {
		defer miniRedis.SetError("")

		mockInput := &balancer.Request{
			ID:    "qwerty-3",
			Value: 1000,
		}
		mockErr := errors.New("unexpected")

		miniRedis.SetError("unexpected")

		res, err := redisBalancer.Add(ctx, mockInput)

		assert.EqualError(t, err, mockErr.Error())
		assert.Nil(t, res)
	})
}
