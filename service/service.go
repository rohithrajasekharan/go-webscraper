package service

import (
	"github.com/gocolly/colly"
	"github.com/rohithrajasekharan/web-scraper/types"
)

func (h *Handler) createOrder(url string) ([]types.Product, error) {
	c := colly.NewCollector()
	c.Visit(url)
	products := []types.Product{}
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := types.Product{}

		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "src")
		product.Name = e.ChildText("h2")
		product.Price = e.ChildText(".price")

		products = append(products, product)
	})
	return products, nil
}
