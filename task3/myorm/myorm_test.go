package myorm

import "testing"

func TestGetDB(t *testing.T) {
	_, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateTable(t *testing.T) {
	mydb, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")
	if err != nil {
		t.Fatal(err)
	}
	err = mydb.CreateTable(UserInfo{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	mydb, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")
	if err != nil {
		t.Fatal(err)
	}
	err = mydb.CreateTable(UserInfo{})
	if err != nil {
		t.Fatal(err)
	}
	u2 := UserInfo{3, "topgoer.com", "女", "足球"}
	mydb.Insert(u2)
}

func TestFindOne(t *testing.T) {
	mydb, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")
	if err != nil {
		t.Fatal(err)
	}
	err = mydb.CreateTable(UserInfo{})
	if err != nil {
		t.Fatal(err)
	}
	ret, err := mydb.FindOne(&UserInfo{}, "ID=3")
	if err != nil {
		t.Fatal(err)
	}
	err = mydb.Update("Name='xiaohong'")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*ret...)
}

func TestFindAll(t *testing.T) {
	mydb, err := GetDB("mysql", "task3:123456@tcp(127.0.0.1:3306)/task3")
	if err != nil {
		t.Fatal(err)
	}
	err = mydb.CreateTable(UserInfo{})
	if err != nil {
		t.Fatal(err)
	}
	ret, err := mydb.FindAll(&UserInfo{}, "Gender=\"man\"")
	mydb.Delete()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*ret...)
}
