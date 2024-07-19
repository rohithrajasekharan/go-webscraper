package service

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/rohithrajasekharan/web-scraper/types"
)

func (h *Handler) createOrder(url string) ([]types.Product, error) {

	c := colly.NewCollector()
	products := []types.Product{}

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := types.Product{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img", "src"),
			Name:  e.ChildText("h2"),
			Price: e.ChildText(".price"),
		}
		products = append(products, product)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	c.Wait()

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found")
	}

	return products, nil
}
