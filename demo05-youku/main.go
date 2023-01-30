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
	// 封装一个chromedp的context
	ctx, cancel := chromedp.NewContext(context.TODO())
	defer cancel()
	url := "https://acz.youku.com/wow/ykpage/act/top_hot?spm=a2hja.14919748_WEBHOME_NEW.search.1"
	var data string
	// 拿到网页数据
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}
	// 传给goquery处理
	r, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	// goquery进行处理
	r.Find("#root > div > div.list-box > div > div:nth-child(1) .item").Each(func(i int, selection *goquery.Selection) {
		// 热门排行
		fmt.Println(selection.Text()[1:])
		// 对应链接
		//fmt.Println(selection.Attr("href"))
	})
}
