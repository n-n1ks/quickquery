package compute

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseQuery(query string) ([]string, error) {
	sm := newStateMachine()

	return sm.parse(query)
}

func isTrailingCharacter(ch rune) bool {
	if ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r' {
		return true
	}

	return false
}

func isCharacter(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') ||
		ch == '_' || ch == '/' || ch == '*' {
		return true
	}

	return false
}
