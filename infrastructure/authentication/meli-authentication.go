package authentication

import (
	"challenge-scrapy/domain"
	"challenge-scrapy/infrastructure/gateways"
	"fmt"
	"github.com/melisource/fury_go-meli-toolkit-restful/rest"
	"net/http"
	"net/url"
	"strings"
)

const refreshTokenPath = "/oauth/token"

type MeliAuthorizationGateway struct {
	client *rest.RequestBuilder
}

func NewMeliAuthorizationGateway(c *rest.RequestBuilder) *MeliAuthorizationGateway {
	return &MeliAuthorizationGateway{
		client: c,
	}
}

/*
GetItems Makes this request example.
curl --location 'https://api.mercadolibre.com/oauth/token' \
--header 'accept: application/json' \
--header 'content-type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=refresh_token' \
--data-urlencode 'client_id=4935816029516044' \
--data-urlencode 'client_secret=mEHurdrHfh3BAqPb36DwcCKDDm0IJvGd' \
--data-urlencode 'refresh_token=TG-6630502c8ea21c00010d19fa-32069679'
*/

func (g MeliAuthorizationGateway) RefreshToken() (*domain.AuthManager, error) {

	var refreshToken domain.AuthManager
	var refreshTokenResponse RefreshTokenResponse

	headers := http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/x-www-form-urlencoded"},
	}

	meliResponse := g.client.Post(refreshTokenPath, urlEncodedAuthData(), rest.Headers(headers))
	response := gateways.NewResponse(meliResponse.Response, meliResponse.Bytes())

	if response.StatusCodeIs2xx() {
		err := response.Hydrate(&refreshTokenResponse)
		if err != nil {
			return nil, fmt.Errorf("error getting auth: %s", response.GetStatus())
		}
		refreshToken = domain.AuthManager{
			AccessToken:  refreshTokenResponse.AccessToken,
			TokenType:    refreshTokenResponse.TokenType,
			ExpiresIn:    refreshTokenResponse.ExpiresIn,
			Scope:        refreshTokenResponse.Scope,
			UserId:       refreshTokenResponse.UserId,
			RefreshToken: refreshTokenResponse.RefreshToken,
		}
		return &refreshToken, nil
	}
	if response.StatusCodeIs4xx() {
		return nil, fmt.Errorf("error getting auth: %s", response.GetStatus())
	}
	if response.StatusCodeIs5xx() {
		return nil, fmt.Errorf("error getting auth: %s", response.GetStatus())
	}

	return nil, fmt.Errorf("error getting auth: %v", meliResponse)
}

func urlEncodedAuthData2() *strings.Reader {
	//todo: agregar config
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	//todo pasar a secret
	data.Set("client_id", "4935816029516044")
	//todo pasar a secret
	data.Set("client_secret", "mEHurdrHfh3BAqPb36DwcCKDDm0IJvGd")
	//todo pasar a secret
	data.Set("refresh_token", "TG-6630502c8ea21c00010d19fa-32069679")

	return strings.NewReader(data.Encode())

}
func urlEncodedAuthData() RefreshTokenRequest {

	return RefreshTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     "4935816029516044",
		ClientSecret: "mEHurdrHfh3BAqPb36DwcCKDDm0IJvGd",
		RefreshToken: "TG-6630502c8ea21c00010d19fa-32069679",
	}

}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	UserId       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}
