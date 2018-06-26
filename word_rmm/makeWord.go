package word_rmm
import (
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	//"log"
	//"fmt"
	"os"
)
func MakeWord(){
	os.Remove("test.docx")
	doc := document.New()
	g := doc.AddParagraph()
//	g.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	g.Properties().SetAlignment(wml.ST_JcCenter)

	r := g.AddRun()
	r.AddTab()

	r.Properties().SetSize(18)
	r.Properties().SetBold(true)
	r.Properties().SetFontFamily("黑体")
	r.AddText("施工任务通知单")

	g = doc.AddParagraph()
//	g.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	g.Properties().SetAlignment(wml.ST_JcCenter)
	r = g.AddRun()
	r.AddTab()
	r.Properties().SetFontFamily("宋体")
	r.Properties().SetSize(12)
	r.Properties().SetBold(false)
	r.AddText("编号：G215.01234.20180514001")

	g1 := doc.AddParagraph()
	//g.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	g1.Properties().SetAlignment(wml.ST_JcUnset)
//	g1.Properties().SetSpacing(0.5,0.5)

//	g1.Properties().SetFirstLineIndent(0.5 * measurement.Inch)
	//r.Properties().SetCharacterSpacing(1.5)
	//r.AddTab()
	r = g1.AddRun()
	r.Properties().SetFontFamily("宋体")
	r.Properties().SetSize(12)
	r.Properties().SetBold(false)
	r.AddText("四川隧唐科技股份有限公司：")

	g1 = doc.AddParagraph()
	g1.Properties().Spacing().SetLineSpacing(18,wml.ST_LineSpacingRuleUnset)
	//g.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	g1.Properties().SetFirstLineIndent(24)
	//g1.Properties().X().Spacing.LineRuleAttr = wml.ST_LineSpacingRuleAuto
	r = g1.AddRun()
	//r.AddTab()
	r.Properties().SetSize(12)
	//r.Properties().SetCharacterSpacing(1.5)
	//r.Properties().SetKerning(4.0)
	r.AddText("现将G215.01234.20180514001号任务通知单下达给你部，请在2018年5月20号前完成，质量保准为$质量标准$，请你部严格按照养护安全作业规程做好安全工作，按时组织施工")

	//添加空白行
	doc.AddParagraph()
	doc.AddParagraph()
	doc.AddParagraph()
	doc.AddParagraph()


	table := doc.AddTable()
//	table.Properties().SetWidthPercent(100)
	table.Properties().SetAlignment(wml.ST_JcTableCenter)
	table.Properties().SetWidthAuto()
	table.Properties().SetCellSpacingAuto()
	table.Properties().SetStyle("GridTable4-Accent1")
	look := table.Properties().TableLook()
	// these have default values in the style, so we manually turn some of them off
	look.SetFirstColumn(false)
	look.SetFirstRow(true)
	look.SetLastColumn(false)
	look.SetLastRow(true)
	look.SetHorizontalBanding(true)

	for r := 0; r < 3; r++ {
		row := table.AddRow()
		if r == 0{
			row.AddCell().AddParagraph().AddRun().AddText("签发单位：四川隧唐科技股份有限公司")
			row.AddCell().AddParagraph().AddRun().AddText("接收单位：四川交投建设工程股份有限公司")
		}
		if r == 1{
			row.AddCell().AddParagraph().AddRun().AddText("____________（签字盖章）")
			row.AddCell().AddParagraph().AddRun().AddText("____________（签字盖章）")
		}
		if r == 2{
			row.AddCell().AddParagraph().AddRun().AddText("2018年05月15日")
			row.AddCell().AddParagraph().AddRun().AddText("2018年05月15日")
		}
	}
	doc.SaveToFile("test.docx")
	return

	/*
	var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

	/*
	doc := document.New()

	// Headers/footers apply to the preceding paragraphs in the document. There
	// is a section properties on the document body itself acessible via
	// BodySection().  To have multiple different headers (aside from the
	// supported even/odd/first), we need to add multiple sections.

	// First add some content
	for i := 0; i < 5; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}

	// Construct our header
	hdr := doc.AddHeader()
	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	run.AddTab()
	run.AddText("My Document Title")

	// Create a new section and apply the header
	para = doc.AddParagraph()
	section := para.Properties().AddSection(wml.ST_SectionMarkNextPage)
	section.SetHeader(hdr, wml.ST_HdrFtrDefault)

	// Add some more content
	for i := 0; i < 5; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}

	hdr = doc.AddHeader()
	para = hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run = para.AddRun()
	run.AddTab()
	run.AddText("Different Title")
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	doc.SaveToFile("header-footer-multiple.docx")
	return
*/

	/*
	doc, err := document.OpenTemplate("template.docx")
	if err != nil {
		log.Fatalf("error opening Windows Word 2016 document: %s", err)
	}

	// We can now print out all styles in the document, verifying that they
	// exist.
	for _, s := range doc.Styles.Styles() {
		fmt.Println("style", s.Name(), "has ID of", s.StyleID(), "type is", s.Type())
	}

	// And create documents setting their style to the style ID (not style name).
	para := doc.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	//para.SetStyle("Title")
	r := para.AddRun()
	r.AddTab()
	r.AddText("My Document Title")
	r.Properties().SetBold(true)
	r.Properties().SetSize(14.5)

	para = doc.AddParagraph()
	para.SetStyle("Subtitle")
	para.AddRun().AddText("Document Subtitle")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	para.AddRun().AddText("Major Section")
	para = doc.AddParagraph()
	para = doc.AddParagraph()
	for i := 0; i < 4; i++ {
		para.AddRun().AddText(lorem)
	}

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	para.AddRun().AddText("Minor Section")
	para = doc.AddParagraph()
	for i := 0; i < 4; i++ {
		para.AddRun().AddText(lorem)
	}

	// using a pre-defined table style
	table := doc.AddTable()
	table.Properties().SetWidthPercent(90)
	table.Properties().SetStyle("GridTable4-Accent1")
	look := table.Properties().TableLook()
	// these have default values in the style, so we manually turn some of them off
	look.SetFirstColumn(false)
	look.SetFirstRow(true)
	look.SetLastColumn(false)
	look.SetLastRow(true)
	look.SetHorizontalBanding(true)

	for r := 0; r < 5; r++ {
		row := table.AddRow()
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("row %d col %d", r+1, c+1))
		}
	}
	doc.SaveToFile("use-template2.docx")*/

	//doc := document.New()
	////增加段落
	//para1 := doc.AddParagraph()
	//para1.SetStyle("Title")
	//para2 := doc.AddParagraph()
	////增加片段
	//run := para1.AddRun()
	//run.AddText("Hello World\n")
	//
	//fmt.Println(para1.Style())
	//run.AddText("Text2")
	//
	//run2 := para2.AddRun()
	//run2.AddText("Text22------>")
	//doc.SaveToFile("hello.docx")

	//
	//doc, err := document.Open("./template.docx")
	//if err != nil {
	//	log.Fatalf("error opening document: %s", err)
	//}
	////doc.Paragraphs()得到包含文档所有的段落的切片
	//for i, para := range doc.Paragraphs() {
	//	fmt.Println("style---->",para.Style())
	//	fmt.Println("X---->",para.X())
	//	//run为每个段落相同格式的文字组成的片段
	//	fmt.Println("-----------第", i, "段-------------")
	//	for j, run := range para.Runs() {
	//		fmt.Print("\t-----------第", j, "格式片段-------------")
	//		fmt.Print(run.Text())
	//	}
	//	fmt.Println()
	//}
}
