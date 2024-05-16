package domain

//todo:corregir no es site
var validReaderType = map[string]Site{
	"csv":   Site("csv"),
	"space": Site("space"),
	"x":     Site("x"),
}

type ReaderType string

func (r ReaderType) IsValid() bool {

	if _, ok := validReaderType[string(r)]; ok {
		return true
	}
	return false
}
func (r ReaderType) String() string {
	return string(r)
}
