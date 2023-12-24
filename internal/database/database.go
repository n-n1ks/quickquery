package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"quickquery/internal/database/compute"
	"quickquery/internal/util"
)

var (
	errInvalidComputeLayer = fmt.Errorf("invalid compute layer")
	errInvalidStorageLayer = fmt.Errorf("invalid storage layer")
	errInvalidLogger       = fmt.Errorf("invalid logger")
)

type ComputerLayer interface {
	HandleQuery(ctx context.Context, queryStr string) (query compute.Query, err error)
}

type StorageLayer interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type Database struct {
	computer ComputerLayer
	storage  StorageLayer
	logger   *zap.Logger
}

func NewDatabase(computer ComputerLayer, storage StorageLayer, logger *zap.Logger) (*Database, error) {
	if computer == nil {
		return nil, errInvalidComputeLayer
	}
	if storage == nil {
		return nil, errInvalidStorageLayer
	}
	if logger == nil {
		return nil, errInvalidLogger
	}

	return &Database{
		computer: computer,
		storage:  storage,
		logger:   logger,
	}, nil
}

func (db *Database) HandleQuery(ctx context.Context, queryStr string) string {
	tx := uuid.NewString()
	ctx = context.WithValue(ctx, util.CtxTxKey, tx)

	db.logger.Debug("handling query", zap.String("tx", tx), zap.String("query", queryStr))

	query, err := db.computer.HandleQuery(ctx, queryStr)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	switch query.CommandID() {
	case compute.GetCommandID:
		return db.handleGet(ctx, query)
	case compute.SetCommandID:
		return db.handleSet(ctx, query)
	case compute.DelCommandID:
		return db.handleDel(ctx, query)
	}

	db.logger.Debug("can't handle query", zap.String("tx", tx), zap.String("query", queryStr))

	return "error: can't handle query"
}

func (db *Database) handleGet(ctx context.Context, query compute.Query) string {
	value, err := db.storage.Get(ctx, query.Arguments()[0])
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return value
}

func (db *Database) handleSet(ctx context.Context, query compute.Query) string {
	if err := db.storage.Set(ctx, query.Arguments()[0], query.Arguments()[1]); err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return "OK"
}

func (db *Database) handleDel(ctx context.Context, query compute.Query) string {
	if err := db.storage.Del(ctx, query.Arguments()[0]); err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return "OK"
}
