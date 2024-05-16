package domain

type CurrencyDescription string

func (c CurrencyDescription) IsValid() bool {
	if c != "" {
		return true
	}
	return false
}
func (c CurrencyDescription) ToString() string {
	return string(c)
}
func (c CurrencyDescription) ToBytes() []byte {
	return []byte(c)
}
