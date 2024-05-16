package gateways

import (
	"challenge-scrapy/domain"
	"challenge-scrapy/infrastructure"
	"challenge-scrapy/infrastructure/cache"
	"challenge-scrapy/infrastructure/responses"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/melisource/fury_go-meli-toolkit-restful/rest"
	"net/http"
	"strconv"
	"strings"
)

const itemsPath = "/items"
const categoryPath = "/categories/%s"
const currenciesPath = "/currencies/%s"
const usersPath = "/users/%s"

type MeliGateway struct {
	ctx         context.Context
	client      *rest.RequestBuilder
	localCache  *cache.InMemoryCache
	conf        *infrastructure.Config
	AuthManager *domain.AuthManager
}

func NewMeliGateway(ctx context.Context,
	c *rest.RequestBuilder,
	localCache *cache.InMemoryCache,
	conf *infrastructure.Config,
	auth *domain.AuthManager) *MeliGateway {
	return &MeliGateway{
		ctx:         ctx,
		client:      c,
		localCache:  localCache,
		conf:        conf,
		AuthManager: auth,
	}
}
func (g *MeliGateway) SetContext(ctx context.Context) {
	g.ctx = ctx
}

/*
GetItems Makes this request example.
curl --location 'https://api.mercadolibre.com/items?ids=MLA750925229%2CMLA594239600&attributes=price%2Cdate_created%2Ccategory_id%2Ccurrency_id%2Cseller_id%2Cid' \
--header 'Authorization: Bearer APP_USR-4935816029516044-050217-4b327bb92a44717f05460ef56c6d8bec-32069679'
*/

func (g *MeliGateway) GetItems(items []string) ([]domain.Item, error) {

	var arrayItem []domain.Item
	arrayItemsResponse := make([]responses.ItemsResponse, 0)
	query := g.getQueryItems(items)
	headerOptionToken := g.getAuthorizationHeader()

	meliResponse := g.client.Get(itemsPath+query, headerOptionToken)
	response := NewResponse(meliResponse.Response, meliResponse.Bytes())

	if response.StatusCodeIs2xx() {
		err := response.Hydrate(&arrayItemsResponse)
		if err != nil {
			return nil, fmt.Errorf("error getting items: %s", response.GetStatus())
		}
		return parseItemsResponseWithCode200ToItem(arrayItemsResponse), nil
	}
	if response.StatusCodeIs4xx() {
		return nil, fmt.Errorf("error getting items: %s", response.GetStatus())
	}
	if response.StatusCodeIs5xx() {
		return nil, fmt.Errorf("error getting items: %s", response.GetStatus())
	}

	return arrayItem, fmt.Errorf("error getting items: %v", meliResponse)
}

/*
GetCategory Makes this request example.
curl --location 'https://api.mercadolibre.com/categories/MLA418285' \
--header 'Authorization: Bearer $ACCESS_TOKEN'
*/

func (g *MeliGateway) GetCategory(categoryID string) (domain.CategoryName, error) {

	var categoryResponse = responses.CategoriesResponse{}

	exist, err := getItemFromCacheIfExist[domain.CategoryName](g.ctx, g.localCache, categoryID)
	if err == nil {
		return exist, nil
	}

	headerOptionToken := g.getAuthorizationHeader()

	meliResponse := g.client.Get(fmt.Sprintf(categoryPath, categoryID), headerOptionToken)
	response := NewResponse(meliResponse.Response, meliResponse.Bytes())

	if response.StatusCodeIs2xx() {
		err := response.Hydrate(&categoryResponse)
		if err != nil {
			return "", fmt.Errorf("error getting category: %s", response.GetStatus())
		}

		category := parseCategoryResponseToCategory(categoryResponse)
		_ = g.localCache.Set(nil, categoryID, category.ToBytes(), 5000)
		return category, nil
	}
	if response.StatusCodeIs4xx() {
		return "", fmt.Errorf("error getting category: %s", response.GetStatus())
	}
	if response.StatusCodeIs5xx() {
		return "", fmt.Errorf("error getting category: %s", response.GetStatus())
	}

	return "", fmt.Errorf("error getting category: %v", meliResponse)
}

/*
	GetCurrency Makes this request example.

curl --location 'https://api.mercadolibre.com/currencies/ARS' \
--header 'Authorization: Bearer $ACCESS_TOKEN'
*/
func (g *MeliGateway) GetCurrency(currencyID string) (domain.CurrencyDescription, error) {

	exist, err := getItemFromCacheIfExist[domain.CurrencyDescription](g.ctx, g.localCache, currencyID)
	if err == nil {
		return exist, nil
	}

	var currencyResponse = responses.CurrenciesResponse{}
	headerOptionToken := g.getAuthorizationHeader()

	meliResponse := g.client.Get(fmt.Sprintf(currenciesPath, currencyID), headerOptionToken)
	response := NewResponse(meliResponse.Response, meliResponse.Bytes())

	if response.StatusCodeIs2xx() {
		err := response.Hydrate(&currencyResponse)
		if err != nil {
			return "", fmt.Errorf("error getting currencies: %s", response.GetStatus())
		}
		currency := parseCurrencyResponseToCurrency(currencyResponse)
		_ = g.localCache.Set(nil, currencyID, currency.ToBytes(), 5000)
		return currency, nil
	}
	if response.StatusCodeIs4xx() {
		return "", fmt.Errorf("error getting currencies: %s", response.GetStatus())
	}
	if response.StatusCodeIs5xx() {
		return "", fmt.Errorf("error getting currencies: %s", response.GetStatus())
	}

	return "", fmt.Errorf("error getting currencies: %v", meliResponse)
}

/*
	GetUserNickname Makes this request example.

curl --location 'https://api.mercadolibre.com/users/59284503' \
--header 'Authorization: Bearer $token'
*/
func (g *MeliGateway) GetUserNickname(sellerID int) (domain.NickName, error) {

	exist, err := getItemFromCacheIfExist[domain.NickName](
		g.ctx, g.localCache, strconv.Itoa(sellerID),
	)
	if err == nil {
		return exist, nil
	}

	var userResponse = responses.UserResponse{}
	headerOptionToken := g.getAuthorizationHeader()

	meliResponse := g.client.Get(
		fmt.Sprintf(usersPath, strconv.Itoa(sellerID)), headerOptionToken,
	)
	response := NewResponse(meliResponse.Response, meliResponse.Bytes())

	if response.StatusCodeIs2xx() {
		err := response.Hydrate(&userResponse)
		if err != nil {
			return "", fmt.Errorf("error getting user: %s", response.GetStatus())
		}
		nickname := parseUserResponseToNickName(userResponse)
		_ = g.localCache.Set(nil, string(sellerID), nickname.ToBytes(), 5000)
		return parseUserResponseToNickName(userResponse), nil
	}
	if response.StatusCodeIs4xx() {
		return "", fmt.Errorf("error getting user: %s", response.GetStatus())
	}
	if response.StatusCodeIs5xx() {
		return "", fmt.Errorf("error getting user: %s", response.GetStatus())
	}

	return "", fmt.Errorf("error getting user: %v", meliResponse)
}

func (g *MeliGateway) getQueryItems(items []string) string {
	idsSeparatedByComma := strings.Join(items, ",")
	query := fmt.Sprintf("?ids=%s&attributes=%s", idsSeparatedByComma, "price,date_created,category_id,currency_id,seller_id,id")
	return query
}
func (g *MeliGateway) getAuthorizationHeader() rest.Option {
	token := g.AuthManager.GetAccessToken()
	//token := "APP_USR-4935816029516044-050319-b6aa45d9e0e1a02ce9c3fc2acfb41bf9-32069679"
	bearerToken := "Bearer " + token
	return rest.Headers(http.Header{"Authorization": []string{bearerToken}})
}
func (g *MeliGateway) SetAuthManager(manager *domain.AuthManager) {
	g.AuthManager = manager

}

func parseItemsResponseWithCode200ToItem(itemsResponses []responses.ItemsResponse) []domain.Item {
	var arrayItem []domain.Item
	for _, itemResponse := range itemsResponses {
		if itemResponse.Code != 200 {
			//TODO: business define do something
			continue
		}
		item := domain.Item{
			ID:         itemResponse.Body.Id,
			Price:      itemResponse.Body.Price,
			Site:       itemResponse.Body.DateCreated,
			CategoryID: itemResponse.Body.CategoryId,
			CurrencyID: itemResponse.Body.CurrencyId,
			SellerID:   itemResponse.Body.SellerId,
			StarTime:   itemResponse.Body.DateCreated,
		}
		arrayItem = append(arrayItem, item)

	}
	return arrayItem
}
func parseCategoryResponseToCategory(categoryResponse responses.CategoriesResponse) domain.CategoryName {
	return domain.CategoryName(categoryResponse.Name)
}
func parseCurrencyResponseToCurrency(categoryResponse responses.CurrenciesResponse) domain.CurrencyDescription {
	return domain.CurrencyDescription(categoryResponse.Description)
}
func parseUserResponseToNickName(userResponse responses.UserResponse) domain.NickName {
	return domain.NickName(userResponse.Nickname)
}

func getItemFromCacheIfExist1[T any](ctx context.Context, cache *cache.InMemoryCache, key string, t T) (T, error) {
	if cache == nil {
		return t, errors.New("cache is nil")
	}

	bytesOfT, err := cache.Get(ctx, key)
	if err != nil {
		return t, err
	}
	if err = json.Unmarshal(bytesOfT, t); err == nil {
		return t, nil
	}
	return t, err
}
func getItemFromCacheIfExist[T any](ctx context.Context, cache *cache.InMemoryCache, key string) (cachedData T, err error) {
	if cache == nil {
		return cachedData, errors.New("cache is nil")
	}

	bytesOfT, err := cache.Get(ctx, key)
	if err != nil {
		return cachedData, err
	}
	if err = json.Unmarshal(bytesOfT, cachedData); err == nil {
		return cachedData, nil
	}
	return cachedData, err
}
func (g *MeliGateway) SetAuth(authManager *domain.AuthManager) {
	g.AuthManager = authManager
}
