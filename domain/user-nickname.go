package domain

type NickName string

func (n NickName) IsValid() bool {
	if n != "" {
		return true
	}
	return false
}
func (n NickName) ToString() string {
	return string(n)
}
func (n NickName) ToBytes() []byte {
	return []byte(n)
}
