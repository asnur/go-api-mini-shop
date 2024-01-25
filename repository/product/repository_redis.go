package product

import (
	"context"
	"encoding/json"
	"go-api-mini-shop/domain"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RepositoryRedis struct {
	dbRedis *redis.Client
	Ctx     context.Context
}

func NewRepositoryRedis(ctx context.Context, client *redis.Client) *RepositoryRedis {
	return &RepositoryRedis{Ctx: ctx, dbRedis: client}
}

func (r *RepositoryRedis) GetProductByCategoryID(id int) ([]*domain.Product, error) {
	var products []*domain.Product

	// int to string
	idStr := strconv.Itoa(id)

	// Get data from redis
	data, err := r.dbRedis.Get(r.Ctx, idStr).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal data to struct
	if err := json.Unmarshal([]byte(data), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *RepositoryRedis) SetProductByCategoryID(id int, products []*domain.Product) error {
	// int to string
	idStr := strconv.Itoa(id)

	// Marshal struct to json
	data, err := json.Marshal(products)

	if err != nil {
		return err
	}

	// Get Time To Live in seconds
	ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))
	ttlDuration := time.Duration(ttl) * time.Second

	// Set data to redis
	if err := r.dbRedis.Set(r.Ctx, idStr, data, ttlDuration).Err(); err != nil {
		return err
	}

	return nil
}
