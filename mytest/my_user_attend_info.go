package mytest

import "github.com/NiuStar/xsql2"

type My_user_attend_info struct {
	CONSISTEND *xsql2.XSqlParam
	USERNAME   *xsql2.XSqlParam
	ID         *xsql2.XSqlParam
	USER_ID    *xsql2.XSqlParam
	PASSWORD   *xsql2.XSqlParam
	FACE       *xsql2.XSqlParam
	FP         *xsql2.XSqlParam
	IDCARD     *xsql2.XSqlParam
}

func (table *My_user_attend_info) initObject() {
	table.ID = &xsql2.XSqlParam{Name: "id", Type_: "int", Target: table, AS_: ""}
	table.USER_ID = &xsql2.XSqlParam{Name: "user_id", Type_: "int", Target: table, AS_: ""}
	table.PASSWORD = &xsql2.XSqlParam{Name: "password", Type_: "string", Target: table, AS_: ""}
	table.FACE = &xsql2.XSqlParam{Name: "face", Type_: "string", Target: table, AS_: ""}
	table.FP = &xsql2.XSqlParam{Name: "fp", Type_: "string", Target: table, AS_: ""}
	table.IDCARD = &xsql2.XSqlParam{Name: "idcard", Type_: "string", Target: table, AS_: ""}
	table.CONSISTEND = &xsql2.XSqlParam{Name: "consistend", Type_: "int", Target: table, AS_: ""}
	table.USERNAME = &xsql2.XSqlParam{Name: "username", Type_: "string", Target: table, AS_: ""}
}
func (table *My_user_attend_info) GetName() string {
	return "my_user_attend_info"
}
func CreateMy_user_attend_info() *My_user_attend_info {
	table := &My_user_attend_info{}
	table.initObject()
	return table
}
