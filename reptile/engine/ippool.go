package engine

import (
	"crypto/tls"
	"distributed/reptile/debugs"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//使用ip池访问url
func Ippool(curl string) (*goquery.Document, bool){
	//容错处理
	pass := true
	defer func() {
		if err := recover();err !=nil{
			if debugs.Debugs{
				fmt.Println("null ip：",err)
			}
			pass = false
		}
	}()

	ip := GetEffectiveIp()
	if ip ==""{
		doc,_ := goquery.NewDocument(curl)
		return doc,true
	}
	proxy, _ := url.Parse("http://"+ip)
	fmt.Println(proxy)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 3,
	}

	res, _ := client.Get(curl)
	doc, err := goquery.NewDocumentFromResponse(res)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	if debugs.Debugs {
		fmt.Println(ip,curl)
	}

	return doc, pass
}

//通过接口获取代理ip
func GetEffectiveIp()string{
	resp, err := http.Get("http://localhost:8080/v2/ip")
	if err !=nil{
		//容错机制，无代理ip使用自身ip
		return ""
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	a := make(map[string]string)
	json.Unmarshal(b, &a)
	ip := a["ip"]
	return ip
}