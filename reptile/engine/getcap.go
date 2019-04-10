//获取读者位置代码未完成
package engine

import (
	"distributed/reptile/debugs"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

var head = "https:"
//评论区后缀
var tail = "?type=1&page="

type GetCap interface {
	GetImgUrl(docBook *goquery.Document)
	GetBookName(docBook *goquery.Document)
	GetAuthor(docBook *goquery.Document)
	GetTypes(docBook *goquery.Document)
	GetComDoc(docBook *goquery.Document) (string, int)
	GetCap()
}

type City struct {
	A海门   int
	A鄂尔多斯 int
	A招远   int
	A舟山   int
	A齐齐哈尔 int
	A盐城   int
	A赤峰   int
	A青岛   int
	A乳山   int
	A金昌   int
	A泉州   int
	A莱西   int
	A日照   int
	A胶南   int
	A南通   int
	A拉萨   int
	A云浮   int
	A梅州   int
	A文登   int
	A上海   int
	A攀枝花  int
	A威海   int
	A承德   int
	A厦门   int
	A汕尾   int
	A潮州   int
	A丹东   int
	A太仓   int
	A曲靖   int
	A烟台   int
	A福州   int
	A瓦房店  int
	A即墨   int
	A抚顺   int
	A玉溪   int
	A张家口  int
	A阳泉   int
	A莱州   int
	A湖州   int
	A汕头   int
	A昆山   int
	A宁波   int
	A湛江   int
	A揭阳   int
	A荣成   int
	A连云港  int
	A葫芦岛  int
	A常熟   int
	A东莞   int
	A河源   int
	A淮安   int
	A泰州   int
	A南宁   int
	A营口   int
	A惠州   int
	A江阴   int
	A蓬莱   int
	A韶关   int
	A嘉峪关  int
	A广州   int
	A延安   int
	A太原   int
	A清远   int
	A中山   int
	A昆明   int
	A寿光   int
	A盘锦   int
	A长治   int
	A深圳   int
	A珠海   int
	A宿迁   int
	A咸阳   int
	A铜川   int
	A平度   int
	A佛山   int
	A海口   int
	A江门   int
	A章丘   int
	A肇庆   int
	A大连   int
	A临汾   int
	A吴江   int
	A石嘴山  int
	A沈阳   int
	A苏州   int
	A茂名   int
	A嘉兴   int
	A长春   int
	A胶州   int
	A银川   int
	A张家港  int
	A三门峡  int
	A锦州   int
	A南昌   int
	A柳州   int
	A三亚   int
	A自贡   int
	A吉林   int
	A阳江   int
	A泸州   int
	A西宁   int
	A宜宾   int
	A呼和浩特 int
	A成都   int
	A大同   int
	A镇江   int
	A桂林   int
	A张家界  int
	A宜兴   int
	A北海   int
	A西安   int
	A金坛   int
	A东营   int
	A牡丹江  int
	A遵义   int
	A绍兴   int
	A扬州   int
	A常州   int
	A潍坊   int
	A重庆   int
	A台州   int
	A南京   int
	A滨州   int
	A贵阳   int
	A无锡   int
	A本溪   int
	A克拉玛依 int
	A渭南   int
	A马鞍山  int
	A宝鸡   int
	A焦作   int
	A句容   int
	A北京   int
	A徐州   int
	A衡水   int
	A包头   int
	A绵阳   int
	A乌鲁木齐 int
	A枣庄   int
	A杭州   int
	A淄博   int
	A鞍山   int
	A溧阳   int
	A库尔勒  int
	A安阳   int
	A开封   int
	A济南   int
	A德阳   int
	A温州   int
	A九江   int
	A邯郸   int
	A临安   int
	A兰州   int
	A沧州   int
	A临沂   int
	A南充   int
	A天津   int
	A富阳   int
	A泰安   int
	A诸暨   int
	A郑州   int
	A哈尔滨  int
	A聊城   int
	A芜湖   int
	A唐山   int
	A平顶山  int
	A邢台   int
	A德州   int
	A济宁   int
	A荆州   int
	A宜昌   int
	A义乌   int
	A丽水   int
	A洛阳   int
	A秦皇岛  int
	A株洲   int
	A石家庄  int
	A莱芜   int
	A常德   int
	A保定   int
	A湘潭   int
	A金华   int
	A岳阳   int
	A长沙   int
	A衢州   int
	A廊坊   int
	A菏泽   int
	A合肥   int
	A武汉   int
	A大庆   int
}

type Book struct {
	Name   string
	Img    string
	Author string
	Types  string
	Sex    int
	City
}

//根据html获取图片链接保存到结构体中
func (b *Book) GetImgUrl(docBook *goquery.Document) {
	if imgUrl, ok := docBook.Find(".book-img").Find("img").Attr("src"); ok {
		if debugs.Debugs {
			fmt.Printf("book.Img:%s", head+imgUrl)
		}
		b.Img = head + imgUrl[:len(imgUrl)-1]
	}
}

//根据html获取图书名保存到结构体中
func (b *Book) GetBookName(docBook *goquery.Document) {
	bookName, errEm := docBook.Find(".book-info").Find("em").Html()
	if errEm != nil {
		fmt.Println(errEm)
	}
	if debugs.Debugs {
		fmt.Println("book.Name:", bookName)
	}
	b.Name = bookName
}

//根据html获取作者昵称保存到结构体中
func (b *Book) GetAuthor(docBook *goquery.Document) {
	author, errA := docBook.Find(".book-info").Find("span").Find("a").Html()
	if errA != nil {
		fmt.Println(errA)
	}
	if debugs.Debugs {
		fmt.Println("book.Author:", author)
	}
	b.Author = author
}

//根据html获取书籍类型保存到结构体中
func (b *Book) GetTypes(docBook *goquery.Document) {
	types, errA := docBook.Find(".book-info").Find(".tag").Find("a").Html()
	if errA != nil {
		fmt.Println(errA)
	}
	if debugs.Debugs {
		fmt.Println("book.Types:", types)
	}
	b.Types = types
}

//获取每本书的评论链接及页数
func (b *Book) GetComDoc(docBook *goquery.Document) (string, int) {
	if comUrl, ok := docBook.Find(".content-nav-wrap").Find(".j_discussion_block").Find("a").Attr("href"); ok {
		//获取重定向url，修改golang源码获取
		resp, _ := http.Get(head + comUrl)
		defer resp.Body.Close()
		//根据重定向后的url解析
		docCom, errCom := goquery.NewDocument(head + http.RedirectMy)
		if errCom != nil {
			fmt.Println("get comment error: ", errCom)
		}
		var end int
		if e, ok := docCom.Find(".page-container").Attr("data-pagemax"); ok {
			end, _ = strconv.Atoi(e)
		}
		if debugs.Debugs {
			fmt.Println("url and count: ", head+http.RedirectMy, end)
		}
		return http.RedirectMy, end
	} else {
		fmt.Println("open comment error")
		return "", 0
	}
}

//获取读者分布位置
//此代码未完成
func (b *Book)GetCap() {
	//根据页数循环
	e := ReturnEnd{"https://www.qidian.com/all"}
	end,_ := e.Start()
	for page := 1; page <= end; page++ {
		doc, err := goquery.NewDocument("https://www.qidian.com/all?orderId=&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0&page=" + strconv.Itoa(page))
		if err != nil {
			fmt.Println("open url error", err)
		}

		doc.Find(".book-img-box").Each(func(i int, contentSelection *goquery.Selection) {
			title, _ := contentSelection.Find("a").Attr("href")
			//根据正确的url寻址
			if debugs.Debugs {
				fmt.Println("book url", head+title)
			}
			docBook, err := goquery.NewDocument(head + title)
			if err != nil {
				fmt.Println("open "+title+" error", err)
			}
			b := &Book{}
			b.GetImgUrl(docBook)
			b.GetBookName(docBook)
			b.GetAuthor(docBook)
			b.GetTypes(docBook)

			comUrl, urlEnd := b.GetComDoc(docBook)
			for pard := 1; pard <= urlEnd; pard++ {
				comDoc, errCom := goquery.NewDocument(head + comUrl + tail + strconv.Itoa(pard))
				if errCom != nil {
					fmt.Println("open comment error: ", errCom)
				}
				if debugs.Debugs {
					fmt.Println(head + comUrl + tail + strconv.Itoa(pard))
				}
				comDoc.Find(".all-post").Find(".post-wrap").Each(func(i int, selection *goquery.Selection) {
					if userurl, ok := selection.Find(".card-wrap").Find("a").Attr("href"); ok {
						//使用ip池
						ComUser := &goquery.Document{}
						for {
							if ComUser, ok = Ippool(head + userurl); ok {
								break
							}
						}
						if ComUser != nil {
							fmt.Println(ComUser.Find(".header-msg-desc").Html())
						}
					}
				})
			}
		})
	}
}
