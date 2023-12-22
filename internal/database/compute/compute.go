package compute

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

var (
	errInvalidSymbol           = errors.New("invalid symbol")
	errInvalidCommand          = errors.New("invalid command")
	errInvalidCommandArguments = errors.New("invalid command arguments")
	errInvalidParser           = errors.New("query parser is invalid")
	errInvalidAnalyzer         = errors.New("query analyzer is invalid")
)

type parser interface {
	ParseQuery(query string) (tokens []string, err error)
}

type analyzer interface {
	AnalyzeQuery(tokens []string) (query Query, err error)
}

type Compute struct {
	parser   parser
	analyzer analyzer
	logger   *zap.Logger
}

func NewCompute(parser parser, analyzer analyzer, logger *zap.Logger) (*Compute, error) {
	if parser == nil {
		return nil, errInvalidParser
	}

	if analyzer == nil {
		return nil, errInvalidAnalyzer
	}

	return &Compute{
		parser:   parser,
		analyzer: analyzer,
		logger:   logger,
	}, nil
}

func (c *Compute) HandleQuery(_ context.Context, queryStr string) (Query, error) {
	tokens, err := c.parser.ParseQuery(queryStr)
	if err != nil {
		return Query{}, err
	}

	query, err := c.analyzer.AnalyzeQuery(tokens)
	if err != nil {
		return Query{}, err
	}

	return query, nil
}
