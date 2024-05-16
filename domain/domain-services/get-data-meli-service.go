package domain_services

import (
	"challenge-scrapy/domain"
	"challenge-scrapy/domain/domain-interfaces"
	"context"
)

type GetDataBusinessService struct {
	service     domain_interfaces.DataBusinessGetter
	authManager *domain.AuthManager
}

func NewGetDataBusinessService(
	service domain_interfaces.DataBusinessGetter,
	auth *domain.AuthManager) *GetDataBusinessService {
	return &GetDataBusinessService{
		service:     service,
		authManager: auth,
	}
}

func (g *GetDataBusinessService) GetFullItems(
	ctx context.Context, blendIDs []domain.BlendID) ([]domain.FullItem, error) {
	g.service.SetContext(ctx)

	var fullItems []domain.FullItem

	arrayBlendIds := g.SplitBlendIdsInBlendIdsArray(blendIDs, 20)

	for _, blendIds := range arrayBlendIds {

		arrayStringBulkBlendIds := g.getIdsByChunks(blendIds)

		items, _ := g.service.GetItems(arrayStringBulkBlendIds)
		if items == nil {
			continue
		}

		for _, item := range items {
			category, _ := g.service.GetCategory(item.CategoryID)
			currency, _ := g.service.GetCurrency(item.CurrencyID)
			nickname, _ := g.service.GetUserNickname(item.SellerID)

			fullItems = append(fullItems, domain.FullItem{
				Id:        item.ID,
				Site:      item.Site,
				Price:     item.Price,
				StartTime: item.StarTime,
				Name:      category.ToString(),
				Nickname:  nickname.ToString(),
				Currency:  currency.ToString(),
			})
		}
	}

	firstItem := 0
	for _, ids := range blendIDs {
		if firstItem == 1 {
			break
		}
		items, _ := g.service.GetItems([]string{ids.ToString()})
		if items == nil {
			continue
		}

		category, _ := g.service.GetCategory(items[firstItem].CategoryID)
		currency, _ := g.service.GetCurrency(items[firstItem].CurrencyID)
		nickname, _ := g.service.GetUserNickname(items[firstItem].SellerID)

		fullItems = append(fullItems, domain.FullItem{
			Id:        string(ids.Id),
			Site:      string(ids.Site),
			Price:     items[firstItem].Price,
			StartTime: items[firstItem].StarTime,
			Name:      category.ToString(),
			Nickname:  nickname.ToString(),
			Currency:  currency.ToString(),
		})
		firstItem++
	}

	return fullItems, nil
}

func (g *GetDataBusinessService) getIdsByChunks(blendIds []domain.BlendID) []string {
	var arrayStringBulkBlendIds []string
	for _, blendID := range blendIds {
		arrayStringBulkBlendIds = append(arrayStringBulkBlendIds, blendID.ToString())
	}
	return arrayStringBulkBlendIds
}

func (g *GetDataBusinessService) SplitBlendIdsInBlendIdsArray(blendIDs []domain.BlendID, chunkSize int) [][]domain.BlendID {
	var chunks [][]domain.BlendID
	for {
		if len(blendIDs) == 0 {
			break
		}

		if len(blendIDs) < chunkSize {
			chunkSize = len(blendIDs)
		}

		chunks = append(chunks, blendIDs[0:chunkSize])
		blendIDs = blendIDs[chunkSize:]
	}

	return chunks
}

func (g *GetDataBusinessService) SetAuth(authManager *domain.AuthManager) {
	g.authManager = authManager
	g.service.SetAuth(authManager)

}
