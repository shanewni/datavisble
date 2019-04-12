package main

import (
	"distributed/reptile/db"
	"distributed/reptile/debugs"
	"distributed/reptile/engine"
	"distributed/reptile/engine/contents"
	"fmt"
	"github.com/go-clog/clog"
	"os"
	"time"
)

func main() {
	//创建日志
	Clog()
	//存放字体破解文件数据
	//月票数被加密，使用破解文件进行破解
	fontCrack := make(map[string]map[string]string)
	for y := 2005; y <= time.Now().Year(); y++ {
		//按照年分组，最多存放12个月的集合
		authors := []engine.Author{}
		for m := 1; m <= 12; m++ {
			//判断数据是否是否已经最新，是则跳出循环
			if contents.JudjeTime(y, m) {
				break
			}
			//根据年月生成url
			url := "https://www.qidian.com/rank/yuepiao?month=" + fmt.Sprintf("%02d", m) + "&year=" + fmt.Sprintf("%d", y) + "&style=2&page="
			re := &engine.ReturnEnd{url + "1"}
			//获取当前根目录的子页数量
			end, errEnd := re.Start()
			if errEnd != nil {
				clog.Fatal(2, "Fail to get end : %v\n", errEnd)
			}
			if end == 0 {
				continue
			}
			//将当前年月的所有页面数据进行解析
			authors = append(authors, contents.JudjeView(end, y, m, url)...)
			if debugs.Debugs {
				clog.Trace("%s %v", engine.FontHead+authors[0].GetFontUrl()+engine.FontTail, authors)
			}
		}
		//解析真实月票数并添加到数据中
		contents.AddCount(authors, fontCrack)
		//插入到mysql中
		tt := time.Now()
		db.InsertSql(authors)
		clog.Trace("插入%d年数据用时%s\n", y, time.Now().Sub(tt).String())
	}

}

//创建日志
func Clog() {
	err := clog.New(clog.FILE, clog.FileConfig{
		Filename: "reptile.log",
	})
	if err != nil {
		fmt.Printf("Fail to create new logger: %v\n", err)
		os.Exit(1)
	}
}
