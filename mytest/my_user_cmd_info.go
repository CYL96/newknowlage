package mytest

import "github.com/NiuStar/xsql2"

type My_user_cmd_info struct {
	CMD_RESULT               *xsql2.XSqlParam
	TRANS_ID                 *xsql2.XSqlParam
	TRANS_STATUS             *xsql2.XSqlParam
	DEVICE_ID                *xsql2.XSqlParam
	CMD_CODE                 *xsql2.XSqlParam
	TRANS_STATUS_UPDATE_TIME *xsql2.XSqlParam
	CMD_PARAM                *xsql2.XSqlParam
}

func (table *My_user_cmd_info) initObject() {
	table.TRANS_STATUS_UPDATE_TIME = &xsql2.XSqlParam{Name: "trans_status_update_time", Type_: "string", Target: table, AS_: ""}
	table.CMD_PARAM = &xsql2.XSqlParam{Name: "cmd_param", Type_: "string", Target: table, AS_: ""}
	table.CMD_RESULT = &xsql2.XSqlParam{Name: "cmd_result", Type_: "string", Target: table, AS_: ""}
	table.TRANS_ID = &xsql2.XSqlParam{Name: "trans_id", Type_: "int", Target: table, AS_: ""}
	table.TRANS_STATUS = &xsql2.XSqlParam{Name: "trans_status", Type_: "int", Target: table, AS_: ""}
	table.DEVICE_ID = &xsql2.XSqlParam{Name: "device_id", Type_: "int", Target: table, AS_: ""}
	table.CMD_CODE = &xsql2.XSqlParam{Name: "cmd_code", Type_: "string", Target: table, AS_: ""}
}
func (table *My_user_cmd_info) GetName() string {
	return "my_user_cmd_info"
}
func CreateMy_user_cmd_info() *My_user_cmd_info {
	table := &My_user_cmd_info{}
	table.initObject()
	return table
}
