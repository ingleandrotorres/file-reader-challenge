package domain

type ID string

func (i ID) IsValid() bool {
	if i != "" {
		return true
	}
	return false
}
