package textformatservice

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TitleCaseFormatter struct{}

func NewTitleCaseFormatter() TitleCaseFormatter {
	return TitleCaseFormatter{}
}

func (this TitleCaseFormatter) Format(text string) string {
	return cases.Title(language.English).String(text)
}
