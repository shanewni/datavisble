package engine

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

//单元测试
func Test_GetImgUrl(t *testing.T) {
	bk := &Book{}
	doc, err := goquery.NewDocument("https://www.qidian.com/all?orderId=&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0&page=1")
	if err != nil {
		fmt.Println("open url error", err)
	}

	doc.Find(".book-img-box").Each(func(i int, contentSelection *goquery.Selection) {
		title, _ := contentSelection.Find("a").Attr("href")
		//根据正确的ril寻址
		docBook, err := goquery.NewDocument(head + title)
		if err != nil {
			fmt.Println("open "+title+" error", err)
		}

		bk.GetImgUrl(docBook)
		bk.GetBookName(docBook)
		bk.GetAuthor(docBook)
		bk.GetTypes(docBook)
		if bk.Img == "" || bk.Name == "" || bk.Author == "" || bk.Types == "" {
			t.Error("error: struct is null")
		}
	})
}

//性能测试
func BenchmarkGetImgUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bk := &Book{}
		doc, err := goquery.NewDocument("https://www.qidian.com/all?orderId=&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0&page=1")
		if err != nil {
			fmt.Println("open url error", err)
		}

		doc.Find(".book-img-box").Each(func(i int, contentSelection *goquery.Selection) {
			title, _ := contentSelection.Find("a").Attr("href")
			//根据正确的ril寻址
			docBook, err := goquery.NewDocument(head + title)
			if err != nil {
				fmt.Println("open "+title+" error", err)
			}

			bk.GetImgUrl(docBook)
			bk.GetBookName(docBook)
			bk.GetAuthor(docBook)
			bk.GetTypes(docBook)
		})
	}
}
