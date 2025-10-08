// Package cache provides a simple Redis-backed cache for Go applications.
// It handles JSON serialization and deserialization of arbitrary data types.
package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// ErrNotFound is returned when a key is not found in the cache.
var ErrNotFound = errors.New("cache: key not found")

// Cache is a Redis cache client.
type Cache struct {
	client *redis.Client
}

// New creates a new Cache instance and connects to the Redis server.
// It pings the server to ensure a valid connection has been established.
func New(ctx context.Context, address, password string, db int) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db, // use default DB
	})

	// Ping the server to check the connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Cache{client: client}, nil
}

// Set stores a value in the cache with a given key and expiration (TTL).
// The value is automatically marshaled to JSON.
func (c *Cache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	// Serialize the object to JSON
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Set the value in Redis
	return c.client.Set(ctx, key, data, ttl).Err()
}

// Get retrieves a value from the cache by its key.
// The value is automatically unmarshaled from JSON into the 'dest' parameter.
// 'dest' must be a pointer to the variable where you want to store the result.
// Returns ErrNotFound if the key does not exist.
func (c *Cache) Get(ctx context.Context, key string, dest any) error {
	// Get the value from Redis
	data, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	// Deserialize the JSON data into the destination object
	return json.Unmarshal([]byte(data), dest)
}

// Delete removes a key from the cache.
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Close gracefully closes the connection to the Redis server.
func (c *Cache) Close() error {
	return c.client.Close()
}