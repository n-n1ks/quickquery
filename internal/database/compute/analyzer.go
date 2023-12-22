package compute

const (
	GetCommandID = iota
	SetCommandID
	DelCommandID
	// Must be the last one.
	commandsCount
)

const (
	GetCommandRequiredArguments = 1
	SetCommandRequiredArguments = 2
	DelCommandRequiredArguments = 1
)

type Analyzer struct {
	handlers [commandsCount]func(query Query) error
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		handlers: [commandsCount]func(query Query) error{
			GetCommandID: analyzeGetCommand,
			SetCommandID: analyzeSetCommand,
			DelCommandID: analyzeDelCommand,
		},
	}
}

func (a *Analyzer) AnalyzeQuery(tokens []string) (Query, error) {
	commandID, err := toCommandID(tokens[0])
	if err != nil {
		return Query{}, err
	}

	query := NewQuery(commandID, tokens[1:])
	handler := a.handlers[commandID]
	err = handler(query)
	if err != nil {
		return Query{}, err
	}

	return query, nil
}

func toCommandID(token string) (int, error) {
	switch token {
	case "GET":
		return GetCommandID, nil
	case "SET":
		return SetCommandID, nil
	case "DEL":
		return DelCommandID, nil
	default:
		return 0, errInvalidCommand
	}
}

func analyzeGetCommand(query Query) error {
	if len(query.Arguments()) != GetCommandRequiredArguments {
		return errInvalidCommandArguments
	}

	return nil
}

func analyzeSetCommand(query Query) error {
	if len(query.Arguments()) != SetCommandRequiredArguments {
		return errInvalidCommandArguments
	}

	return nil
}

func analyzeDelCommand(query Query) error {
	if len(query.Arguments()) != DelCommandRequiredArguments {
		return errInvalidCommandArguments
	}

	return nil
}
