package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// 创建连接器
	c := colly.NewCollector()
	// 抓取到html页面的时候会调用,前面的字符实际是个标签选择器
	c.OnHTML(".post-item-title", func(e *colly.HTMLElement) {
		// 拿到父标签
		fmt.Println(e.Name)
		// 拿到文本内容
		fmt.Println(e.Text)
		// 拿到整个链接
		url := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Println(url)
	})
	// 发送请求会调用
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("url: ", request.URL)
	})

	// 发送请求
	c.Visit("https://www.cnblogs.com/")
}
