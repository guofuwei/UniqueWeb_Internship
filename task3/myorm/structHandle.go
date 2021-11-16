package myorm

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type StructMsg struct {
	FieldName  []string
	FieldType  []string
	FieldValue []interface{}
}

func GetStructMsg(arg interface{}, leng int) (structmsg *StructMsg) {
	structmsg = new(StructMsg)
	structmsg.FieldName = make([]string, leng)
	structmsg.FieldType = make([]string, leng)
	structmsg.FieldValue = make([]interface{}, leng)
	getType := reflect.TypeOf(arg)
	getValue := reflect.ValueOf(arg)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		structmsg.FieldName[i] = field.Name
		structmsg.FieldType[i] = field.Type.Name()
		structmsg.FieldValue[i] = value
		// fmt.Println(reflect.TypeOf(value))
		// fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	return structmsg
}

func GetAllRet(rows *sql.Rows, strucptr interface{}) *[]interface{} {
	result := make([]interface{}, 0)
	s := reflect.ValueOf(strucptr).Elem()
	leng := s.NumField()
	onerow := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		onerow[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		err := rows.Scan(onerow...)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		result = append(result, s.Interface())
	}
	return &result
}

func GetRet(row *sql.Row, strucptr interface{}) (*[]interface{}, error) {
	s := reflect.ValueOf(strucptr).Elem()
	leng := s.NumField()
	onerow := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		onerow[i] = s.Field(i).Addr().Interface()
	}
	err := row.Scan(onerow...)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		err = errors.New("there is no ret")
		return nil, err
	}
	return &onerow, nil
}

func AutoCreateTable(struc interface{}, mydb *MYDB) (err error) {
	s := reflect.ValueOf(struc)
	t := reflect.TypeOf(struc)
	leng := s.NumField()
	mydb.valueNum = leng
	mydb.tablename = reflect.TypeOf(struc).Name()
	mydb.presql = ""
	//
	fieldname := make([]string, leng)
	onerowsql := make([]string, leng)
	indexstr := make([]string, leng)
	uniqueindexstr := make([]string, leng)
	primary_key_name := ""
	//开始处理结构体
	// t := reflect.TypeOf(struc)
	// if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
	// 	return fmt.Errorf("参数应该为结构体指针")
	// }
	count1 := 0
	count2 := 0
	for i := 0; i < leng; i++ {
		//存储字段名
		fieldname[i] = t.Field(i).Name
		onerowsql[i] += fmt.Sprintf("`%s` ", fieldname[i])
		//尝取出tag
		tag := t.Field(i).Tag
		//对tag进行处理
		theprimarykey := tag.Get("primary_key")
		thetype := tag.Get("type")
		theindex := tag.Get("index")
		theuniqueindex := tag.Get("unique_index")
		//转换go到mysql的字段类型
		modelType := MyormModel(s.Field(i))
		//开始处理sql
		if thetype == "" {
			onerowsql[i] += modelType
		} else {
			onerowsql[i] += thetype
		}
		if theprimarykey != "" {
			primary_key_name += t.Field(i).Name
		}
		if theindex != "" {
			indexstr[count1] = theindex
			count1 += 1
		}
		if theuniqueindex != "" {
			uniqueindexstr[count2] = theuniqueindex
			count2 += 1
		}
	}
	fieldsql := strings.Join(onerowsql, ",")
	if primary_key_name != "" {
		fieldsql += fmt.Sprintf(",PRIMARY KEY (`%s`)", primary_key_name)
	} else {
		fieldsql += fmt.Sprintf(",PRIMARY KEY (`%s`)", fieldname[0])
	}
	// 开始进行索引的生成
	for _, oneindex := range indexstr {
		if oneindex == "" {
			break
		}
		fieldsql += fmt.Sprintf(",INDEX idx_%s (%s)", oneindex, oneindex)
	}
	for _, oneunindex := range uniqueindexstr {
		if oneunindex == "" {
			break
		}
		fieldsql += fmt.Sprintf(",UNIQUE idx_%s (%s)", oneunindex, oneunindex)
	}
	//开始生成最后的sql语句
	allsql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", mydb.tablename, fieldsql)
	// fmt.Println(allsql)
	_, err = mydb.dbptr.Exec(allsql)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
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
