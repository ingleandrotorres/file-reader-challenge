package domain

type Item struct {
	ID         string  `json:"id"`
	Site       string  `json:"site"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category"`
	CurrencyID string  `json:"currency_id"`
	SellerID   int     `json:"seller_id"`
	StarTime   string  `json:"date_created"`
}

func (i Item) IsValid() bool {
	if i.ID != "" && i.Site != "" && i.Price != 0 && i.CategoryID != "" && i.CurrencyID != "" && i.SellerID != 0 {
		return true
	}
	return false
}
