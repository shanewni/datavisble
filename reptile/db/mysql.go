package db

import (
	"database/sql"
	"distributed/reptile/engine"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//插入数据
func InsertSql(authors []engine.Author) {
	db, err := sql.Open("mysql", "root:52172d++@tcp(127.0.0.1:3306)/?charset=utf8")
	defer db.Close()
	if err != nil {
		fmt.Println("open sql error: ", err)
	}
	value := []string{}
	for _, v := range authors {
		value = append(value, fmt.Sprintf("(%d,%s,%s,%s,%d,%d)", v.GetUid(), "'"+v.GetType()+"'", "'"+v.GetName()+"'", "'"+v.GetAuthor()+"'", v.GetCount(), v.GetTime()))
	}
	v := strings.Join(value, ",")
	//开启事务
	tx, _ := db.Begin()
	_, errInsert := tx.Query(`INSERT ippool.author VALUE ` + v)
	if errInsert != nil {
		tx.Rollback()
		fmt.Printf("INSERT author error: ", errInsert)
	} else {
		tx.Commit()
	}
}
