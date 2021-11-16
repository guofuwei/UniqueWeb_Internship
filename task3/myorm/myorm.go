package myorm

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type MYDB struct {
	dbptr     *sql.DB
	tablename string
	valueNum  int
	presql    string
}

// func getTest(arg interface{}) {
// 	getType := reflect.TypeOf(arg)
// 	getValue := reflect.ValueOf(arg)
// 	for i := 0; i < getType.NumField(); i++ {
// 		field := getType.Field(i)
// 		value := getValue.Field(i).Interface()
// 		fmt.Println(reflect.TypeOf(value))
// 		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
// 	}
// }

//将结构体自动创建Table
func (mydb *MYDB) CreateTable(struc interface{}) (err error) {
	err = AutoCreateTable(struc, mydb)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//获取数据库链接
func GetDB(args ...interface{}) (mydb *MYDB, err error) {
	mydb = new(MYDB)
	//开始处理参数
	if len(args) != 2 {
		err = errors.New("invalid database source")
		return nil, err
	}
	//开始获取参数
	driver := args[0].(string)
	source := args[1].(string)
	//连接并测试
	mydb.dbptr, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = mydb.dbptr.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return mydb, err
}

//插入数据
func (mydb *MYDB) Insert(args ...interface{}) (err error) {
	mydb.presql = ""
	leng := mydb.valueNum
	for _, arg := range args {
		structmsg := GetStructMsg(arg, leng)
		insertfield := strings.Join(structmsg.FieldName, ",")
		insertvalue := strings.Repeat("?,", mydb.valueNum)
		insertvalue = strings.TrimSuffix(insertvalue, ",")
		stmt, err := mydb.dbptr.Prepare(fmt.Sprintf("insert into %s(%s) values(%s);", mydb.tablename, insertfield, insertvalue))
		if err != nil {
			log.Fatal(err)
			return err
		}
		_, err = stmt.Exec(structmsg.FieldValue...)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	// stmt, err := mydb.dbptr.Prepare("insert into task2 values($1,$2,$3)")
	// _, err = stmt.Exec(filename, code, expiretime)
	return nil
}

//查找所有的数据
func (mydb *MYDB) FindAll(struc interface{}, args ...string) (*[]interface{}, error) {
	str := strings.Join(args, " ")
	mydb.presql = str
	// querystr := "select " + query + " from t " + str + ";"
	querystr := fmt.Sprintf("select * from %s where %s;", mydb.tablename, str)
	rows, err := mydb.dbptr.Query(querystr)
	if err != nil {
		log.Fatal(err)
		defer rows.Close()
		return nil, err
	}
	return GetAllRet(rows, struc), err
}

//查找单个数据
func (mydb *MYDB) FindOne(strucptr interface{}, args ...string) (*[]interface{}, error) {
	str := strings.Join(args, " ")
	// leng := len(strings.Split(query, ","))
	mydb.presql = str
	querystr := fmt.Sprintf("select * from %s where %s;", mydb.tablename, str)
	row := mydb.dbptr.QueryRow(querystr)
	return GetRet(row, strucptr)
}

//删除
func (mydb *MYDB) Delete() error {
	if mydb.presql == "" {
		return errors.New("there is no parameter")
	}
	delsql := fmt.Sprintf("delete from %s where %s;", mydb.tablename, mydb.presql)
	mydb.dbptr.Exec(delsql)
	return nil
}

//更新
func (mydb *MYDB) Update(upvalue string) error {
	if mydb.presql == "" {
		return errors.New("there is no parameter")
	}
	upsql := fmt.Sprintf("update %s set %s where %s;", mydb.tablename, upvalue, mydb.presql)
	_, err := mydb.dbptr.Exec(upsql)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
