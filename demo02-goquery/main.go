package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://www.cnblogs.com/", nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	resp, _ := client.Do(request)
	reader, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()

	// 筛选class属性
	reader.Find(".post-item-title").Each(func(i int, selection *goquery.Selection) {
		// 拿文本
		text := selection.Text()
		fmt.Printf("text: %v\n", text)
		// 拿属性
		after, _ := selection.Attr("href")
		fmt.Printf("after: %v\n", after)
	})

}
