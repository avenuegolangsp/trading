package domain

// Stock represents a stock asset
type Stock struct {
	Symbol      string  `json:"symbol"`
	Company     string  `json:"company"`
	Sector      string  `json:"sector"`
	MinPrice    float64 `json:"min_price"`
	MarketCap   string  `json:"market_cap"`
	Description string  `json:"description"`
}
