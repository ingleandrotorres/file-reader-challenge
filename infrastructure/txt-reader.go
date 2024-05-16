package infrastructure

import (
	"challenge-scrapy/domain"
	"encoding/csv"
	"fmt"
	"os"
)

type TxtReader struct{}

func NewTxtReader() TxtReader {
	return TxtReader{}
}
func (r TxtReader) Read(url, name, format, separator string) ([]domain.BlendID, error) {

	nameFile := fmt.Sprintf("%s%s.%s", url, name, format)
	file, err := os.Open(nameFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' //aqui separator.

	var blendID []domain.BlendID

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		blendID = append(blendID, domain.BlendID{
			Id:   domain.ID(record[0]),
			Site: domain.Site(record[1])})
	}
	return blendID, nil
}
