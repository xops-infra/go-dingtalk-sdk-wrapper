package go_dingtalk_sdk_wrapper

import "time"

type TokenDetail struct {
	// Dingtalk token
	Token    string `json:"token"`
	CreateAt int64  `json:"create_at"`
	ExpireIn int64  `json:"expire_in"`
}

func (t *TokenDetail) IsExpire() bool {
	timeNow := time.Now().Unix()
	return timeNow > t.CreateAt+t.ExpireIn
}

type Cache interface {
	Set() error
	Get() (string, error)
}

type MemoryCache struct {
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

func (r *MemoryCache) Set() error {

	return nil
}

func (r *MemoryCache) Get() error {
	return nil
}
