package mocks

import (
	"context"
	"example/main/internal/fetcher"
	"example/main/internal/model"
	"sync"
)

// Ensure, that SourceMock does implement fetcher.Source.
// If this is not the case, regenerate this file with moq.
var _ fetcher.Source = &SourceMock{}

// SourceMock is a mock implementation of fetcher.Source.
//
//	func TestSomethingThatUsesSource(t *testing.T) {
//
//		// make and configure a mocked fetcher.Source
//		mockedSource := &SourceMock{
//			FetchFunc: func(ctx context.Context) ([]model.Item, error) {
//				panic("mock out the Fetch method")
//			},
//			IDFunc: func() int64 {
//				panic("mock out the ID method")
//			},
//			NameFunc: func() string {
//				panic("mock out the Name method")
//			},
//		}
//
//		// use mockedSource in code that requires fetcher.Source
//		// and then make assertions.
//
//	}
type SourceMock struct {
	// FetchFunc mocks the Fetch method.
	FetchFunc func(ctx context.Context) ([]model.Item, error)

	// IDFunc mocks the ID method.
	IDFunc func() int64

	// NameFunc mocks the Name method.
	NameFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Fetch holds details about calls to the Fetch method.
		Fetch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
	}
	lockFetch sync.RWMutex
	lockID    sync.RWMutex
	lockName  sync.RWMutex
}

// Fetch calls FetchFunc.
func (mock *SourceMock) Fetch(ctx context.Context) ([]model.Item, error) {
	if mock.FetchFunc == nil {
		panic("SourceMock.FetchFunc: method is nil but Source.Fetch was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFetch.Lock()
	mock.calls.Fetch = append(mock.calls.Fetch, callInfo)
	mock.lockFetch.Unlock()
	return mock.FetchFunc(ctx)
}

// FetchCalls gets all the calls that were made to Fetch.
// Check the length with:
//
//	len(mockedSource.FetchCalls())
func (mock *SourceMock) FetchCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFetch.RLock()
	calls = mock.calls.Fetch
	mock.lockFetch.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *SourceMock) ID() int64 {
	if mock.IDFunc == nil {
		panic("SourceMock.IDFunc: method is nil but Source.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//
//	len(mockedSource.IDCalls())
func (mock *SourceMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *SourceMock) Name() string {
	if mock.NameFunc == nil {
		panic("SourceMock.NameFunc: method is nil but Source.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//
//	len(mockedSource.NameCalls())
func (mock *SourceMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}
