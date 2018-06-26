package word_rmm

import (
	"baliance.com/gooxml/document"
	"fmt"
)

func ExampleOpen() {
	doc, err := document.Open("123.docx")
	if err != nil {
		fmt.Println("error opening document: %s", err)
		return
	}
	for _, para := range doc.Paragraphs() {
		fmt.Println("--",para.Properties().Style())
		for _, run := range para.Runs() {
			fmt.Print(run.Text())
			fmt.Println("--",run.Properties().Fonts(),run.Properties().Fonts(),run.Properties().IsItalic())
		}
		fmt.Println()
	}
}