package getter

import (
	"distributed/proxypool/pkg/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-clog/clog"
	"strconv"
	"strings"
)

func Xila() (result []*models.IP) {
	for i :=1; i <= 15;i++{
		pollURL,_ := goquery.NewDocument("http://www.xiladaili.com/https/"+strconv.Itoa(i)+"/")
		pollURL.Find(".fl-table").Find("tbody").Find("tr").Each(func(i int, selection *goquery.Selection) {
			ret,_:= selection.Find("td").Html()
			ipport := strings.Split(ret," ")
			IP := models.NewIP()
			IP.Data = ipport[0]
			IP.Type1 = "http"
			//IP.Type2 = "https"
			result = append(result, IP)
			clog.Info("[kxila] done")
		})
	}
	return
}