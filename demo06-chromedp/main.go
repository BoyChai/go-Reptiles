package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

func main() {
	//// 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
	//ctx, cancel := chromedp.NewContext(
	//	context.Background(),
	//	chromedp.WithLogf(log.Printf),
	//)
	//defer cancel()
	//var nodes []*cdp.Node
	//err := chromedp.Run(ctx,
	//	chromedp.Navigate("https://www.cnblogs.com/"),
	//	chromedp.WaitVisible(`#footer`, chromedp.ByID),
	//	chromedp.Nodes(`.//a[@class="titlelnk"]`, &nodes),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("get nodes:", len(nodes))
	//// print titles
	//for _, node := range nodes {
	//	fmt.Println(node.Children[0].NodeValue, ":", node.AttributeValue("href"))
	//}
	//ctx, cancel := chromedp.NewContext(context.Background()) // chromedp.WithDebugf(log.Printf), 是否显示调试信息
	//
	//defer cancel()
	//
	//// 设置超时
	//ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	//defer cancel()
	//// 导航到页面，等待元素，单机
	//var example string
	//err := chromedp.Run(ctx,
	//	chromedp.Navigate(`https://pkg.go.dev/time`),
	//	chromedp.WaitVisible(`body > footer`),
	//	chromedp.Click(`#example-After`, chromedp.NodeVisible),
	//	chromedp.Value(`#example-After textarea`, &example),
	//)
	//err := chromedp.Run(ctx,
	//	chromedp.Navigate(`https://pkg.go.dev/time`),
	//	//chromedp.Navigate(`https://acz.youku.com/wow/ykpage/act/top_hot?spm=a2hja.14919748_WEBHOME_NEW.search.1`),
	//	chromedp.WaitVisible(`body #pkg-constants`),
	//	chromedp.Text(`#pkg-constants`, &example),
	//)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("数据" + example)/
	//http://127.0.0.1:8848/shop_demo/aa.html
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := "https://acz.youku.com/wow/ykpage/act/top_hot?spm=a2hja.14919748_WEBHOME_NEW.search.1"
	var data string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	r, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	r.Find("#root > div > div.list-box > div > div:nth-child(1) .item").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		fmt.Println(text)
	})
	//ctx, cancel := chromedp.NewContext(context.Background())
	//defer cancel()
	//url := "http://127.0.0.1:8848/shop_demo/aa.html"
	//var data string
	//if err := chromedp.Run(ctx,
	//	chromedp.Navigate(url),
	//	//chromedp.Value(`#test`, &data),
	//	chromedp.OuterHTML("#test", &data, chromedp.ByQuery),
	//); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(data)

}
