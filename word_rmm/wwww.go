package word_rmm

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/schema/soo/wml"
	"errors"
	"fmt"
	"nqc.cn/xsql"
	"os"
	"time"
)
var xs *xsql.XSql
func SetXSql() {
	xs = xsql.InitSql("root", "Suitang@20170601", "192.168.1.133", "3306","rmm_test")
}
func ReportWord(diseaseID string) error {
	s := xsql.CreateInstance(xs)
	s.Qurey("select A.EngineerID,A.ConstructorID,A.SupervisiorID,C.RoadName,A.ID as DiseaseID,D.OrganizationName AS ConstructionOrganization,E.OrganizationName AS SupervisionOrganization,A.TekDemand,A.SafeDemand,DiseaseSerial, Cause,Dealing,Deadline,StartTime,StructureName,StructureID " +
		"from (table_diseases_info as A LEFT JOIN table_construction_organizations as D ON A.ConstructionOrganizationID =D.ID)LEFT JOIN table_supervision_organizations AS E ON A.SupervisionOrganizationID = E.ID,table_structure as B,table_roads_info AS C  " +
		"where A.StructureID = B.ID AND C.ID = B.RoadID and  A.ID =" + diseaseID)
	j1 := make(map[string]string)
	j1["EngineerID"] = "string"
	j1["ConstructorID"] = "string"
	j1["RoadName"] = "string"
	j1["ConstructionOrganization"] = "string"
	j1["SupervisionOrganization"] = "string"
	j1["TekDemand"] = "string"
	j1["SafeDemand"] = "string"
	j1["DiseaseSerial"] = "string"
	j1["Cause"] = "string"
	j1["Dealing"] = "string"
	j1["Deadline"] = "int64"
	j1["StartTime"] = "int64"
	j1["StructureName"] = "string"
	s.SetTableColType(j1)
	l1 := s.Execute()

	s.Qurey("SELECT C.CompanyName,A.Signature AS engineerSign,B1.Signature AS ConstructorSign,B2.Signature AS supervisorSign " +
		"FROM (table_users AS A LEFT JOIN table_users AS B1 ON B1.ID =1)LEFT JOIN table_users AS B2 ON B2.ID =1,table_company AS C  " +
		"WHERE A.ID=7 AND A.CompanyID=C.ID")
	j2 := make(map[string]string)
	j2["CompanyName"] = "string"
	j2["engineerSign"] = "string"
	j2["ConstructorSign"] = "string"
	j2["supervisorSign"] = "string"
	s.SetTableColType(j2)
	l2 := s.Execute()
	StartTime := time.Unix(l1[0]["StartTime"].(int64), 0).Format("2006年01月02日")
	EndTime := time.Unix(l1[0]["Deadline"].(int64), 0).Format("2006年01月02日")
//	fmt.Println(l2)

	os.Remove("./file/MissionFile/"+l1[0]["DiseaseSerial"].(string))
	doc := document.New()

	g := doc.AddParagraph()
	g.Properties().SetAlignment(wml.ST_JcCenter)
	r := g.AddRun()
	r.AddTab()
	r.Properties().SetSize(18)
	r.Properties().SetBold(true)
	r.Properties().SetFontFamily("黑体")
	r.AddText("质检报告")
	doc.AddParagraph()

	oiref := docaddimg(doc, l2[0]["engineerSign"].(string))
	ciref := docaddimg(doc, l2[0]["ConstructorSign"].(string))
	siref := docaddimg(doc, l2[0]["supervisorSign"].(string))

	if oiref == nil || ciref == nil || siref == nil {
		return errors.New("img open failed.")
	}

	table := doc.AddTable()
	table.Properties().SetWidth(420)
	table.Properties().SetAlignment(wml.ST_JcTableCenter)

	doc.AddParagraph()
	// 4 inches wide
	borders := table.Properties().Borders()
	// thin borders
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	//mytableaddtext(row,wml.ST_JcUnset,"`````````qqqqqqqqqqqqqqqqqqqqqqqqqq")
	//mytableaddtext(row,wml.ST_JcUnset,"`33333")

	row := table.AddRow()
	row.Properties().SetHeight(24, wml.ST_HeightRuleExact)
	mytableaddtext(row, 2, wml.ST_JcCenter, "病害ID")
	mytableaddtext(row, 2, wml.ST_JcCenter, l1[0]["DiseaseSerial"].(string))

	row = table.AddRow()
	row.Properties().SetHeight(24, wml.ST_HeightRuleExact)
	mytableaddtext(row, 2, wml.ST_JcCenter, "业主方")
	mytableaddtext(row, 2, wml.ST_JcCenter, l2[0]["CompanyName"].(string))

	row = table.AddRow()
	row.Properties().SetHeight(24, wml.ST_HeightRuleExact)
	mytableaddtext(row, 1, wml.ST_JcCenter, "监理方")
	mytableaddtext(row, 1, wml.ST_JcCenter, l1[0]["SupervisionOrganization"].(string))
	mytableaddtext(row, 1, wml.ST_JcCenter, "施工方")
	mytableaddtext(row, 1, wml.ST_JcCenter, l1[0]["ConstructionOrganization"].(string))

	row = table.AddRow()
	row.Properties().SetHeight(24, wml.ST_HeightRuleExact)
	mytableaddtext(row, 1, wml.ST_JcCenter, "道路")
	mytableaddtext(row, 1, wml.ST_JcCenter, l1[0]["RoadName"].(string))
	mytableaddtext(row, 1, wml.ST_JcCenter, "位置")
	mytableaddtext(row, 1, wml.ST_JcCenter, l1[0]["StructureName"].(string))

	row = table.AddRow()
	row.Properties().SetHeight(24, wml.ST_HeightRuleExact)
	mytableaddtext(row, 1, wml.ST_JcCenter, "病害原因")
	mytableaddtext(row, 1, wml.ST_JcCenter, l1[0]["Cause"].(string))
	mytableaddtext(row, 1, wml.ST_JcCenter, "处理方法")
	mytableaddtext(row, 1, wml.ST_JcCenter, l1[0]["Dealing"].(string))

	row = table.AddRow()
	row.Properties().SetHeight(24, wml.ST_HeightRuleExact)
	mytableaddtext(row, 1, wml.ST_JcCenter, "开始时间")
	mytableaddtext(row, 1, wml.ST_JcCenter, StartTime)
	mytableaddtext(row, 1, wml.ST_JcCenter, "结束时间")
	mytableaddtext(row, 1, wml.ST_JcCenter, EndTime)

	row = table.AddRow()
	row.Properties().SetHeight(240, wml.ST_HeightRuleExact)
	mytableaddtext(row, 1, wml.ST_JcCenter, "结论")
	mytableaddtext(row, 3, wml.ST_JcUnset, l1[0]["TekDemand"].(string)+"。"+l1[0]["SafeDemand"].(string)+"。")

	row = table.AddRow()
	row.Properties().SetHeight(80, wml.ST_HeightRuleExact)
	mytableaddtext(row, 1, wml.ST_JcCenter, "备注")
	mytableaddtext(row, 3, wml.ST_JcUnset, "")

	table = doc.AddTable()
	table.Properties().SetWidth(420)

	row = table.AddRow()
	mytableaddtextWithImg(row, *oiref, wml.ST_JcUnset, "业主方：")
	mytableaddtextWithImg(row, *siref, wml.ST_JcUnset, "监理方：")
	mytableaddtextWithImg(row, *ciref, wml.ST_JcUnset, "施工方：")

	// column span / merged cells
	//cell.Properties().SetColumnSpan(2)
	//
	//run := cell.AddParagraph().AddRun()
	//run.AddText("Cells can span multiple columns")
	//
	//row = table.AddRow()
	//cell = row.AddCell()
	//cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	//cell.AddParagraph().AddRun().AddText("Vertical Merge")
	//row.AddCell().AddParagraph().AddRun().AddText("")
	//
	//row = table.AddRow()
	//cell = row.AddCell()
	//cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	//cell.AddParagraph().AddRun().AddText("Vertical Merge 2")
	//row.AddCell().AddParagraph().AddRun().AddText("")
	//
	//row = table.AddRow()
	//row.AddCell().AddParagraph().AddRun().AddText("Street Address")
	//row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")
	doc.SaveToFile("./file/MissionFile/"+l1[0]["DiseaseSerial"].(string))
	return nil
}
func mytableaddtext(row document.Row, columnspan int, align wml.ST_Jc, text string) {
	//	c.Properties().SetWidthPercent(100)
	//	c.Properties().SetWidth(480)
	c := row.AddCell()
	n := float64(columnspan) * 25
	c.Properties().SetWidthPercent(n)
	c.Properties().SetVerticalMerge(wml.ST_MergeUnset)
	c.Properties().SetColumnSpan(columnspan)
	p := c.AddParagraph()
	//	p.Properties().SetStartIndent(24)
	p.Properties().SetAlignment(align)
	p.Properties().Spacing().SetLineSpacing(18, wml.ST_LineSpacingRuleExact)
	r := p.AddRun()
	r.Properties().SetFontFamily("宋体")
	r.Properties().SetSize(12)
	r.AddText(text)
	return
}
func mytableaddtextWithImg(row document.Row, iref common.ImageRef, align wml.ST_Jc, text string) error {
	c := row.AddCell()
	//	c.Properties().SetWidthPercent(100)
	c.Properties().SetWidth(480)
	p := c.AddParagraph()
	//p.Properties().SetStartIndent(24)
	p.Properties().SetAlignment(align)
	p.Properties().Spacing().SetLineSpacing(18, wml.ST_LineSpacingRuleExact)
	r := p.AddRun()
	r.Properties().SetFontFamily("宋体")
	r.Properties().SetSize(12)
	r.AddText(text)
	anchored, err := r.AddDrawingAnchored(iref)
	if err != nil {
		fmt.Println("unable to add anchored image: %s", err)
		return err
	}
	anchored.SetSize(78, 32)
	//anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	//anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetXOffset(60)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextLargest)
	return nil
}

func mydocaddimg(doc *document.Document, imgpath string) *common.ImageRef {
	img, err := common.ImageFromFile("." + imgpath)
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
