package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://gorm.io/zh_CN/docs/index.html", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "__gads=ID=1ff95293228eebb1:T=1664981362:S=ALNI_MapoWfwx-k68gOAOz0nx65W-_4X9Q; _ga_4CQQXWHK3C=GS1.1.1668576263.1.0.1668576273.0.0.0; _ga=GA1.2.857141578.1664981361; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1671520085,1672639828,1673058960,1673686500; __gpi=UID=00000a1ac4c9fbfc:T=1664981362:RT=1673686501:S=ALNI_MbrH9zh8AdT72gH6WpdzbYw8LmmnQ; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8GXQNXLgcs5PrnWvMs4xAGPLB0VrfoK2z3QuYEIS1BGGaoZv4-Q174CI8ozOhZaCZy3pIuqYpHdi1XwtF2z1xl4knAqbHuCli4frenxROiwI2_4lu9U7F0BfODcWE_aFSR4XhjO_vJBFEfdOjik5xxw")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode != 200 {
		log.Println("Status Code Error: ", resp.StatusCode)
	}
	value, err := io.ReadAll(resp.Body)
	defer func() { _ = resp.Body.Close() }()
	fmt.Println(string(value))
}
