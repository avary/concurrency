package models

import (
	"time"
)

type CacheItem struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
	Node      string    `json:"node,omitempty"`
}
