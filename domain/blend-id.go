package domain

type BlendID struct {
	Id   ID   `json:"id"`
	Site Site `json:"site"`
}

func NewBlendID(id ID, site Site) *BlendID {
	return &BlendID{
		Id:   id,
		Site: site,
	}
}
func (b BlendID) ToString() string {
	return string(b.Id) + string(b.Site)
}
