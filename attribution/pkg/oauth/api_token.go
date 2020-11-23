package oauth

import (
	"context"
	"sync"
	"time"
)

var (
	DefaultIntervalTime = 31 * time.Second
	TokenPrefixKey      = "marketing-api-access-token"
)

type Token struct {
	lock  *sync.RWMutex
	token string
	store store
}

func (t *Token) Get() string {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.token
}

func (t *Token) Set(token string) error {
	return t.store.Set(TokenPrefixKey, token)
}

func (t *Token) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <- ctx.Done():
				break
			default:
				_ = t.fetch()
				time.Sleep(DefaultIntervalTime)
			}
		}
	}()
}

func (t *Token) fetch() error {
	token, err := t.store.Get(TokenPrefixKey)
	if err != nil {
		return err
	}

	t.lock.Lock()
	t.token = token
	t.lock.Unlock()
	return nil
}

func NewToken(store store) *Token {
	return &Token{
		lock: &sync.RWMutex{},
		store: store,
	}
}