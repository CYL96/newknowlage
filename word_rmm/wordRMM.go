package word_rmm

import (
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	//"log"
	//"fmt"
	"os"
	"baliance.com/gooxml/measurement"
	"fmt"
	"baliance.com/gooxml/common"
	"errors"
)
func Wrodbuild(diseaseID string,odertime string,confirmtime,completetime string,qualityStandard string,owner string,ownersignPATH string,constructor string,constructorsignPATH string)error{
	os.Remove(diseaseID+".docx")
	doc := document.New()

	g := doc.AddParagraph()
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
	g.Properties().Spacing().SetLineSpacing(24,wml.ST_LineSpacingRuleExact)
	r = g.AddRun()
	r.AddTab()
	r.Properties().SetFontFamily("宋体")
	r.Properties().SetSize(12)
	r.Properties().SetBold(false)
	r.AddText("编号："+diseaseID)

	g1 := doc.AddParagraph()
	//g.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	g1.Properties().SetAlignment(wml.ST_JcUnset)
	g1.Properties().Spacing().SetLineSpacing(24,wml.ST_LineSpacingRuleExact)
	//	g1.Properties().SetSpacing(0.5,0.5)
	//	g1.Properties().SetFirstLineIndent(0.5 * measurement.Inch)
	//r.Properties().SetCharacterSpacing(1.5)
	//r.AddTab()
	r = g1.AddRun()
	r.Properties().SetFontFamily("宋体")
	r.Properties().SetSize(12)
	r.Properties().SetBold(false)
	r.AddText(constructor+"：")

	g1 = doc.AddParagraph()
	g1.Properties().Spacing().SetLineSpacing(24,wml.ST_LineSpacingRuleExact)
	//g.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	g1.Properties().SetFirstLineIndent(24)
	//g1.Properties().X().Spacing.LineRuleAttr = wml.ST_LineSpacingRuleAuto
	r = g1.AddRun()
	//r.AddTab()
	r.Properties().SetSize(12)
	//r.Properties().SetCharacterSpacing(1.5)
	//r.Properties().SetKerning(4.0)
	r.AddText("现将"+diseaseID+"号任务通知单下达给你部，请在"+completetime+"前完成，质量保准为"+qualityStandard+"，请你部严格按照养护安全作业规程做好安全工作，按时组织施工。")

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
			//row.AddCell().AddParagraph().AddRun().AddText("签发单位：四川隧唐科技股份有限公司")
			tableaddtext(row,24,wml.ST_JcUnset,24,"宋体",12,"签发单位："+owner)
			tableaddtext(row,24,wml.ST_JcUnset,24,"宋体",12,"接收单位："+constructor)
		}
		if r == 1{
			oiref := docaddimg(doc,ownersignPATH)
			ciref := docaddimg(doc,constructorsignPATH)
			if oiref == nil || ciref == nil{
				return errors.New("img open failed")
			}
			err := tableaddtextWithImg(row,*oiref,24,wml.ST_JcUnset,24,"宋体",12,"(签字盖章)")
			if err != nil {
				return err
			}
			err = tableaddtextWithImg(row,*ciref,24,wml.ST_JcUnset,24,"宋体",12,"(签字盖章)")
			if err != nil {
				return err
			}
			//tableaddtext(row,24,wml.ST_JcUnset,24,"宋体",12,"____________（签字盖章）")
			//row.AddCell().AddParagraph().AddRun().AddText("____________（签字盖章）")
			//row.AddCell().AddParagraph().AddRun().AddText("____________（签字盖章）")
		}
		if r == 2{
			tableaddtext(row,24,wml.ST_JcUnset,24,"宋体",12,odertime)
			tableaddtext(row,24,wml.ST_JcUnset,24,"宋体",12,confirmtime)
			//row.AddCell().AddParagraph().AddRun().AddText("2018年05月15日")
		}
	}
	doc.SaveToFile(diseaseID+".docx")
	return nil
}

func paragraphaddtext(p document.Paragraph,startIndent measurement.Distance,align wml.ST_Jc,LineSpacing measurement.Distance,font string,fontsize measurement.Distance,text string){
	p.Properties().SetStartIndent(0.5*measurement.Inch)
	p.Properties().SetAlignment(align)
	p.Properties().Spacing().SetLineSpacing(LineSpacing,wml.ST_LineSpacingRuleExact)
	r := p.AddRun()
	r.Properties().SetFontFamily(font)
	r.Properties().SetSize(fontsize)
	r.AddText(text)
	return
}
func tableaddtext(row document.Row,startIndent measurement.Distance,align wml.ST_Jc,LineSpacing measurement.Distance,font string,fontsize measurement.Distance,text string){
	p :=row.AddCell().AddParagraph()
	p.Properties().SetStartIndent(startIndent)
	p.Properties().SetAlignment(align)
	p.Properties().Spacing().SetLineSpacing(LineSpacing,wml.ST_LineSpacingRuleExact)
	r := p.AddRun()
	r.Properties().SetFontFamily(font)
	r.Properties().SetSize(fontsize)
	r.AddText(text)
	return
}
func tableaddtextWithImg(row document.Row,iref common.ImageRef,startIndent measurement.Distance,align wml.ST_Jc,LineSpacing measurement.Distance,font string,fontsize measurement.Distance,text string)error{
	p :=row.AddCell().AddParagraph()

	p.Properties().SetStartIndent(startIndent)
	p.Properties().SetAlignment(align)
	p.Properties().Spacing().SetLineSpacing(LineSpacing,wml.ST_LineSpacingRuleExact)
	r := p.AddRun()
	r.Properties().SetFontFamily(font)
	r.Properties().SetSize(fontsize)

	anchored,err := r.AddDrawingAnchored(iref)
	if err != nil {
		fmt.Println("unable to add anchored image: %s", err)
		return err
	}
	anchored.SetSize(78,32)
	//anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	//anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetXOffset(40)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextLargest)

	r.AddText(text)
	return nil
}

func docaddimg(doc *document.Document,imgpath string) *common.ImageRef{
	img ,err := common.ImageFromFile("123.png")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	iref, err := doc.AddImage(img)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &iref
}

