package domain

var validSites = map[string]Site{
	"MLA": Site("MLA"),
	"MLB": Site("MLB"),
	"MCO": Site("MCO"),
}

type Site string

func (i Site) IsValid() bool {

	if _, ok := validSites[string(i)]; ok {
		return true
	}
	return false
}
