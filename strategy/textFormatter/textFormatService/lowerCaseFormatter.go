package textformatservice

import "strings"

type LowerCaseFormatter struct{}

func NewLowerCaseFormatter() LowerCaseFormatter {
	return LowerCaseFormatter{}
}

func (this LowerCaseFormatter) Format(text string) string {
	return strings.ToLower(text)
}
