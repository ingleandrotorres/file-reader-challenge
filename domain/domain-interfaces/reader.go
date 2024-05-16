package domain_interfaces

import "challenge-scrapy/domain"

type FileReader interface {
	Read(url, name, format, separator string) ([]domain.BlendID, error)
}

type SourceRepository interface {
	GetData() ([]domain.BlendID, error)
}
