package textformatservice

import "strings"

type UpperCaseFormatter struct{}

func NewUpperCaseFormatter() UpperCaseFormatter {
	return UpperCaseFormatter{}
}

func (this UpperCaseFormatter) Format(text string) string {
	return strings.ToUpper(text)
}
