package mytest

import "github.com/NiuStar/xsql2"

type Tip_state_info struct {
	STATE *xsql2.XSqlParam
	NAME  *xsql2.XSqlParam
}

func (table *Tip_state_info) initObject() {
	table.STATE = &xsql2.XSqlParam{Name: "state", Type_: "int", Target: table, AS_: ""}
	table.NAME = &xsql2.XSqlParam{Name: "name", Type_: "string", Target: table, AS_: ""}
}
func (table *Tip_state_info) GetName() string {
	return "tip_state_info"
}
func CreateTip_state_info() *Tip_state_info {
	table := &Tip_state_info{}
	table.initObject()
	return table
}
