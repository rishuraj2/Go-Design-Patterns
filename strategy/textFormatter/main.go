package main

import textformatservice "textformatter/textFormatService"

func main() {
	editor := textformatservice.NewTextEditor(textformatservice.NewTitleCaseFormatter())

	// [Title Case] expected output: Hello Guys!
	editor.PublishText("heLLo guys!")

	editor.SetFormatter(textformatservice.NewLowerCaseFormatter())

	// [lower case] expected output: hello guys!
	editor.PublishText("heLLo guys!")
}
