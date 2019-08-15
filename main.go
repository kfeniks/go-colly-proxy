package main

import (
//	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks4://51.158.68.26:8811", "socks5://51.158.68.26:8811")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// On every a element which has href attribute call callback
	c.OnHTML("#offers_table", func(e *colly.HTMLElement) {
		// Print link
		// motoList := e.Text

		e.ForEach("table", func(someInt int, e *colly.HTMLElement) {

			// motoList := e.Text


			fmt.Println(someInt)

			//	fmt.Printf("text found: %q -> %s\n", motoList, motoList)
		})

	//	fmt.Println(e.Text)

	//	fmt.Printf("text found: %q -> %s\n", motoList, motoList)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(e))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Print the response
	c.OnResponse(func(r *colly.Response) {
	//	log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.olx.ua/transport/moto/")
}