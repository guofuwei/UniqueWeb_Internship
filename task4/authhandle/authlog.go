package authhandle

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserInfo struct {
	Username     string `db:"username"`
	Password     string `db:"password"`
	Identify     string `db:"identify"`
	WriteLog     bool   `db:"write_log"`
	CatOtherLog  bool   `db:"cat_other_log"`
	EditOtherLog bool   `db:"edit_other_log"`
}

var Db *sqlx.DB
var err error

func init() {
	Db, err = sqlx.Open("mysql", "task:123456@tcp(127.0.0.1:3306)/task")
	if err != nil {
		log.Fatal(err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetLogAuth(username string) map[string]bool {
	user := UserInfo{}
	err := Db.Get(&user, "select * from userinfo where username=?;", username)
	if err != nil {
		log.Fatal(err)
	}
	logAuth := make(map[string]bool, 3)
	logAuth["writeLog"] = user.WriteLog
	logAuth["catOtherLog"] = user.CatOtherLog
	logAuth["editOtherLog"] = user.EditOtherLog
	return logAuth
}
