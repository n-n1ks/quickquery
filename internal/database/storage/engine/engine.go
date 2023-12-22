package engine

import (
	"context"
)

type Engine interface {
	Get(context.Context, string) (string, bool)
	Set(context.Context, string, string)
	Del(context.Context, string)
}
