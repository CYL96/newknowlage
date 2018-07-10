package mysql

/*

数据库调用方法说明

首先进行InitSql()数据库初始化操作，告诉数据库要去操作哪个数据库

CreateInstance()创建查询对象，每次查询创建一个对象，方便自动化维护查询对象

然后调用增删改查的各种方法，进行语句拼接，也可以直接调用Qurey方法，直接写语句

其中查询方法可以使用三种方式
1、select 把要查询的表名称以及所属字段及字段属性列入其中 如：Select("users","id","int","name","string","phone","string","age","float")
2、先调用SetTableName或SetTableColType或者SetTableColTypeString对要查询的数据库的字段进行初始化设计，方便查询到以后进行序列化的时候自动转换
推荐使用SetTableColType或者SetTableColTypeString，这样减少数据库请求次数
SetTableName会先去调取数据库字段属性，调取成功以后再去调用实际请求的查询语句，这样方便开发，但不利于数据库本身
3、先调用Select2()，不需输入类型，如Select("users","id","name","phone",,"age") 或者 Select("users","*")
然后调用SetTableColType或者SetTableColTypeString对要查询的数据库的字段进行初始化设计，方便查询到以后进行序列化的时候自动转换，如果里面没有采用重命名的字段名称，可以忽略

然后调用Execute方法，进行语句执行，返回结果，也可以直接调用ExecuteForJson，返回json序列化后的字符串
*/


import (
	"database/sql"
	"sync"
	"bytes"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"nqc.cn/log"
	"strings"
	"strconv"
	"encoding/json"
)

const LifeTime int64  =  60 * 60

type XSql struct {
	db         *sql.DB
	name string
	password string
	ip string
	port string
	sqlName string
	mLock *sync.RWMutex
	time_last int64
	//DB *sql.DB
}

type XSqlOrder struct {
	xs         *XSql
	reqBuffer  bytes.Buffer
	//selectKeys map[string]string
	tableName  []string
	colType map[string]string
	ch uint8
}

func CreateInstance(xs *XSql) *XSqlOrder {
	o := new (XSqlOrder)
	o.xs = xs
	//o.selectKeys = make(map[string]string)
	o.colType = make(map[string]string)
	o.reqBuffer.Grow(4096)
	return o
}
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//初始化数据库
func InitSql(name string, password string, ip string, port string, sqlName string) *XSql {
	db := createDB(name,password,ip,port,sqlName)
	fmt.Println("初始化数据库成功")
	s := new(XSql)
	s.mLock = new(sync.RWMutex)
	s.db = db
	s.name = name
	s.password = password
	s.ip = ip
	s.port = port
	s.sqlName = sqlName
	s.time_last = time.Now().Unix()
	//s.ch = make(chan uint8,100)

	//go timer(s)

	return s
}
func createDB(name string, password string, ip string, port string, sqlName string) *sql.DB {
	db, err := sql.Open("mysql", name+":"+password+"@tcp("+ip+":"+port+")/"+sqlName+"?charset=utf8mb4")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.Ping()

	checkErr(err)


	return db
}

func (s *XSqlOrder) ClearColType() {
	s.colType = make(map[string]string)
}
func (s *XSqlOrder) ClearBuffer() {
	s.reqBuffer.Reset()
}

func (s *XSqlOrder) Select(name string, keys ...string) { //第一个参数为列表名，第二个为参数类型，int，string，float
	//parses(keys)
	fmt.Println("SELECT..........")
	s.ClearBuffer()
	s.reqBuffer.WriteString("select ")
	s.ClearColType()
	//selectKeys = keys
	var nextKey string
	for i:= 0;i<len(keys);i++ {
		if i%2 == 0 {
			s.reqBuffer.WriteString(keys[i])
			if i != len(keys)-2{
				s.reqBuffer.WriteString(",")
			}
			nextKey = strings.ToLower(keys[i])
		} else {
			s.colType[nextKey] = strings.ToLower(keys[i])
		}
	}
	//s.reqBuffer.Truncate(s.reqBuffer.Len()-1)
	fmt.Println(s.reqBuffer.String())

	s.reqBuffer.WriteString(" from ")
	s.reqBuffer.WriteString( name)

	//s.selectKeys = selectKeys
	s.tableName = append(s.tableName,name)

}

func (s *XSqlOrder) Insert(values map[string]interface{}, name string) {

	s.ClearBuffer()
	s.reqBuffer.WriteString("INSERT INTO ")
	s.reqBuffer.WriteString( name)
	s.reqBuffer.WriteString(" (")
	//s.selectKeys = make(map[string]string)

	var valueString bytes.Buffer
	valueString.Grow(2048)
	//sorted_keys := make([]string, 0)
	var index = 0
	for key, value := range values {
		index ++
		s.reqBuffer.WriteString(key)
		if index != len(values) {
			s.reqBuffer.WriteString(",")
		}
		//valueString += value + ","
		//value := values[k]
		switch value.(type) {
		case int:
			{

				valueString .WriteString( strconv.FormatInt(int64(value.(int)), 10))
				if index != len(values) {
					valueString .WriteString( ",")
				}
			}
			break
		case int64:
			{

				valueString .WriteString(  strconv.FormatInt(value.(int64), 10))
				if index != len(values) {
					valueString .WriteString( ",")
				}
			}
			break
		case float32:
			{

				valueString .WriteString(strconv.FormatFloat(float64(value.(float32)), 'f', 6, 32))
				if index != len(values) {
					valueString .WriteString(",")
				}
			}
			break
		case float64:
			{

				valueString .WriteString( strconv.FormatFloat(value.(float64), 'f', 6, 64))
				if index != len(values) {
					valueString .WriteString( ",")
				}
			}
			break
		case string:
			{

				valueString .WriteString( "\"")
				valueString .WriteString( value.(string) )
				valueString .WriteString("\"")
				if index != len(values) {
					valueString .WriteString( ",")
				}

			}
			break
		case []byte:
			{

				valueString .WriteString( "\"")
				valueString .WriteString( string(value.([]byte)))
				valueString .WriteString("\"")
				if index != len(values) {
					valueString .WriteString( ",")
				}
			}
			break
		}
		//sorted_keys = append(sorted_keys, key)
	}
	s.reqBuffer.WriteString(") VALUES ( ")
	s.reqBuffer.WriteString( valueString.String())
	//s.reqString = Substr(reqString, 0, len(reqString)-1)
	s.reqBuffer.WriteString( ")")
	fmt.Println(s.reqBuffer.String())

}
func (s *XSqlOrder) Insert_(values map[string]interface{}, name string) {
	s.ClearBuffer()
	s.reqBuffer.WriteString("INSERT INTO ")
	s.reqBuffer.WriteString(name)
	s.reqBuffer.WriteString(" (")
	//s.selectKeys = make(map[string]string)

	var valueString bytes.Buffer
	//sorted_keys := make([]string, 0)
	var index = 0
	for key, value := range values {
		index ++
		s.reqBuffer.WriteString( key)
		if index != len(values) {
			s.reqBuffer.WriteString( ",")
		}
		//valueString += value + ","
		//value := values[k]
		switch value.(type) {
		case int:
			{

				valueString.WriteString( strconv.FormatInt(int64(value.(int)), 10))
				if index != len(values) {
					valueString.WriteString( ",")
				}
			}
			break
		case int64:
			{

				valueString.WriteString( strconv.FormatInt(value.(int64), 10))
				if index != len(values) {
					valueString.WriteString( ",")
				}
			}
			break
		case float32:
			{

				valueString.WriteString( strconv.FormatFloat(float64(value.(float32)), 'f', 6, 32))
				if index != len(values) {
					valueString.WriteString( ",")
				}
			}
			break
		case float64:
			{

				valueString.WriteString( strconv.FormatFloat(value.(float64), 'f', 6, 64))
				if index != len(values) {
					valueString.WriteString( ",")
				}
			}
			break
		case string:
			{

				valueString.WriteString( "'")
				valueString.WriteString(value.(string))
				valueString.WriteString("'")
				if index != len(values) {
					valueString.WriteString( ",")
				}

			}
			break
		case []byte:
			{

				valueString.WriteString( "\"")
				valueString.WriteString(string(value.([]byte)))
				valueString.WriteString("\"")
				if index != len(values) {
					valueString.WriteString( ",")
				}
			}
			break
		}
		//sorted_keys = append(sorted_keys, key)
	}
	s.reqBuffer.WriteString(") VALUES ( ")
	s.reqBuffer.WriteString( valueString.String())
	//s.reqString = Substr(reqString, 0, len(reqString)-1)
	s.reqBuffer.WriteString( ")")

}
func (s *XSqlOrder) MulitInsert(list []map[string]interface{}, name string) {

	s.ClearBuffer()
	s.reqBuffer.WriteString("INSERT INTO ")
	s.reqBuffer.WriteString(name)
	s.reqBuffer.WriteString(" (")
	//s.selectKeys = make(map[string]string)

	var valueString bytes.Buffer
	//sorted_keys := make([]string, 0)
	var key_list []string
	if len(list) > 0 {
		var index = 0;
		//key 为字段
		for key, _ := range list[0] {
			key_list = append(key_list,key)
			index++
			s.reqBuffer.WriteString( key)
			if index != len(list[0]) {
				s.reqBuffer.WriteString(",")
			}
		}
	}

	for index_k , values := range list {
		index_k++
		//var index = 0
		for index,key_value := range key_list {
			value := values[key_value]
			index ++
			switch value.(type) {
			case int:
				{

					valueString.WriteString( strconv.FormatInt(int64(value.(int)), 10))
					if index != len(values) {
						valueString.WriteString( ",")
					}
				}
				break
			case int64:
				{

					valueString.WriteString( strconv.FormatInt(value.(int64), 10))
					if index != len(values) {
						valueString.WriteString( ",")
					}
				}
				break
			case float32:
				{

					valueString.WriteString( strconv.FormatFloat(float64(value.(float32)), 'f', 6, 32))
					if index != len(values) {
						valueString.WriteString( ",")
					}
				}
				break
			case float64:
				{

					valueString.WriteString( strconv.FormatFloat(value.(float64), 'f', 6, 64))
					if index != len(values) {
						valueString.WriteString( ",")
					}
				}
				break
			case string:
				{

					valueString.WriteString( "\"")
					valueString.WriteString(value.(string))
					valueString.WriteString("\"")
					if index != len(values) {
						valueString.WriteString( ",")
					}
				}
				break
			case []byte:
				{

					valueString.WriteString( "\"")
					valueString.WriteString( string(value.([]byte)))
					valueString.WriteString("\"")
					if index != len(values) {
						valueString.WriteString( ",")
					}
				}
				break
			}
		}

		if index_k != len(list)  {
			valueString.WriteString( "),( ")
		}
	}
	fmt.Println("reqString: ",s.reqBuffer.String())
	fmt.Println("valueString: ",valueString)

	s.reqBuffer.WriteString(") VALUES ( ")
	s.reqBuffer.WriteString(valueString.String())
	fmt.Println("reqString: ",s.reqBuffer.String())
	//s.reqString = Substr(reqString, 0, len(reqString)-1)
	s.reqBuffer.WriteString(")")

}

func (s *XSqlOrder) Update(values map[string]interface{}, name string) {
	s.ClearBuffer()
	s.reqBuffer.WriteString( "UPDATE ")
	s.reqBuffer.WriteString(name)
	s.reqBuffer.WriteString(" SET ")
	//s.selectKeys = make(map[string]string)
	var index = 0
	for key, value := range values {
		index++
		switch value.(type) {
		case int:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatInt(int64(value.(int)), 10))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case int64:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatInt(value.(int64), 10))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case float32:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatFloat(float64(value.(float32)), 'f', 6, 32))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case float64:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 6, 64))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case string:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString("\"")
				s.reqBuffer.WriteString(value.(string))
				s.reqBuffer.WriteString("\"")
				if index != len(values) {
					s.reqBuffer.WriteString( ",")
				}
			}
			break
		case []byte:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString("\"")
				s.reqBuffer.WriteString(string(value.([]byte)) + "\"")
				if index != len(values) {
					s.reqBuffer.WriteString( ",")
				}
			}
			break
		}
	}

	//s.reqString = Substr(reqString, 0, len(reqString)-2)

}
func (s *XSqlOrder) Update_(values map[string]interface{}, name string) {
	s.ClearBuffer()
	s.reqBuffer.WriteString( "UPDATE ")
	s.reqBuffer.WriteString(name)
	s.reqBuffer.WriteString(" SET ")
	//s.selectKeys = make(map[string]string)
	var index = 0
	for key, value := range values {
		index++
		switch value.(type) {
		case int:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatInt(int64(value.(int)), 10))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case int64:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatInt(value.(int64), 10))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case float32:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatFloat(float64(value.(float32)), 'f', 6, 32))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case float64:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 6, 64))
				if index != len(values) {
					s.reqBuffer.WriteString( ", ")
				}
			}
			break
		case string:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString("'")
				s.reqBuffer.WriteString(value.(string))
				s.reqBuffer.WriteString("'")
				if index != len(values) {
					s.reqBuffer.WriteString( ",")
				}
			}
			break
		case []byte:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString("\"")
				s.reqBuffer.WriteString(string(value.([]byte)) + "\"")
				if index != len(values) {
					s.reqBuffer.WriteString( ",")
				}
			}
			break
		}
	}

	//s.reqString = Substr(reqString, 0, len(reqString)-2)

}


func (s *XSqlOrder) Delete(name string) {
	s.ClearBuffer()
	s.reqBuffer.WriteString("DELETE FROM ")
	s.reqBuffer.WriteString(name)
	//s.selectKeys = make(map[string]string)
}
func (s *XSqlOrder) Where(values map[string]interface{}) {
	s.reqBuffer.WriteString( " where ")
	var index = 0
	for key, value := range values {
		index ++
		switch value.(type) {
		case int:
			{
				s.reqBuffer.WriteString(  key)
				s.reqBuffer.WriteString( "=")
				s.reqBuffer.WriteString( strconv.FormatInt(int64(value.(int)), 10))
			}
			break
		case int64:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString( "=")

				s.reqBuffer.WriteString(strconv.FormatInt(value.(int64), 10))
			}
			break
		case float32:
			{
				s.reqBuffer.WriteString(  key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatFloat(float64(value.(float32)), 'f', 0, 32))
			}
			break
		case float64:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 0, 64))
			}
			break
		case string:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString("\"")
				s.reqBuffer.WriteString(value.(string))
				s.reqBuffer.WriteString("\"")
			}
			break
		case []byte:
			{
				s.reqBuffer.WriteString( key)
				s.reqBuffer.WriteString("=")
				s.reqBuffer.WriteString("\"")
				s.reqBuffer.WriteString(string(value.([]byte)))
			}
			break
		}
		if index != len(values) {
			s.reqBuffer.WriteString( " AND ")
		}
	}
	//s.reqString = Substr(reqString, 0, len(reqString)-4)
}
func (s *XSqlOrder) AddSuf(suffixes string) { //添加sql尾部参数
	s.reqBuffer.WriteString(" ")
	s.reqBuffer.WriteString(suffixes)

}

func (s *XSqlOrder) Count(name string) {
	s.colType = make(map[string]string)
	s.ClearBuffer()
	s.reqBuffer.WriteString( "select count(*) as count from ")
	s.reqBuffer.WriteString(name)
	s.colType["count"] = "int"
}
func (s *XSqlOrder) CountMore(name string,tag string) {
	s.colType = make(map[string]string)
	s.ClearBuffer()
	s.reqBuffer.WriteString( "select count(")
	s.reqBuffer.WriteString(tag)
	s.reqBuffer.WriteString(") as count from ")
	s.reqBuffer.WriteString(name)
	s.colType["count"] = "int"
}

func (s *XSqlOrder) Value() int64 {
	list := s.Execute()
	for _, value := range list {
		for _, value_c := range value {
			return value_c.(int64)
		}
	}
	return 0
}

func (s *XSqlOrder) Qurey(suffixes string) { //执行sql语句
	//s.colType = make(map[string]string)
	s.ClearBuffer()
	s.reqBuffer.WriteString( suffixes)
}

func (s *XSqlOrder) ExecuteForJson() string { //执行sql语句得到json

	body, err := json.Marshal(s.Execute())
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
func (s *XSqlOrder) GetSQLString() string {
	return s.reqBuffer.String()
}
func (s *XSqlOrder) ExecuteNoResult() {
	//SQL
	fmt.Println("ExecuteNoResult执行sql语句: " + s.reqBuffer.String())
	s.ch = 0
	s.xs.mLock.RLock()
	go timer(s)
	rows, _ := s.xs.db.Query(s.reqBuffer.String())
	s.xs.mLock.RUnlock()
	s.ch = 1
	rows.Close()
}
func (s *XSqlOrder) ExecuteForLastInsertId() int64{
	//SQL
	fmt.Println("ExecuteForLastInsertId执行sql语句: " + s.reqBuffer.String())
	s.ch = 0
	s.xs.mLock.RLock()
	go timer(s)
	ret,err := s.xs.db.Exec(s.reqBuffer.String())
	if err != nil{
		fmt.Println("err:",err)
	}
	LastInsertId,err := ret.LastInsertId()
	if err != nil{
		fmt.Println("err:",err)
	}
	//if RowsAffected, err := ret.RowsAffected(); nil == err {
	//	fmt.Println("RowsAffected:", RowsAffected)
	//}
	s.xs.mLock.RUnlock()
	s.ch = 1
	//defer ret.Close()
	return LastInsertId
}
func (s *XSqlOrder) Execute2() []map[string]interface{} { //SQL
	fmt.Println("Execute执行sql语句: " + s.reqBuffer.String())

	s.ch = 0

	s.xs.mLock.RLock()
	go timer(s)

	rows, err := s.xs.db.Query(s.reqBuffer.String())

	s.xs.mLock.RUnlock()
	s.ch = 1
	if err != nil {
		fmt.Println("error: ",err)
		s.xs.mLock.RLock()
		s.xs.db.Close()
		db := createDB(s.xs.name,s.xs.password,s.xs.ip,s.xs.port,s.xs.sqlName)
		s.xs.db = db
		s.xs.time_last = time.Now().Unix()

		rows, err = s.xs.db.Query(s.reqBuffer.String())
		checkErr(err)
		return nil
	}

	defer rows.Close()

	columns, err2 := rows.Columns()
	if err2 != nil {
		log.Write(err2) // proper error handling instead of panic in your app
		return nil
	}

	if len(columns) <= 0 {
		return nil
	}

	// Make a slice for the values
	values := make([]interface{}, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]

	}
	var results []map[string]interface{}

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		t := make(map[string]interface{})
		//fmt.Println(len(selectKeys))
		for i, col := range values {
			//	fmt.Println(i)
			//fmt.Println(col)
			//fmt.Println(selectKeys[columns[i]])
			if col == nil {
				t[columns[i]] = nil
			} else {
				switch s.colType[strings.ToLower(columns[i])] {
				case "int":
					{

						t[columns[i]] = byte2Int(col.([]byte))
					}
					break
				case "float":
					{
						fmt.Println(columns[i])
						t[columns[i]] = byte2Float(col.([]byte))
					}
					break
				case "string":
					{

						t[columns[i]] = byte2String(col.([]byte))
					}
					break
				default:
					{

					}
					break
				}
			}

		}
		results = append(results, t)

	}
	//rows.Close()

	return results

}
func (s *XSqlOrder) Execute() []map[string]interface{} { //SQL
	fmt.Println("Execute执行sql语句: " + s.reqBuffer.String())

	s.ch = 0
	s.xs.mLock.RLock()

	go timer(s)

	rows, err := s.xs.db.Query(s.reqBuffer.String())


	s.xs.mLock.RUnlock()
	s.ch = 1
	if err != nil {
		fmt.Println("error: ",err)
		s.xs.mLock.RLock()
		s.xs.db.Close()
		db := createDB(s.xs.name,s.xs.password,s.xs.ip,s.xs.port,s.xs.sqlName)
		s.xs.db = db
		s.xs.time_last = time.Now().Unix()

		rows, err = s.xs.db.Query(s.reqBuffer.String())
		defer rows.Close()
		checkErr(err)
		return nil
	}

	defer rows.Close()
	columns, err2 := rows.Columns()
	if err2 != nil {
		log.Write(err2) // proper error handling instead of panic in your app
		return nil
	}

	if len(columns) <= 0 {
		return nil
	}

	// Make a slice for the values
	values := make([]interface{}, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]

		//fmt.Println(columns[i])
	}
	//return nil
	var results []map[string]interface{}

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("come in")
			//panic(err.Error()) // proper error handling instead of panic in your app
		}
		t := make(map[string]interface{})
		for i, col := range values {

			if col == nil {
				t[columns[i]] = nil
			} else {
				switch s.colType[strings.ToLower(columns[i])] {
				case "int":
					{

						t[columns[i]] = byte2Int(col.([]byte))
					}
					break
				case "float":
					{
						//fmt.Println(columns[i])
						t[columns[i]] = byte2Float(col.([]byte))
					}
					break
				case "string":
					{

						t[columns[i]] = byte2String(col.([]byte))
					}
					break
				default:
					{
						t[columns[i]] = getInitValue(col.([]byte))
					}
					break
				}
			}

		}
		results = append(results, t)

	}

	return results
}

func byte2Int(value []byte) int64 {

	result, err := strconv.ParseInt(string(value), 10, 64)
	checkErr(err)
	return result
}
func byte2Float(value []byte) float64 {

	result, err := strconv.ParseFloat(string(value), 64)
	checkErr(err)
	return result
}
func byte2String(value []byte) string {
	return string(value)
}
func getInitValue(pval []byte) interface{} {
	result_int,ok := ParseInt(pval)
	if !ok {
		result_float, ok := ParseFloat(pval)
		if !ok {
			//fmt.Println("string")
			return string(pval)
		}
		//fmt.Println("float")
		return result_float
	}else{
		s := string(pval)
		a := strings.Split(s,"0")
		if strings.EqualFold(a[0],""){
			return string(pval)
		}
		//fmt.Println("int")
		return result_int
	}
}
func ParseInt(value []byte) (int64,bool) {
	result, err := strconv.ParseInt(string(value), 10, 64)
	if err != nil {
		return 0,false
	}
	return result,true
}
func ParseFloat(value []byte) (float64,bool) {
	result, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		return 0,false
	}
	return result,true
}

func (s *XSqlOrder)SetTableName(name ...string) map[string]string {
	s.tableName = name
	var sqlbuf bytes.Buffer
	sqlbuf.Grow(4096)
	sqlbuf.WriteString( "SELECT column_name,data_type FROM INFORMATION_SCHEMA.columns WHERE TABLE_NAME='")
	for i :=0;i<len(name);i++{
		sqlbuf.WriteString( name[i])
		if i != len(name) - 1 {
			sqlbuf.WriteString(  "' OR TABLE_NAME='")
		}
	}
	sqlbuf.WriteString(  "' ")
	s.ch = 0
	s.xs.mLock.RLock()
	rows,err := s.xs.db.Query(sqlbuf.String())

	s.xs.mLock.RUnlock()
	s.ch = 1
	if err != nil {
		fmt.Println("error: ",err)
		s.xs.mLock.RLock()
		s.xs.db.Close()
		db := createDB(s.xs.name,s.xs.password,s.xs.ip,s.xs.port,s.xs.sqlName)
		s.xs.db = db
		s.xs.time_last = time.Now().Unix()

		rows, err = s.xs.db.Query(s.reqBuffer.String())
		defer rows.Close()
		checkErr(err)
		return nil
	}
	defer rows.Close()
	//t := make(map[string]string)
	for rows.Next() {
		var column_name string
		var data_type string
		err = rows.Scan(&column_name,&data_type)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if strings.Contains(data_type,"int")  {
			s.colType[column_name] = "int"
		} else if strings.Contains(data_type,"float")  {
			s.colType[column_name] = "float"
		} else if strings.Contains(data_type,"bool")  {
			s.colType[column_name] = "bool"
		} else {
			s.colType[column_name] = "string"
		}
	}
	//rows.Close()
	//s.colType = t
	return s.colType
}
//int  float string
func (s *XSqlOrder) SetTableColType(data_type map[string]string) {
	//SQL
	for key,value := range data_type {
		s.colType[strings.ToLower(key)] = strings.ToLower(value)
	}
}
func (s *XSqlOrder) SetTableColTypeString(data_types ...string) {
	var nextKey string
	for i:=0;i<len(data_types);i++ {
		if i%2 == 0 {
			nextKey = strings.ToLower(data_types[i])
		} else {
			s.colType[nextKey] = strings.ToLower(data_types[i])
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Write(err)
	}
}
func timer(s *XSqlOrder) {
	//nt := int64(time.Now().Unix())
	/*timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			{
				v := <- s.ch
				fmt.Println(v)
				//s.createNewDB()
				if v == 1 {
					return
				}
			}
		}
	}*/
	fmt.Println("timer")
	for i := 0 ; i < 1000 ; i ++ {
		//fmt.Println("ti8mer:" , i)
		//s.createNewDB()
		if s.ch == 1 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("重新生成:")
	s.xs.db.Close()
	db := createDB(s.xs.name,s.xs.password,s.xs.ip,s.xs.port,s.xs.sqlName)
	s.xs.db = db
	s.xs.time_last = time.Now().Unix()
	fmt.Println("重新生成:OK")
}