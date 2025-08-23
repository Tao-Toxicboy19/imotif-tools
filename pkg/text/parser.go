package text

import (
	"fmt"
	"strings"
)

type Parser struct {
	args []string
}

func NewParser(args []string) *Parser {
	return &Parser{
		args: args,
	}
}

func (p *Parser) Parse(fallbackErrMsg ...string) (string, error) {
	if p.args == nil {
		return "", fmt.Errorf("args is nil")
	}

	if len(p.args) < 2 {
		return "", fmt.Errorf("args is too short")
	}

	str := strings.Join(p.args[2:], " ")
	if strings.TrimSpace(str) == "" {
		if len(fallbackErrMsg) > 0 && strings.TrimSpace(fallbackErrMsg[0]) != "" {
			return "", fmt.Errorf(fallbackErrMsg[0])
		}
		return "", fmt.Errorf("parsed input is empty")
	}

	return str, nil
}
