package sqlconect

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"

	"log"
)

var SqlDB *sql.DB

func init() {
	db, err := sql.Open("mysql", "web:Lomoz1986@/sitemanager?parseTime=true")
	if err != nil {
		log.Println(err)
	}
	SqlDB = db
}
