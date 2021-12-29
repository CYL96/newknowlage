package mytest

import "github.com/NiuStar/xsql2"

type T_mbo_account struct {
	LEDGER_TYPE    *xsql2.XSqlParam
	LEDGER_DURATIN *xsql2.XSqlParam
	LEDGER_END     *xsql2.XSqlParam
	TITLE          *xsql2.XSqlParam
	ACCOUNT        *xsql2.XSqlParam
	TRAIN_DURATION *xsql2.XSqlParam
	SCHOOL         *xsql2.XSqlParam
	MAJOR          *xsql2.XSqlParam
	ID_NO          *xsql2.XSqlParam
	ETEL           *xsql2.XSqlParam
	QUALIFY        *xsql2.XSqlParam
	PWD            *xsql2.XSqlParam
	WECHAT         *xsql2.XSqlParam
	ADMIN          *xsql2.XSqlParam
	ID_PLACE       *xsql2.XSqlParam
	GRADUATE       *xsql2.XSqlParam
	DOC_COMPANY    *xsql2.XSqlParam
	TEL            *xsql2.XSqlParam
	NATIVE         *xsql2.XSqlParam
	LEDGER_START   *xsql2.XSqlParam
	CERTIFICATE    *xsql2.XSqlParam
	ATYPE          *xsql2.XSqlParam
	SEX            *xsql2.XSqlParam
	TRAIN_SALARY   *xsql2.XSqlParam
	LEDGER_DATE    *xsql2.XSqlParam
	SOCIAL_DATE    *xsql2.XSqlParam
	USERID         *xsql2.XSqlParam
	MARITAL        *xsql2.XSqlParam
	BANK_CARD      *xsql2.XSqlParam
	QQ             *xsql2.XSqlParam
	LEDGER         *xsql2.XSqlParam
	IID            *xsql2.XSqlParam
	PIC            *xsql2.XSqlParam
	STATE          *xsql2.XSqlParam
	DID            *xsql2.XSqlParam
	JOB            *xsql2.XSqlParam
	DESCRIPTION    *xsql2.XSqlParam
	BIRTH_TYPE     *xsql2.XSqlParam
	ID             *xsql2.XSqlParam
	LEDGER_COMPANY *xsql2.XSqlParam
	LEDGER_COUNTS  *xsql2.XSqlParam
	MUTUAL_FUND    *xsql2.XSqlParam
	LOGIN_TIME     *xsql2.XSqlParam
	BANK           *xsql2.XSqlParam
	E_MAIL         *xsql2.XSqlParam
	QUALIFY_SALARY *xsql2.XSqlParam
	AUTHORITY      *xsql2.XSqlParam
	EDUCATION      *xsql2.XSqlParam
	EMPLOYEESHIP   *xsql2.XSqlParam
	DOC_ADDR       *xsql2.XSqlParam
	ANAME          *xsql2.XSqlParam
	BIRTH          *xsql2.XSqlParam
	NATION         *xsql2.XSqlParam
	BIRTH_DATE     *xsql2.XSqlParam
	ENAME          *xsql2.XSqlParam
	SOCIAL_COMPANY *xsql2.XSqlParam
	ENTRY          *xsql2.XSqlParam
	LEAVE          *xsql2.XSqlParam
}

func (table *T_mbo_account) initObject() {
	table.NATION = &xsql2.XSqlParam{Name: "nation", Type_: "string", Target: table, AS_: ""}
	table.ENTRY = &xsql2.XSqlParam{Name: "entry", Type_: "int", Target: table, AS_: ""}
	table.LEAVE = &xsql2.XSqlParam{Name: "leave", Type_: "int", Target: table, AS_: ""}
	table.BIRTH_DATE = &xsql2.XSqlParam{Name: "birth_date", Type_: "string", Target: table, AS_: ""}
	table.ENAME = &xsql2.XSqlParam{Name: "ename", Type_: "string", Target: table, AS_: ""}
	table.SOCIAL_COMPANY = &xsql2.XSqlParam{Name: "social_company", Type_: "int", Target: table, AS_: ""}
	table.TITLE = &xsql2.XSqlParam{Name: "title", Type_: "string", Target: table, AS_: ""}
	table.ACCOUNT = &xsql2.XSqlParam{Name: "account", Type_: "string", Target: table, AS_: ""}
	table.TRAIN_DURATION = &xsql2.XSqlParam{Name: "train_duration", Type_: "int", Target: table, AS_: ""}
	table.LEDGER_TYPE = &xsql2.XSqlParam{Name: "ledger_type", Type_: "int", Target: table, AS_: ""}
	table.LEDGER_DURATIN = &xsql2.XSqlParam{Name: "ledger_duratin", Type_: "int", Target: table, AS_: ""}
	table.LEDGER_END = &xsql2.XSqlParam{Name: "ledger_end", Type_: "int", Target: table, AS_: ""}
	table.ETEL = &xsql2.XSqlParam{Name: "etel", Type_: "string", Target: table, AS_: ""}
	table.QUALIFY = &xsql2.XSqlParam{Name: "qualify", Type_: "int", Target: table, AS_: ""}
	table.PWD = &xsql2.XSqlParam{Name: "pwd", Type_: "string", Target: table, AS_: ""}
	table.WECHAT = &xsql2.XSqlParam{Name: "wechat", Type_: "string", Target: table, AS_: ""}
	table.SCHOOL = &xsql2.XSqlParam{Name: "school", Type_: "string", Target: table, AS_: ""}
	table.MAJOR = &xsql2.XSqlParam{Name: "major", Type_: "string", Target: table, AS_: ""}
	table.ID_NO = &xsql2.XSqlParam{Name: "id_no", Type_: "string", Target: table, AS_: ""}
	table.DOC_COMPANY = &xsql2.XSqlParam{Name: "doc_company", Type_: "int", Target: table, AS_: ""}
	table.TEL = &xsql2.XSqlParam{Name: "tel", Type_: "string", Target: table, AS_: ""}
	table.NATIVE = &xsql2.XSqlParam{Name: "native", Type_: "string", Target: table, AS_: ""}
	table.ADMIN = &xsql2.XSqlParam{Name: "admin", Type_: "int", Target: table, AS_: ""}
	table.ID_PLACE = &xsql2.XSqlParam{Name: "id_place", Type_: "string", Target: table, AS_: ""}
	table.GRADUATE = &xsql2.XSqlParam{Name: "graduate", Type_: "int", Target: table, AS_: ""}
	table.ATYPE = &xsql2.XSqlParam{Name: "atype", Type_: "int", Target: table, AS_: ""}
	table.SEX = &xsql2.XSqlParam{Name: "sex", Type_: "int", Target: table, AS_: ""}
	table.LEDGER_START = &xsql2.XSqlParam{Name: "ledger_start", Type_: "int", Target: table, AS_: ""}
	table.CERTIFICATE = &xsql2.XSqlParam{Name: "certificate", Type_: "string", Target: table, AS_: ""}
	table.USERID = &xsql2.XSqlParam{Name: "userid", Type_: "string", Target: table, AS_: ""}
	table.MARITAL = &xsql2.XSqlParam{Name: "marital", Type_: "int", Target: table, AS_: ""}
	table.TRAIN_SALARY = &xsql2.XSqlParam{Name: "train_salary", Type_: "int", Target: table, AS_: ""}
	table.LEDGER_DATE = &xsql2.XSqlParam{Name: "ledger_date", Type_: "int", Target: table, AS_: ""}
	table.SOCIAL_DATE = &xsql2.XSqlParam{Name: "social_date", Type_: "int", Target: table, AS_: ""}
	table.IID = &xsql2.XSqlParam{Name: "iid", Type_: "int", Target: table, AS_: ""}
	table.PIC = &xsql2.XSqlParam{Name: "pic", Type_: "string", Target: table, AS_: ""}
	table.BANK_CARD = &xsql2.XSqlParam{Name: "bank_card", Type_: "string", Target: table, AS_: ""}
	table.QQ = &xsql2.XSqlParam{Name: "QQ", Type_: "string", Target: table, AS_: ""}
	table.LEDGER = &xsql2.XSqlParam{Name: "ledger", Type_: "int", Target: table, AS_: ""}
	table.DID = &xsql2.XSqlParam{Name: "did", Type_: "int", Target: table, AS_: ""}
	table.JOB = &xsql2.XSqlParam{Name: "job", Type_: "string", Target: table, AS_: ""}
	table.STATE = &xsql2.XSqlParam{Name: "state", Type_: "int", Target: table, AS_: ""}
	table.DESCRIPTION = &xsql2.XSqlParam{Name: "description", Type_: "string", Target: table, AS_: ""}
	table.BIRTH_TYPE = &xsql2.XSqlParam{Name: "birth_type", Type_: "int", Target: table, AS_: ""}
	table.ID = &xsql2.XSqlParam{Name: "id", Type_: "int", Target: table, AS_: ""}
	table.LEDGER_COMPANY = &xsql2.XSqlParam{Name: "ledger_company", Type_: "int", Target: table, AS_: ""}
	table.LOGIN_TIME = &xsql2.XSqlParam{Name: "login_time", Type_: "int", Target: table, AS_: ""}
	table.BANK = &xsql2.XSqlParam{Name: "bank", Type_: "string", Target: table, AS_: ""}
	table.LEDGER_COUNTS = &xsql2.XSqlParam{Name: "ledger_counts", Type_: "int", Target: table, AS_: ""}
	table.MUTUAL_FUND = &xsql2.XSqlParam{Name: "mutual_fund", Type_: "int", Target: table, AS_: ""}
	table.E_MAIL = &xsql2.XSqlParam{Name: "e_mail", Type_: "string", Target: table, AS_: ""}
	table.AUTHORITY = &xsql2.XSqlParam{Name: "authority", Type_: "int", Target: table, AS_: ""}
	table.EDUCATION = &xsql2.XSqlParam{Name: "education", Type_: "int", Target: table, AS_: ""}
	table.QUALIFY_SALARY = &xsql2.XSqlParam{Name: "qualify_salary", Type_: "int", Target: table, AS_: ""}
	table.EMPLOYEESHIP = &xsql2.XSqlParam{Name: "employeeship", Type_: "int", Target: table, AS_: ""}
	table.ANAME = &xsql2.XSqlParam{Name: "aname", Type_: "string", Target: table, AS_: ""}
	table.BIRTH = &xsql2.XSqlParam{Name: "birth", Type_: "int", Target: table, AS_: ""}
	table.DOC_ADDR = &xsql2.XSqlParam{Name: "doc_addr", Type_: "string", Target: table, AS_: ""}
}
func (table *T_mbo_account) GetName() string {
	return "t_mbo_account"
}
func CreateT_mbo_account() *T_mbo_account {
	table := &T_mbo_account{}
	table.initObject()
	return table
}
