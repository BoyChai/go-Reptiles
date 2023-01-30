package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("#pl_top_realtimehot tbody .td-02 a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		request.Headers.Set("Cookie", "SUB=_2AkMUzPfgf8NxqwFRmP0QyGrha45zywjEieKikAY7JRMxHRl-yT9kqmMitRB6P0zZD0Pn_miqx0LR8GFBDShFYrynEzoC; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WWY8ac1N0I98cU_boYXm99P; SINAGLOBAL=2947613495381.2007.1670412505930; _s_tentry=www.google.com; UOR=,,www.google.com; Apache=4338821339332.2837.1674976440626; ULV=1674976440629:4:2:1:4338821339332.2837.1674976440626:1673056507776")
	})
	c.Visit("https://s.weibo.com/top/summary?cate=realtimehot")
	//
	//client := &http.Client{}
	//request, _ := http.NewRequest("GET", "https://s.weibo.com/top/summary?cate=realtimehot", nil)
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	//request.Header.Set("Cookie", "SUB=_2AkMUzPfgf8NxqwFRmP0QyGrha45zywjEieKikAY7JRMxHRl-yT9kqmMitRB6P0zZD0Pn_miqx0LR8GFBDShFYrynEzoC; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WWY8ac1N0I98cU_boYXm99P; SINAGLOBAL=2947613495381.2007.1670412505930; _s_tentry=www.google.com; UOR=,,www.google.com; Apache=4338821339332.2837.1674976440626; ULV=1674976440629:4:2:1:4338821339332.2837.1674976440626:1673056507776")
	//resp, _ := client.Do(request)
	//body, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(body))
}
