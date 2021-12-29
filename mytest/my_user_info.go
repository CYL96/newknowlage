package mytest

import "github.com/NiuStar/xsql2"

type My_user_info struct {
	USERNAME *xsql2.XSqlParam
	ID       *xsql2.XSqlParam
	POSITION *xsql2.XSqlParam
}

func (table *My_user_info) initObject() {
	table.USERNAME = &xsql2.XSqlParam{Name: "username", Type_: "string", Target: table, AS_: ""}
	table.ID = &xsql2.XSqlParam{Name: "id", Type_: "int", Target: table, AS_: ""}
	table.POSITION = &xsql2.XSqlParam{Name: "position", Type_: "int", Target: table, AS_: ""}
}
func (table *My_user_info) GetName() string {
	return "my_user_info"
}
func CreateMy_user_info() *My_user_info {
	table := &My_user_info{}
	table.initObject()
	return table
}
