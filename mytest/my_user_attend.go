package mytest

import "github.com/NiuStar/xsql2"

type My_user_attend struct {
	TIMEOUT  *xsql2.XSqlParam
	STATUS   *xsql2.XSqlParam
	DEL      *xsql2.XSqlParam
	USERNAME *xsql2.XSqlParam
	USER_ID  *xsql2.XSqlParam
	ID       *xsql2.XSqlParam
	TIMEIN   *xsql2.XSqlParam
}

func (table *My_user_attend) initObject() {
	table.USERNAME = &xsql2.XSqlParam{Name: "username", Type_: "string", Target: table, AS_: ""}
	table.USER_ID = &xsql2.XSqlParam{Name: "user_id", Type_: "int", Target: table, AS_: ""}
	table.ID = &xsql2.XSqlParam{Name: "id", Type_: "int", Target: table, AS_: ""}
	table.TIMEIN = &xsql2.XSqlParam{Name: "timein", Type_: "string", Target: table, AS_: ""}
	table.TIMEOUT = &xsql2.XSqlParam{Name: "timeout", Type_: "string", Target: table, AS_: ""}
	table.STATUS = &xsql2.XSqlParam{Name: "status", Type_: "int", Target: table, AS_: ""}
	table.DEL = &xsql2.XSqlParam{Name: "del", Type_: "int", Target: table, AS_: ""}
}
func (table *My_user_attend) GetName() string {
	return "my_user_attend"
}
func CreateMy_user_attend() *My_user_attend {
	table := &My_user_attend{}
	table.initObject()
	return table
}
