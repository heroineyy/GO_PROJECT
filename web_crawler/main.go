package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"time"
)

//1、请求
//2、解析
//3、打印
func main() {
	//获取客户端
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	//根据url的规律，for循环获取10张页面
	for i := 0; i < 10; i++ {
		index := i * 25
		//strconv.Itoa(index)将整型转换成string型
		req, _ := http.NewRequest("GET", "https://movie.douban.com/top250?start="+strconv.Itoa(index)+"&filter=", nil)

		//User-Agent
		//即用户代理，网站服务器通过识别它来确定用户所使用的操作系统版本、CPU 类型、浏览器版本等信息，并判断通过判断它来给客户端发送不同的页面。
		//在返回值的头部设置user-agent
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
		//发起 HTTP 请求
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		defer resp.Body.Close()

		//resp.Body就是html文件
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//#content 表示id为content
		//.grid_view 表示class为grid_view
		doc.Find("#content .grid_view li").Each(func(i int, s *goquery.Selection) {
			content := s.Find(".title").Eq(0).Text()
			fmt.Println(content)
		})
	}

}
