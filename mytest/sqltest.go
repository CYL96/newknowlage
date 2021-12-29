package mytest

import "github.com/NiuStar/xsql2"

type Sqltest struct {
	IDDD   *xsql2.XSqlParam
	IDD    *xsql2.XSqlParam
	IDDDD  *xsql2.XSqlParam
	IDDDDD *xsql2.XSqlParam
	ID     *xsql2.XSqlParam
}

func (table *Sqltest) initObject() {
	table.IDD = &xsql2.XSqlParam{Name: "IDD", Type_: "int", Target: table, AS_: ""}
	table.IDDDD = &xsql2.XSqlParam{Name: "IDDDD", Type_: "float", Target: table, AS_: ""}
	table.IDDDDD = &xsql2.XSqlParam{Name: "IDDDDD", Type_: "string", Target: table, AS_: ""}
	table.ID = &xsql2.XSqlParam{Name: "ID", Type_: "int", Target: table, AS_: ""}
	table.IDDD = &xsql2.XSqlParam{Name: "IDDD", Type_: "string", Target: table, AS_: ""}
}
func (table *Sqltest) GetName() string {
	return "sqltest"
}
func CreateSqltest() *Sqltest {
	table := &Sqltest{}
	table.initObject()
	return table
}
