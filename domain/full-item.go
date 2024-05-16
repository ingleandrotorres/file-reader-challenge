package domain

type FullItem struct {
	Id          string  `json:"id"`
	Site        string  `json:"site"`
	Price       float64 `json:"price"`
	StartTime   string  `json:"start_time"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nickname    string  `json:"nickname"`
	Currency    string  `json:"currency"`
}

func ToBytes() {

}
