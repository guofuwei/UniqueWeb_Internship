package db_operate

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "practice_task2"
	password = "123456"
	dbname   = "practice_task2"
)

type Code struct {
	Filename   string
	Expiretime int64
	Code       string
}

var Db *sql.DB

func init() {
	Db = GetDB()
}

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("successfull connected!")
	return db
}

func Insert(db *sql.DB, filename string, code string, expiretime int64) {
	stmt, err := db.Prepare("insert into task2(filename,code,expiretime) values($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(filename, code, expiretime)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert into user_tbl success")
	}
}

func Select(db *sql.DB) map[string]string {
	rows, err := db.Query("select filename,code from task2 where expiretime > (select floor(extract(epoch from now())));")
	if err != nil {
		log.Fatal(err)
	}
	codes := Code{}
	nodemap := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&codes.Filename, &codes.Code)
		if err != nil {
			log.Fatal(err)
		}
		nodemap[codes.Filename] = codes.Code
	}
	return nodemap
}

func SelectOne(db *sql.DB, filename string) map[string]Code {
	rows, err := db.Query("select filename,expiretime,code from task2 where expiretime > (select floor(extract(epoch from now()))) and filename=$1;", filename)
	if err != nil {
		log.Fatal(err)
	}
	codes := Code{}
	nodemap := make(map[string]Code)
	for rows.Next() {
		err := rows.Scan(&codes.Filename, &codes.Expiretime, &codes.Code)
		if err != nil {
			log.Fatal(err)
		}
		nodemap[codes.Filename] = codes
	}
	return nodemap
}

// func UpdateUser(db *sql.DB) {
// 	stmt, err := db.Prepare("UPDATE  user_tbl  set username=$1 WHERE  id=$2")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec("Linux", 6)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("udpate user_tbl success")
// 	}
// }
