package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {

	})
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("1")
		fmt.Println(request.URL)
	})
	//c.OnHTML("#sm-offer-list .normalcommon-offer-card ad-item a", func(element *colly.HTMLElement) {
	c.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println("2")
		fmt.Println(element.Text)
	})

	c.Visit("https://s.1688.com/selloffer/offer_search.htm?keywords=%CC%AB%D1%F4%C4%DC&n=y&netType=1%2C11%2C16&spm=a260k.dacugeneral.search.0")
}
