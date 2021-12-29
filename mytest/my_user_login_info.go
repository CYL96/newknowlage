package mytest

import "github.com/NiuStar/xsql2"

type My_user_login_info struct {
	TEMP_PASSWORD *xsql2.XSqlParam
	ID            *xsql2.XSqlParam
	LOGIN_TIME    *xsql2.XSqlParam
}

func (table *My_user_login_info) initObject() {
	table.TEMP_PASSWORD = &xsql2.XSqlParam{Name: "temp_password", Type_: "string", Target: table, AS_: ""}
	table.ID = &xsql2.XSqlParam{Name: "ID", Type_: "int", Target: table, AS_: ""}
	table.LOGIN_TIME = &xsql2.XSqlParam{Name: "login_time", Type_: "string", Target: table, AS_: ""}
}
func (table *My_user_login_info) GetName() string {
	return "my_user_login_info"
}
func CreateMy_user_login_info() *My_user_login_info {
	table := &My_user_login_info{}
	table.initObject()
	return table
}
