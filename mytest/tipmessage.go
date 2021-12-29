package mytest

import "github.com/NiuStar/xsql2"

type Tipmessage struct {
	STATE        *xsql2.XSqlParam
	NEXTTIPTIME  *xsql2.XSqlParam
	DEL          *xsql2.XSqlParam
	ADD_UID      *xsql2.XSqlParam
	ADD_NAME     *xsql2.XSqlParam
	GET_NAME     *xsql2.XSqlParam
	MID          *xsql2.XSqlParam
	GET_UID      *xsql2.XSqlParam
	TYPE         *xsql2.XSqlParam
	STARTTIPTIME *xsql2.XSqlParam
	INTERVALTIME *xsql2.XSqlParam
	MOBILE       *xsql2.XSqlParam
	MESSAGE      *xsql2.XSqlParam
	WXNAME       *xsql2.XSqlParam
	REMINDNUM    *xsql2.XSqlParam
	UNIT         *xsql2.XSqlParam
	DISTANCETIME *xsql2.XSqlParam
}

func (table *Tipmessage) initObject() {
	table.DEL = &xsql2.XSqlParam{Name: "del", Type_: "int", Target: table, AS_: ""}
	table.ADD_UID = &xsql2.XSqlParam{Name: "add_uid", Type_: "string", Target: table, AS_: ""}
	table.ADD_NAME = &xsql2.XSqlParam{Name: "add_name", Type_: "string", Target: table, AS_: ""}
	table.GET_NAME = &xsql2.XSqlParam{Name: "get_name", Type_: "string", Target: table, AS_: ""}
	table.MID = &xsql2.XSqlParam{Name: "mid", Type_: "int", Target: table, AS_: ""}
	table.GET_UID = &xsql2.XSqlParam{Name: "get_uid", Type_: "string", Target: table, AS_: ""}
	table.TYPE = &xsql2.XSqlParam{Name: "type", Type_: "string", Target: table, AS_: ""}
	table.STARTTIPTIME = &xsql2.XSqlParam{Name: "startTipTime", Type_: "string", Target: table, AS_: ""}
	table.INTERVALTIME = &xsql2.XSqlParam{Name: "intervalTime", Type_: "string", Target: table, AS_: ""}
	table.MOBILE = &xsql2.XSqlParam{Name: "mobile", Type_: "string", Target: table, AS_: ""}
	table.MESSAGE = &xsql2.XSqlParam{Name: "message", Type_: "string", Target: table, AS_: ""}
	table.WXNAME = &xsql2.XSqlParam{Name: "wxname", Type_: "string", Target: table, AS_: ""}
	table.REMINDNUM = &xsql2.XSqlParam{Name: "remindnum", Type_: "int", Target: table, AS_: ""}
	table.UNIT = &xsql2.XSqlParam{Name: "unit", Type_: "string", Target: table, AS_: ""}
	table.DISTANCETIME = &xsql2.XSqlParam{Name: "distanceTime", Type_: "int", Target: table, AS_: ""}
	table.STATE = &xsql2.XSqlParam{Name: "state", Type_: "int", Target: table, AS_: ""}
	table.NEXTTIPTIME = &xsql2.XSqlParam{Name: "nextTipTime", Type_: "string", Target: table, AS_: ""}
}
func (table *Tipmessage) GetName() string {
	return "tipmessage"
}
func CreateTipmessage() *Tipmessage {
	table := &Tipmessage{}
	table.initObject()
	return table
}
