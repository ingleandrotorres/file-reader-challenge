package main

import (
	"challenge-scrapy/application/controller"
	"challenge-scrapy/domain"
	domain_services "challenge-scrapy/domain/domain-services"
	"challenge-scrapy/infrastructure"
	"challenge-scrapy/infrastructure/authentication"
	"challenge-scrapy/infrastructure/cache"
	"challenge-scrapy/infrastructure/gateways"
	"challenge-scrapy/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"github.com/melisource/fury_go-meli-toolkit-restful/rest"
	"log"
	"time"
)

const ReadFilePath = "/challenge-app/reader/:save-key"
const PingPath = "/ping"
const Port = ":8080"

var authManager *domain.AuthManager

func main() {

	meliClient := getMeliClient()
	cacheClient := getCacheClient()
	infraConfig := infrastructure.NewConfig()

	authGateway := authentication.NewMeliAuthorizationGateway(meliClient)
	authManager, _ = authGateway.RefreshToken()

	gateway := gateways.NewMeliGateway(nil, meliClient, cacheClient, infraConfig, authManager)
	businessService := domain_services.NewGetDataBusinessService(gateway, authManager)

	repository, _ := repositories.NewInMemoryMockKvsClient()

	router := gin.Default()
	router.Group("/challenge-app/reader")

	//todo: authManager refactor to middleware
	//router.Use(authenticationMiddleware(authGateway, authManager))
	router.GET(ReadFilePath, controller.NewBuildChallengeEntity(
		infraConfig,
		businessService,
		repository,
	).Read)
	router.GET(PingPath, controller.NewPingController().Ping)

	err := router.Run(Port)
	if err != nil {
		log.Fatal("error running server", err)
	}
}

func getMeliClient() *rest.RequestBuilder {
	meliToolkitRestClient := &rest.RequestBuilder{
		Timeout:        5 * time.Second,
		ConnectTimeout: 5 * time.Second,
		UserAgent:      "User-Agent",
		EnableCache:    false,
		BaseURL:        "https://api.mercadolibre.com",
	}
	return meliToolkitRestClient
}

func getCacheClient() *cache.InMemoryCache {
	//client, err := redis.NewCache("example")
	//if err != nil {
	//	return nil
	//}

	return cache.NewInMemoryCache()
}

func authenticationMiddleware(authGateway *authentication.MeliAuthorizationGateway,
	manager *domain.AuthManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		//authGateway.SetContext(c)
		manager, _ = authGateway.RefreshToken()
		c.Header("Authorization", "Bearer "+manager.AccessToken)
		c.Next()
	}
}
