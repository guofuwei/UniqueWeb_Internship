package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	ID   int `type:"integer"`
	Name string
}

type StructMsg struct {
	FieldName  []string
	FieldType  []string
	FieldValue []interface{}
}

func GetStructMsg(arg interface{}) (structmsg *StructMsg) {
	structmsg = new(StructMsg)
	structmsg.FieldName = make([]string, 2)
	structmsg.FieldType = make([]string, 2)
	structmsg.FieldValue = make([]interface{}, 2)
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

func getTest(arg interface{}) {
	getType := reflect.TypeOf(arg)
	getValue := reflect.ValueOf(arg)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		// addr := get.Elem().Field(i)
		// fmt.Println(addr)
		// fieldinfo := field.Tag
		// // fmt.Println(reflect.TypeOf(value))
		// fmt.Println(fieldinfo.Get("type"))
		fmt.Printf("%s: %v = %v", field.Name, field.Type, value)
	}
}

func getTest2(arg interface{}) {
	addr := reflect.ValueOf(arg).Elem().Field(0).Addr().Interface()
	fmt.Println(addr)
}

func main() {
	// var s1 []string = []string{"hobby=10", "and", "ID=3"}
	//将一系列字符串连接为一个字符串，之间用sep来分隔
	// s2 := strings.Join(s1, " ")
	// querystr := "select * from task2 " + s2 + ";"
	// fmt.Println(querystr)
	// test := Test{0, "test"}
	// fmt.Println(test)
	getTest2(&Test{})
	// structmsg := GetStructMsg(test)
	// fmt.Println(structmsg)
}
