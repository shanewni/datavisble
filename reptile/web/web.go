package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-clog/clog"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

type author struct {
	Uid    int    `json:"uid"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Count  int    `json:"count"`
	Stamp  int64  `json:"stamp"`
}

func main() {
	http.HandleFunc("/", a)
	if err := http.ListenAndServe(":3333", nil); err != nil {
		clog.Fatal(2, "Fail to ListenAndServe : %v\n", err)
		//fmt.Println("ListenAndServe:", err)
	}
}

//实现数据接口
func a(w http.ResponseWriter, r *http.Request) {
	//允许跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//获取参数
	r.ParseForm()
	if len(r.Form["year"]) > 0 || len(r.Form["mouth"]) > 0{
		db, _ := sql.Open("mysql", "root:52172d++@tcp(127.0.0.1:3306)/?charset=utf8")
		defer db.Close()

		authors := []author{}
		times := "2006-01-02 15:04"
		t, _ := time.Parse(times,fmt.Sprintf("%s-%s-%s %s",r.Form["year"][0],fmt.Sprintf("%02s",r.Form["mouth"][0]),"01","00:00"))
		row, _ := db.Query("SELECT * FROM ippool.author WHERE time = "+fmt.Sprintf("%d",t.Unix()))
		for row.Next() {
			a := author{}
			row.Scan(&a.Uid, &a.Type, &a.Name, &a.Author, &a.Count, &a.Stamp)
			authors = append(authors, a)
		}

		bytes, _ := json.Marshal(authors)
		fmt.Fprint(w, string(bytes))
	} else {
		fmt.Fprint(w, "获取get请求参数出错")
		return
	}
}
