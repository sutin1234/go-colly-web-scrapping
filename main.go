package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type StockQuote struct {
	StockTitle      string
	StockUnit       string
	GrowthTo        string
	PercentGrowthTo string
}

func main() {
	quotes := []StockQuote{}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/115.0")
		fmt.Printf("Visiting: %v\n", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response StatusCode: %v\n", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("onError: %v\n", err.Error())
	})
	c.OnHTML("div.quote-info", func(h *colly.HTMLElement) {

		div := h.DOM
		stockTitle := strings.TrimSpace(div.Find(".symbol").Text())
		stockUnit := strings.TrimSpace(div.Find(".symbol-title").Text())
		growthTo := strings.TrimSpace(div.Find(".value").Text())
		percentGrowthTo := div.Find(".theme-success").Text()

		quote := StockQuote{
			StockTitle:      stockTitle,
			StockUnit:       stockUnit,
			GrowthTo:        growthTo,
			PercentGrowthTo: percentGrowthTo,
		}
		quotes = append(quotes, quote)

		fmt.Println("Report KTB Stock in SET")

		for _, q := range quotes {
			fmt.Printf("StockTitle: %v\n", q.StockTitle)
			fmt.Printf("StockUnit: %v\n", q.StockUnit)
			fmt.Printf("GrowthTo: %v\n", q.GrowthTo)
			fmt.Printf("PercentGrowthTo: %v\n", q.PercentGrowthTo)
		}
	})

	err := c.Visit("https://www.set.or.th/th/market/product/stock/quote/KTB/price")
	if err != nil {
		panic(err.Error())

	}
}
