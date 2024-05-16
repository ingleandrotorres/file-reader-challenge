package domain_interfaces

import "challenge-scrapy/infrastructure/responses"

type Gateway interface {
	GetItems(items []string) ([]responses.ItemsResponse, error)
}
