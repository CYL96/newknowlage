package word_rmm

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
	"os"
	"fmt"
)

func Init() {
	os.Remove("ouput.docx")


	doc := document.New()
	para := doc.AddParagraph()
	para.SetStyle("Title")
	para.AddRun().AddText("Hello World!")
	doc.SaveToFile("document.docx")
	para = doc.AddParagraph()
	para.SetStyle("Subtitle")
	para.AddRun().AddText("Document Subtitle")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	para.AddRun().AddText("Major Section")

	para = doc.AddParagraph()
	//左对称
	para.Properties().SetFirstLineIndent(2 * measurement.Inch)

	run := para.AddRun()
	run.AddText("A run is a string of characters with the same formatting. ")

	run = para.AddRun()
	//加粗
	run.Properties().SetBold(true)
	//字体
	run.Properties().SetFontFamily("黑体")
	//字体大小
	run.Properties().SetSize(15)
	//颜色
	run.Properties().SetColor(color.Red)
	run.AddText("Multiple runs with different formatting can exist in the same paragraph. ")

	run = para.AddRun()
	run.AddText("Adding breaks to a run will insert line breaks after the run. ")
	run.AddBreak()
	run.AddBreak()
	run = createParaRun(doc, "Runs support styling options:")

	run = createParaRun(doc, "small caps")
	run.Properties().SetSmallCaps(true)

	run = createParaRun(doc, "strike")
	run.Properties().SetStrikeThrough(true)

	run = createParaRun(doc, "double strike")
	run.Properties().SetDoubleStrikeThrough(true)

	run = createParaRun(doc, "outline")
	run.Properties().SetOutline(true)

	run = createParaRun(doc, "emboss")
	run.Properties().SetEmboss(true)

	run = createParaRun(doc, "shadow")
	run.Properties().SetShadow(true)

	run = createParaRun(doc, "imprint")
	run.Properties().SetImprint(true)

	run = createParaRun(doc, "highlighting")
	run.Properties().SetHighlight(wml.ST_HighlightColorYellow)

	run = createParaRun(doc, "underline")
	run.Properties().SetUnderline(wml.ST_UnderlineWavyDouble, color.Red)

	run = createParaRun(doc, "text effects")
	run.Properties().SetEffect(wml.ST_TextEffectAntsRed)

	nd := doc.Numbering.Definitions()[0]

	for i := 1; i < 5; i++ {
		p := doc.AddParagraph()
		p.SetNumberingLevel(i - 1)
		p.SetNumberingDefinition(nd)
		run := p.AddRun()
		run.AddText(fmt.Sprintf("Level %d", i))
	}

	doc.SaveToFile("ouput.docx")

}
func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}