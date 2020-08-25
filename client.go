package dpredis

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ONSdigital/dp-sessions-api/session"
	"github.com/go-redis/redis"
)

var (
	ErrEmptySessionID = errors.New("session id required but was empty")
	ErrEmptySession   = errors.New("session is empty")
	ErrEmptyAddress   = errors.New("address is empty")
	ErrEmptyPassword  = errors.New("password is empty")
	ErrInvalidTTL     = errors.New("ttl should not be zero")
)

// Client - structure for the redis client
type Client struct {
	client RedisClienter
	ttl time.Duration
}

// Config - config options for the redis client
type Config struct {
	Addr     string
	Password string
	Database int
	TTL      time.Duration
}

// NewClient - returns new redis client with provided config options
func NewClient(c Config) (*Client, error) {
	if c.Addr == "" {
		return nil, ErrEmptyAddress
	}

	if c.Password == "" {
		return nil, ErrEmptyPassword
	}

	if c.TTL == 0 {
		return nil, ErrInvalidTTL
	}

	return &Client{
		client: redis.NewClient(&redis.Options{
			Addr:     c.Addr,
			Password: c.Password,
			DB:       c.Database,
		}),
		ttl: c.TTL,
	}, nil
}

// Set - add session to redis
func (c *Client) SetSession(s *session.Session) error {
	if s == nil {
		return ErrEmptySession
	}

	sJSON, err := s.MarshalJSON()
	if err != nil {
		return err
	}

	err = c.client.Set(s.ID, sJSON, c.ttl).Err()
	if err != nil {
		return fmt.Errorf("redis client.Set returned an unexpected error: %w", err)
	}

	return nil
}

// GetByID - gets a session from redis using its ID
func (c *Client) GetByID(id string) (*session.Session, error) {
	if id == "" {
		return nil, ErrEmptySessionID
	}

	msg, err := c.client.Get(id).Result()
	if err != nil {
		return nil, err
	}

	var s *session.Session

	err = json.Unmarshal([]byte(msg), &s)
	if err != nil {
		return nil, err
	}

	// Refresh TTL on access and update LastAccessed in session
	s.LastAccessed = time.Now()
	err = c.client.Expire(s.ID, c.ttl).Err()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (c *Client) DeleteAll() error {
	err := c.client.FlushAll().Err()
	if err != nil {
		return err
	}

	return nil
}

// Set - redis implementation of Set
func (c *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(key, value, expiration)
}

// Get - redis implementation of Get
func (c *Client) Get(key string) *redis.StringCmd {
	return c.client.Get(key)
}

// FlushAll - redis implementation of FlushAll
func (c *Client) FlushAll() *redis.StatusCmd {
	return c.client.FlushAll()
}

// Expire - redis implementation of Expire
func (c *Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return c.client.Expire(key, expiration)
}

// Ping - redis implementation of Ping
func (c *Client) Ping() *redis.StatusCmd {
	return c.client.Ping()
}
