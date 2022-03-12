package DBS

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB
func Init(){
	var err error
	Db, err =sql.Open("mysql", "root:09070810.@tcp(localhost:3306)/ginHello")
	if err != nil{
		log.Panicln("err",err.Error())
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
}