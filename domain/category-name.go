package domain

type CategoryName string

func (c CategoryName) ToString() string {
	return string(c)
}

func (c CategoryName) IsValid() bool {
	if c != "" {
		return true
	}
	return false
}
func (c CategoryName) ToBytes() []byte {
	return []byte(c)
}
