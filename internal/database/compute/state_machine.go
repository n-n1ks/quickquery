package compute

import (
	"errors"
	"strings"
)

const (
	initialState = iota
	letterFoundState
	whitespaceFoundState
)

var (
	errUnknownTransition = errors.New("unknown transition")
)

type transitionAction func(ch rune)

type stateMachine struct {
	state int
	// From status -> to status -> transition action.
	transitions  map[int]map[int]transitionAction
	stringBuffer strings.Builder
	tokens       []string
}

func newStateMachine() *stateMachine {
	sm := stateMachine{
		state:        initialState,
		stringBuffer: strings.Builder{},
		tokens:       []string{},
	}

	sm.transitions = map[int]map[int]transitionAction{
		initialState: {
			letterFoundState:     sm.handleLetterFound,
			whitespaceFoundState: sm.handleWhitespaceFound,
		},
		letterFoundState: {
			letterFoundState:     sm.handleLetterFound,
			whitespaceFoundState: sm.handleWhitespaceFound,
		},
		whitespaceFoundState: {
			letterFoundState:     sm.handleLetterFound,
			whitespaceFoundState: sm.handleWhitespaceFound,
		},
	}

	return &sm
}

func (s *stateMachine) parse(query string) ([]string, error) {
	for _, ch := range query {
		var transitionTo int

		switch {
		case isCharacter(ch):
			transitionTo = letterFoundState
		case isTrailingCharacter(ch):
			transitionTo = whitespaceFoundState
		default:
			return []string{}, errInvalidSymbol
		}

		transition, err := s.getTransition(s.state, transitionTo)
		if err != nil {
			return []string{}, err
		}

		transition(ch)
	}

	// If after the last character we are still in the letterFoundState, then we need to add the last token to the list.
	s.handleWhitespaceFound(' ')

	return s.tokens, nil
}

func (s *stateMachine) getTransition(from, to int) (transitionAction, error) {
	transition, ok := s.transitions[from][to]
	if !ok {
		return nil, errUnknownTransition
	}

	return transition, nil
}

func (s *stateMachine) handleLetterFound(ch rune) {
	s.state = letterFoundState

	s.stringBuffer.WriteRune(ch)
}

func (s *stateMachine) handleWhitespaceFound(_ rune) {
	// We don't need to do anything if we are already in the whitespaceFoundState.
	if s.state == whitespaceFoundState {
		return
	}

	s.state = whitespaceFoundState

	if s.stringBuffer.Len() > 0 {
		s.tokens = append(s.tokens, s.stringBuffer.String())
		s.stringBuffer.Reset()
	}
}
