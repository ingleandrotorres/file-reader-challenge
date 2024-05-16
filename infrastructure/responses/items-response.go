package responses

// https://api.mercadolibre.com/items?ids=MLA750925229,,MLA594239600&attributes=price,date_created,category_id,currency_id,seller_id
type ItemsResponse struct {
	Code int `json:"code"`
	Body struct {
		Id          string  `json:"id"`           //get this
		SellerId    int     `json:"seller_id"`    //get this
		CategoryId  string  `json:"category_id"`  //get this
		Price       float64 `json:"price"`        //get this
		CurrencyId  string  `json:"currency_id"`  //get this
		DateCreated string  `json:"date_created"` //start_time?
	} `json:"body"`
}
