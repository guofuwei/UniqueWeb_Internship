package myorm

import (
	"testing"
	"time"
)

type UserInfo struct {
	ID     uint   `unique_index:"ID"`
	Name   string `index:"Name"`
	Gender string
	Hobby  string
}

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func TestGetStructMsg(t *testing.T) {
	u1 := UserInfo{1, "枯藤", "男", "篮球"}
	// u2 := UserInfo{2, "topgoer.com", "女", "足球"}
	structmsg := GetStructMsg(u1, 4)
	t.Log(*structmsg)
}

func TestGetRet(t *testing.T) {
	// mydb, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")

}

func TestAutoCreateTable(t *testing.T) {
	u1 := UserInfo{1, "枯藤", "男", "篮球"}
	mydb, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")
	if err != nil {
		t.Fatal(err)
	}
	err = AutoCreateTable(u1, mydb)
	if err != nil {
		t.Fatal(err)
	}
}
