package textformatservice

import "fmt"

type TextFormatter interface {
	Format(text string) string
}

type TextEditor struct {
	formatter TextFormatter
}

func NewTextEditor(formatter TextFormatter) TextEditor {
	return TextEditor{
		formatter: formatter,
	}
}

func (this *TextEditor) SetFormatter(formatter TextFormatter) {
	this.formatter = formatter
}

func (this *TextEditor) PublishText(text string) {
	fmt.Println(this.formatter.Format(text))
}
