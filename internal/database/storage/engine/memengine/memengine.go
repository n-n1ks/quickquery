package memengine

import (
	"context"
)

type Engine struct {
	data map[string]string
}

func NewEngine() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}

func (e *Engine) Get(_ context.Context, key string) (string, bool) {
	value, ok := e.data[key]
	return value, ok
}

func (e *Engine) Set(_ context.Context, key, value string) {
	e.data[key] = value
}

func (e *Engine) Del(_ context.Context, key string) {
	delete(e.data, key)
}
