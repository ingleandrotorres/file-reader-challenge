package services

import (
	"challenge-scrapy/domain"
	"challenge-scrapy/domain/domain-interfaces"
	"challenge-scrapy/infrastructure"
	"errors"
)

type ReaderResolver struct {
}

func NewReaderResolver() ReaderResolver {
	return ReaderResolver{}
}

const (
	csv = "csv"
	txt = "txt"
)

func (t ReaderResolver) Select(readerType domain.ReaderType) (domain_interfaces.FileReader, error) {

	readerFunc := getAvailableReaders()[readerType]
	if readerFunc == nil {
		return nil, errors.New("invalid reader type")
	}

	return readerFunc(), nil
}

func getAvailableReaders() map[domain.ReaderType]func() domain_interfaces.FileReader {
	return map[domain.ReaderType]func() domain_interfaces.FileReader{
		"csv": getCsvReader(),
		"txt": getTxtReader(),
	}
}

func getCsvReader() func() domain_interfaces.FileReader {
	return func() domain_interfaces.FileReader {
		return infrastructure.NewCSVReader()
	}
}
func getTxtReader() func() domain_interfaces.FileReader {
	return func() domain_interfaces.FileReader {
		return infrastructure.NewTxtReader()
	}
}
