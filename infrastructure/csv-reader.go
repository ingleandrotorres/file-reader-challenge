package infrastructure

import (
	"challenge-scrapy/domain"
	"encoding/csv"
	"fmt"
	"os"
)

type CSVReader struct{}

func NewCSVReader() CSVReader {
	return CSVReader{}
}
func (r CSVReader) Read(url, name, format, separator string) ([]domain.BlendID, error) {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	nameFile := fmt.Sprintf("%s%s.%s", url, name, format)
	file, err := os.Open(pwd + "/resources/" + nameFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ',' //aqui separator.

	var blendID []domain.BlendID

	count := 0
	for {
		count++
		if count == 1 {
			continue
		}
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
