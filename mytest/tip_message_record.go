package mytest

import "github.com/NiuStar/xsql2"

type Tip_message_record struct {
	STATE   *xsql2.XSqlParam
	UID     *xsql2.XSqlParam
	ID      *xsql2.XSqlParam
	MESSAGE *xsql2.XSqlParam
}

func (table *Tip_message_record) initObject() {
	table.UID = &xsql2.XSqlParam{Name: "uid", Type_: "int", Target: table, AS_: ""}
	table.ID = &xsql2.XSqlParam{Name: "id", Type_: "int", Target: table, AS_: ""}
	table.MESSAGE = &xsql2.XSqlParam{Name: "message", Type_: "string", Target: table, AS_: ""}
	table.STATE = &xsql2.XSqlParam{Name: "state", Type_: "int", Target: table, AS_: ""}
}
func (table *Tip_message_record) GetName() string {
	return "tip_message_record"
}
func CreateTip_message_record() *Tip_message_record {
	table := &Tip_message_record{}
	table.initObject()
	return table
}
