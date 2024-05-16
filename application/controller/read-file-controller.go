package controller

import (
	"challenge-scrapy/domain"
	"challenge-scrapy/domain/domain-interfaces"
	domain_services "challenge-scrapy/domain/domain-services"
	"challenge-scrapy/infrastructure/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuildChallengeEntity struct {
	config              domain_interfaces.Config
	businessService     *domain_services.GetDataBusinessService
	challengeRepository domain_interfaces.ChallengeRepository
}

func NewBuildChallengeEntity(c domain_interfaces.Config,
	b *domain_services.GetDataBusinessService,
	r domain_interfaces.ChallengeRepository,
) *BuildChallengeEntity {
	return &BuildChallengeEntity{
		config:              c,
		businessService:     b,
		challengeRepository: r,
	}

}

func (c *BuildChallengeEntity) Read(ctx *gin.Context) {

	readerType := domain.ReaderType(ctx.Query("readerType"))
	if !readerType.IsValid() {
		ctx.JSON(http.StatusBadRequest, "query readerType parameter is required")
		return
	}

	reader, err := services.NewReaderResolver().Select(readerType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "readerType not supported")
		return
	}

	readerOption := domain.NewReaderOption()
	blendIDs, err := reader.Read(readerOption.GetUrlFile(),
		"data",
		readerType.String(),
		readerOption.GetSeparator(),
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fullItems, err := c.businessService.GetFullItems(ctx, blendIDs)
	if err != nil {
		return
	}

	key := ctx.Param("save-key")

	err = c.challengeRepository.SaveFullItem(ctx, key, &fullItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	getFullItem, err := c.challengeRepository.GetFullItem(ctx, key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getFullItem)
}
