package mocks

import (
	"context"
	"example/main/internal/fetcher"
	"example/main/internal/model"
	"sync"
)

// Ensure, that SourcesProviderMock does implement fetcher.SourcesProvider.
// If this is not the case, regenerate this file with moq.
var _ fetcher.SourcesProvider = &SourcesProviderMock{}

// SourcesProviderMock is a mock implementation of fetcher.SourcesProvider.
//
//	func TestSomethingThatUsesSourcesProvider(t *testing.T) {
//
//		// make and configure a mocked fetcher.SourcesProvider
//		mockedSourcesProvider := &SourcesProviderMock{
//			SourcesFunc: func(ctx context.Context) ([]model.Source, error) {
//				panic("mock out the Sources method")
//			},
//		}
//
//		// use mockedSourcesProvider in code that requires fetcher.SourcesProvider
//		// and then make assertions.
//
//	}
type SourcesProviderMock struct {
	// SourcesFunc mocks the Sources method.
	SourcesFunc func(ctx context.Context) ([]model.Source, error)

	// calls tracks calls to the methods.
	calls struct {
		// Sources holds details about calls to the Sources method.
		Sources []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockSources sync.RWMutex
}

// Sources calls SourcesFunc.
func (mock *SourcesProviderMock) Sources(ctx context.Context) ([]model.Source, error) {
	if mock.SourcesFunc == nil {
		panic("SourcesProviderMock.SourcesFunc: method is nil but SourcesProvider.Sources was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockSources.Lock()
	mock.calls.Sources = append(mock.calls.Sources, callInfo)
	mock.lockSources.Unlock()
	return mock.SourcesFunc(ctx)
}

// SourcesCalls gets all the calls that were made to Sources.
// Check the length with:
//
//	len(mockedSourcesProvider.SourcesCalls())
func (mock *SourcesProviderMock) SourcesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockSources.RLock()
	calls = mock.calls.Sources
	mock.lockSources.RUnlock()
	return calls
}
