package myorm

import (
	"reflect"
	"time"
)

func MyormModel(value reflect.Value) (newvalue string) {
	switch value.Interface().(type) {
	case int, int32:
		newvalue = "interger"
	case int8:
		newvalue = "tinyint"
	case int16:
		newvalue = "smallint"
	case int64:
		newvalue = "bigint"
	case uint, uint32:
		newvalue = "integer unsigned"
	case uint16:
		newvalue = "smallint unsigned"
	case uint64:
		newvalue = "bigint unsigned"
	case float32, float64:
		newvalue = "double precision"
	case bool:
		newvalue = "bool"
	case string:
		newvalue = "varchar(100)"
	case time.Time:
		newvalue = "datetime"
	case byte:
		newvalue = "tinyint unsigned"
	}
	return newvalue
}
