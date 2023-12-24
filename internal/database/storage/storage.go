package storage

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"quickquery/internal/database/storage/engine"
	"quickquery/internal/util"
)

var (
	errInvalidEngine = errors.New("invalid engine")
	errKeyNotFound   = errors.New("key not found")
	errInvalidLogger = errors.New("invalid logger")
)

type Storage struct {
	engine engine.Engine
	logger *zap.Logger
}

func NewStorage(engine engine.Engine, logger *zap.Logger) (*Storage, error) {
	if engine == nil {
		return nil, errInvalidEngine
	}
	if logger == nil {
		return nil, errInvalidLogger
	}

	return &Storage{
		engine: engine,
		logger: logger,
	}, nil
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	if ctx.Err() != nil {
		tx, _ := ctx.Value(util.CtxTxKey).(string)
		s.logger.Debug("context is canceled", zap.String("tx", tx))
	}

	s.engine.Set(ctx, key, value)

	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	if ctx.Err() != nil {
		tx, _ := ctx.Value(util.CtxTxKey).(string)
		s.logger.Debug("context is canceled", zap.String("tx", tx))
	}

	value, ok := s.engine.Get(ctx, key)
	if !ok {
		return "", errKeyNotFound
	}

	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	if ctx.Err() != nil {
		tx, _ := ctx.Value(util.CtxTxKey).(string)
		s.logger.Debug("context is canceled", zap.String("tx", tx))
	}

	s.engine.Del(ctx, key)

	return nil
}
