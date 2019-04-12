package engine

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-clog/clog"
	"strconv"
)

type Option interface {
	Start()(int, error)
}

type ReturnEnd struct {
	Url string
}

//获取当前页数最高值
func (r *ReturnEnd) Start() (int, error) {
	docStart, errStart := goquery.NewDocument(r.Url)
	if errStart != nil {
		clog.Fatal(2, "open root error:", errStart)
		//fmt.Println("open root error：", errStart)
	}
	if end, ok := docStart.Find("#page-container").Attr("data-pagemax"); ok {
		return strconv.Atoi(end)
	}
	return 0, fmt.Errorf("get end error")
}
