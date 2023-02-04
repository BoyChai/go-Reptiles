package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strconv"
	"strings"
)

func main() {
	url := "https://s.1688.com/selloffer/offer_search.htm?keywords=%CC%AB%D1%F4%C4%DC&n=y&netType=1%2C11%2C16&spm=a260k.dacugeneral.search.0"
	// 传给goquery处理
	//r, err := goquery.NewDocumentFromReader(getGoods(url))
	////r, err := goquery.NewDocumentFromReader(goodsLink(url))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//r.Find(".fui-paging-num").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//})
	fmt.Println(getGoodsPage(url))

}

// 拿到商品页面
func getGoods(url string) *strings.Reader {
	var resp []string
	var data string
	// 拿到网页数据
	fmt.Println(getGoodsPage(url))

	for i := 0; i < getGoodsPage(url); i++ {
		var d string
		ctx, _ := chromedp.NewContext(context.TODO())
		if err := chromedp.Run(ctx,
			chromedp.Navigate(fmt.Sprint(url, "&beginPage=", i+1)),
			chromedp.OuterHTML("html", &d, chromedp.ByQuery),
		); err != nil {
			log.Fatal(err)
		}
		resp[i] = d
	}
	for i := 0; i < len(resp); i++ {
		data = data + resp[i]
	}
	return strings.NewReader(data)
}

// 商品页数
func getGoodsPage(url string) int {
	ctx, cancel := chromedp.NewContext(context.TODO())
	defer cancel()
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
	var p string
	r.Find(".fui-paging-num").Each(func(i int, selection *goquery.Selection) {
		p = selection.Text()
	})
	page, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return page

}
