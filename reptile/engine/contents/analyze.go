package contents

import (
	"distributed/reptile/debugs"
	"distributed/reptile/engine"
	"github.com/go-clog/clog"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//根据html页面进行解析
func JudjeView(end, y, m int, url string) []engine.Author {
	var authors []engine.Author
	tt := time.Now()
	//按照月度排行榜进行记录
	sort := 0
	for page := 1; page <= end; page++ {
		doc, errPage := goquery.NewDocument(url + fmt.Sprintf("%d", page))
		if errPage != nil {
			clog.Fatal(2, "Fail to goquery.NewDocument(url+page) error : %v\n", errPage)
			//fmt.Println("goquery.NewDocument(url+page) error", errPage)
		}
		doc.Find(".rank-table-list").Find("tbody").Find("tr").Each(func(i int, selection *goquery.Selection) {
			author := engine.Author{}
			//添加uid
			sort += 1
			author.AddUid(sort)

			//16进制月票编码，待破解
			sexUnicode := fmt.Sprintf("%x", []rune(selection.Find(".month").Find("span").Text()))
			author.AddSexUnicode(sexUnicode)

			//字体文件名称
			getUrl := `url\(\'https://qidian.gtimg.com/qd_anti_spider/([^.]+).ttf\'\)`
			reg := regexp.MustCompile(getUrl)
			url := reg.FindStringSubmatch(selection.Find(".month").Find("style").Text())
			author.AddFontUrl(url[1])

			//小说类型
			//utf-8汉字占3个字节，切片保留两个中文
			bookType := selection.Find(".type").Text()[3:9]
			author.AddType(bookType)

			//小说名称
			bookName := selection.Find(".name").Text()
			author.AddName(bookName)

			//小说作者
			aut := selection.Find(".author").Text()
			author.AddAuthor(aut)

			//所在时间戳
			old, _ := time.Parse("2006-01", fmt.Sprintf("%02d-%02d", y, m))
			author.AddTime(old.Unix())

			if debugs.Debugs {
				clog.Trace("%v", author)
			}

			authors = append(authors, author)
		})
		//防封
		time.Sleep(time.Second * 2)
	}
	clog.Trace("%d年%02d月数据以爬完,共%d页，用时%s,ip防封停顿%d秒\n", y, m, end, time.Now().Sub(tt).String(), end*2)
	return authors
}

//将真实月票数放入结构体中
func AddCount(authors []engine.Author, fontCrack map[string]map[string]string) {
	for i := 0; i < len(authors); i++ {
		font := authors[i].GetFontUrl()

		//下载不存在的破解文件
		if fontCrack[font] == nil {
			//下载字体文件
			authors[i].CreateFontFile()
			//编译字体文件
			authors[i].BuildFontFile()
			//提取破解内容
			fontCrack[font] = authors[i].CraCkXML()
		}

		sexs := strings.Split(authors[i].GetSexCount()[1:len(authors[i].GetSexCount())-1], " ")
		count := ""
		for i := 0; i < len(sexs); i++ {
			key := fontCrack[font][sexs[i]]
			strCount := engine.Analyze[key]
			count += strCount
		}
		countNum, _ := strconv.Atoi(count)
		authors[i].AddCount(countNum)
	}
}

//判断数据url是否最新
func JudjeTime(y, m int) bool {
	old, _ := time.Parse("2006-01", fmt.Sprintf("%02d-%02d", y, m))
	new, _ := time.Parse("2006-01", fmt.Sprintf("%02d-%02d", time.Now().Year(), time.Now().Month()))
	if old.Unix() > new.Unix() {
		clog.Trace("爬取数据结束")
		return true
	} else {
		return false
	}
}
