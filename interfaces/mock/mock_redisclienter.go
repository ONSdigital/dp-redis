// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ONSdigital/dp-redis/interfaces"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	lockRedisClienterMockPing sync.RWMutex
	lockRedisClienterMockSet  sync.RWMutex
)

// Ensure, that RedisClienterMock does implement RedisClienter.
// If this is not the case, regenerate this file with moq.
var _ interfaces.RedisClienter = &RedisClienterMock{}

// RedisClienterMock is a mock implementation of interfaces.RedisClienter.
//
//     func TestSomethingThatUsesRedisClienter(t *testing.T) {
//
//         // make and configure a mocked interfaces.RedisClienter
//         mockedRedisClienter := &RedisClienterMock{
//             PingFunc: func() *redis.StatusCmd {
// 	               panic("mock out the Ping method")
//             },
//             SetFunc: func(in1 string, in2 interface{}, in3 time.Duration) *redis.StatusCmd {
// 	               panic("mock out the Set method")
//             },
//         }
//
//         // use mockedRedisClienter in code that requires interfaces.RedisClienter
//         // and then make assertions.
//
//     }
type RedisClienterMock struct {
	// PingFunc mocks the Ping method.
	PingFunc func() *redis.StatusCmd

	// SetFunc mocks the Set method.
	SetFunc func(in1 string, in2 interface{}, in3 time.Duration) *redis.StatusCmd

	// calls tracks calls to the methods.
	calls struct {
		// Ping holds details about calls to the Ping method.
		Ping []struct {
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 interface{}
			// In3 is the in3 argument value.
			In3 time.Duration
		}
	}
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
func (mock *RedisClienterMock) Set(in1 string, in2 interface{}, in3 time.Duration) *redis.StatusCmd {
	if mock.SetFunc == nil {
		panic("RedisClienterMock.SetFunc: method is nil but RedisClienter.Set was just called")
	}
	callInfo := struct {
		In1 string
		In2 interface{}
		In3 time.Duration
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	lockRedisClienterMockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	lockRedisClienterMockSet.Unlock()
	return mock.SetFunc(in1, in2, in3)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//     len(mockedRedisClienter.SetCalls())
func (mock *RedisClienterMock) SetCalls() []struct {
	In1 string
	In2 interface{}
	In3 time.Duration
} {
	var calls []struct {
		In1 string
		In2 interface{}
		In3 time.Duration
	}
	lockRedisClienterMockSet.RLock()
	calls = mock.calls.Set
	lockRedisClienterMockSet.RUnlock()
	return calls
}
