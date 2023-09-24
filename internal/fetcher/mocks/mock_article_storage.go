package mocks

import (
	"context"
	"example/main/internal/fetcher"
	"example/main/internal/model"
	"sync"
)

// Ensure, that ArticleStorageMock does implement fetcher.ArticleStorage.
// If this is not the case, regenerate this file with moq.
var _ fetcher.ArticleStorage = &ArticleStorageMock{}

// ArticleStorageMock is a mock implementation of fetcher.ArticleStorage.
//
//	func TestSomethingThatUsesArticleStorage(t *testing.T) {
//
//		// make and configure a mocked fetcher.ArticleStorage
//		mockedArticleStorage := &ArticleStorageMock{
//			StoreFunc: func(ctx context.Context, article model.Article) error {
//				panic("mock out the Store method")
//			},
//		}
//
//		// use mockedArticleStorage in code that requires fetcher.ArticleStorage
//		// and then make assertions.
//
//	}
type ArticleStorageMock struct {
	// StoreFunc mocks the Store method.
	StoreFunc func(ctx context.Context, article model.Article) error

	// calls tracks calls to the methods.
	calls struct {
		// Store holds details about calls to the Store method.
		Store []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Article is the article argument value.
			Article model.Article
		}
	}
	lockStore sync.RWMutex
}

// Store calls StoreFunc.
func (mock *ArticleStorageMock) Store(ctx context.Context, article model.Article) error {
	if mock.StoreFunc == nil {
		panic("ArticleStorageMock.StoreFunc: method is nil but ArticleStorage.Store was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Article model.Article
	}{
		Ctx:     ctx,
		Article: article,
	}
	mock.lockStore.Lock()
	mock.calls.Store = append(mock.calls.Store, callInfo)
	mock.lockStore.Unlock()
	return mock.StoreFunc(ctx, article)
}

// StoreCalls gets all the calls that were made to Store.
// Check the length with:
//
//	len(mockedArticleStorage.StoreCalls())
func (mock *ArticleStorageMock) StoreCalls() []struct {
	Ctx     context.Context
	Article model.Article
} {
	var calls []struct {
		Ctx     context.Context
		Article model.Article
	}
	mock.lockStore.RLock()
	calls = mock.calls.Store
	mock.lockStore.RUnlock()
	return calls
}
