// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package dpredis

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	lockRedisClienterMockExpire   sync.RWMutex
	lockRedisClienterMockFlushAll sync.RWMutex
	lockRedisClienterMockGet      sync.RWMutex
	lockRedisClienterMockPing     sync.RWMutex
	lockRedisClienterMockSet      sync.RWMutex
)

// Ensure, that RedisClienterMock does implement RedisClienter.
// If this is not the case, regenerate this file with moq.
var _ RedisClienter = &RedisClienterMock{}

// RedisClienterMock is a mock implementation of RedisClienter.
//
//     func TestSomethingThatUsesRedisClienter(t *testing.T) {
//
//         // make and configure a mocked RedisClienter
//         mockedRedisClienter := &RedisClienterMock{
//             ExpireFunc: func(key string, expiration time.Duration) *redis.BoolCmd {
// 	               panic("mock out the Expire method")
//             },
//             FlushAllFunc: func() *redis.StatusCmd {
// 	               panic("mock out the FlushAll method")
//             },
//             GetFunc: func(key string) *redis.StringCmd {
// 	               panic("mock out the Get method")
//             },
//             PingFunc: func() *redis.StatusCmd {
// 	               panic("mock out the Ping method")
//             },
//             SetFunc: func(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
// 	               panic("mock out the Set method")
//             },
//         }
//
//         // use mockedRedisClienter in code that requires RedisClienter
//         // and then make assertions.
//
//     }
type RedisClienterMock struct {
	// ExpireFunc mocks the Expire method.
	ExpireFunc func(key string, expiration time.Duration) *redis.BoolCmd

	// FlushAllFunc mocks the FlushAll method.
	FlushAllFunc func() *redis.StatusCmd

	// GetFunc mocks the Get method.
	GetFunc func(key string) *redis.StringCmd

	// PingFunc mocks the Ping method.
	PingFunc func() *redis.StatusCmd

	// SetFunc mocks the Set method.
	SetFunc func(key string, value interface{}, expiration time.Duration) *redis.StatusCmd

	// calls tracks calls to the methods.
	calls struct {
		// Expire holds details about calls to the Expire method.
		Expire []struct {
			// Key is the key argument value.
			Key string
			// Expiration is the expiration argument value.
			Expiration time.Duration
		}
		// FlushAll holds details about calls to the FlushAll method.
		FlushAll []struct {
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Key is the key argument value.
			Key string
		}
		// Ping holds details about calls to the Ping method.
		Ping []struct {
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value interface{}
			// Expiration is the expiration argument value.
			Expiration time.Duration
		}
	}
}

// Expire calls ExpireFunc.
func (mock *RedisClienterMock) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	if mock.ExpireFunc == nil {
		panic("RedisClienterMock.ExpireFunc: method is nil but RedisClienter.Expire was just called")
	}
	callInfo := struct {
		Key        string
		Expiration time.Duration
	}{
		Key:        key,
		Expiration: expiration,
	}
	lockRedisClienterMockExpire.Lock()
	mock.calls.Expire = append(mock.calls.Expire, callInfo)
	lockRedisClienterMockExpire.Unlock()
	return mock.ExpireFunc(key, expiration)
}

// ExpireCalls gets all the calls that were made to Expire.
// Check the length with:
//     len(mockedRedisClienter.ExpireCalls())
func (mock *RedisClienterMock) ExpireCalls() []struct {
	Key        string
	Expiration time.Duration
} {
	var calls []struct {
		Key        string
		Expiration time.Duration
	}
	lockRedisClienterMockExpire.RLock()
	calls = mock.calls.Expire
	lockRedisClienterMockExpire.RUnlock()
	return calls
}

// FlushAll calls FlushAllFunc.
func (mock *RedisClienterMock) FlushAll() *redis.StatusCmd {
	if mock.FlushAllFunc == nil {
		panic("RedisClienterMock.FlushAllFunc: method is nil but RedisClienter.FlushAll was just called")
	}
	callInfo := struct {
	}{}
	lockRedisClienterMockFlushAll.Lock()
	mock.calls.FlushAll = append(mock.calls.FlushAll, callInfo)
	lockRedisClienterMockFlushAll.Unlock()
	return mock.FlushAllFunc()
}

// FlushAllCalls gets all the calls that were made to FlushAll.
// Check the length with:
//     len(mockedRedisClienter.FlushAllCalls())
func (mock *RedisClienterMock) FlushAllCalls() []struct {
} {
	var calls []struct {
	}
	lockRedisClienterMockFlushAll.RLock()
	calls = mock.calls.FlushAll
	lockRedisClienterMockFlushAll.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *RedisClienterMock) Get(key string) *redis.StringCmd {
	if mock.GetFunc == nil {
		panic("RedisClienterMock.GetFunc: method is nil but RedisClienter.Get was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	lockRedisClienterMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockRedisClienterMockGet.Unlock()
	return mock.GetFunc(key)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedRedisClienter.GetCalls())
func (mock *RedisClienterMock) GetCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	lockRedisClienterMockGet.RLock()
	calls = mock.calls.Get
	lockRedisClienterMockGet.RUnlock()
	return calls
}

// Ping calls PingFunc.
func (mock *RedisClienterMock) Ping() *redis.StatusCmd {
	if mock.PingFunc == nil {
		panic("RedisClienterMock.PingFunc: method is nil but RedisClienter.Ping was just called")
	}
	callInfo := struct {
	}{}
	lockRedisClienterMockPing.Lock()
	mock.calls.Ping = append(mock.calls.Ping, callInfo)
	lockRedisClienterMockPing.Unlock()
	return mock.PingFunc()
}

// PingCalls gets all the calls that were made to Ping.
// Check the length with:
//     len(mockedRedisClienter.PingCalls())
func (mock *RedisClienterMock) PingCalls() []struct {
} {
	var calls []struct {
	}
	lockRedisClienterMockPing.RLock()
	calls = mock.calls.Ping
	lockRedisClienterMockPing.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *RedisClienterMock) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if mock.SetFunc == nil {
		panic("RedisClienterMock.SetFunc: method is nil but RedisClienter.Set was just called")
	}
	callInfo := struct {
		Key        string
		Value      interface{}
		Expiration time.Duration
	}{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	lockRedisClienterMockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	lockRedisClienterMockSet.Unlock()
	return mock.SetFunc(key, value, expiration)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//     len(mockedRedisClienter.SetCalls())
func (mock *RedisClienterMock) SetCalls() []struct {
	Key        string
	Value      interface{}
	Expiration time.Duration
} {
	var calls []struct {
		Key        string
		Value      interface{}
		Expiration time.Duration
	}
	lockRedisClienterMockSet.RLock()
	calls = mock.calls.Set
	lockRedisClienterMockSet.RUnlock()
	return calls
}
