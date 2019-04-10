package engine

import (
	"distributed/reptile/debugs"
	"fmt"
	"github.com/donnie4w/dom4g"
	"io"
	"net/http"
	"os"
	"os/exec"
)

var FontHead = "https://qidian.gtimg.com/qd_anti_spider/"
var FontTail = ".ttf"

var Analyze = map[string]string{
	"one":    "1",
	"two":    "2",
	"three":  "3",
	"four":   "4",
	"five":   "5",
	"six":    "6",
	"seven":  "7",
	"eight":  "8",
	"nine":   "9",
	"zero":   "0",
	"period": ".",
}

type GetAuthor interface {
	AddUid(id int)
	AddType(t string)
	AddName(name string)
	AddAuthor(aut string)
	AddCount(count int)
	AddTime(tm int64)
	AddSexUnicode(sex string)
	AddFontUrl(font string)
	GetFontUrl() string
	GetSexCount() string
	GetUid() int
	GetType() string
	GetName() string
	GetAuthor() string
	GetCount() int
	GetTime() int64
	CreateFontFile()
	BuildFontFile()
	CraCkXML() map[string]string
}

type Author struct {
	uid          int
	bookType     string
	bookName     string
	bookAuthor   string
	bookCount    int
	timestamp    int64
	unicodeSex   string
	ttfFontStyle string
}

//添加小说排名
func (a *Author) AddUid(id int) {
	a.uid = id
}

//添加小说类型
func (a *Author) AddType(t string) {
	a.bookType = t
}

//添加小说名字
func (a *Author) AddName(name string) {
	a.bookName = name
}

//添加小说作者昵称
func (a *Author) AddAuthor(aut string) {
	a.bookAuthor = aut
}

//添加小说月票数
func (a *Author) AddCount(count int) {
	a.bookCount = count
}

//添加月票书所在时间
func (a *Author) AddTime(tm int64) {
	a.timestamp = tm
}

//添加月票数的16进制/为破解准备
func (a *Author) AddSexUnicode(sex string) {
	a.unicodeSex = sex
}

//添加月票数的破解文件url
func (a *Author) AddFontUrl(font string) {
	a.ttfFontStyle = font
}

//获取字体格式的破解文件url
func (a *Author) GetFontUrl() string {
	return a.ttfFontStyle
}

//16进制月票数
func (a *Author) GetSexCount() string {
	return a.unicodeSex
}

//获取小说排名
func (a *Author) GetUid() int {
	return a.uid
}

//获取小说类型
func (a *Author) GetType() string {
	return a.bookType
}

//获取小说名字
func (a *Author) GetName() string {
	return a.bookName
}

//获取小说作者昵称
func (a *Author) GetAuthor() string {
	return a.bookAuthor
}

//获取小说月票数
func (a *Author) GetCount() int {
	return a.bookCount
}

//获取月票数所在时间
func (a *Author) GetTime() int64 {
	return a.timestamp
}

//下载破解字体的文件
func (a *Author) CreateFontFile() {
	res, errGet := http.Get(FontHead + a.ttfFontStyle + FontTail)
	defer res.Body.Close()
	if errGet != nil {
		fmt.Println("Get font error", errGet)
	}
	f, errCreate := os.Create("./reptile/crack/" + a.ttfFontStyle + FontTail)
	if errCreate != nil {
		fmt.Println("Create error", errCreate)
		return
	}
	io.Copy(f, res.Body)
}

//编译字体文件，供破解使用
func (a *Author) BuildFontFile() {
	cmd := exec.Command("ttx", "./reptile/crack/"+a.ttfFontStyle+FontTail)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	if debugs.Debugs {
		fmt.Println(string(out))
	}

}

//破解字体文件xml
func (a *Author) CraCkXML() map[string]string {
	table := make(map[string]string)
	file, errCrackXml := os.Open("./reptile/crack/" + a.ttfFontStyle + ".ttx")
	if errCrackXml != nil {
		fmt.Println(errCrackXml)
	}
	ele, _ := dom4g.LoadByStream(file)
	eles := ele.Node("cmap").Node("cmap_format_12").Nodes("map")

	for _, v := range eles {
		code, _ := v.AttrValue("code")
		name, _ := v.AttrValue("name")
		table[code[2:]] = name
	}
	return table
}
