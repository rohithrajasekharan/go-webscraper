package types

type Product struct {
	Url   string `json:"url"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ScrapeRequest struct {
	Url string `json:"url"`
}
