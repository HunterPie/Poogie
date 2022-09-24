package version

import (
	"strconv"

	"github.com/Haato3o/poogie/core/features/version/models"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type VersionController struct {
	service *VersionService
}

func (c *VersionController) GetLatestVersion(ctx *gin.Context) {
	supporterToken := utils.ExtractSupporterToken(ctx)
	latest, err := c.service.GetLatestFileVersion(ctx, supporterToken)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	http.Ok(ctx, models.LatestVersionResponse{
		LatestVersion: latest,
	})
}

func (c *VersionController) GetLatestBinary(ctx *gin.Context) {
	supporterToken := utils.ExtractSupporterToken(ctx)
	latest, err := c.service.GetLatestFileVersion(ctx, supporterToken)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	latestBinary, err := c.service.GetFileByVersion(ctx, latest, supporterToken)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	ctx.Header("Content-Length", strconv.Itoa(len(latestBinary)))

	http.OkZip(ctx, latestBinary)
}

func (c *VersionController) GetBinaryByVersion(ctx *gin.Context) {
	supporterToken := utils.ExtractSupporterToken(ctx)
	version := ctx.Param("version")

	binary, err := c.service.GetFileByVersion(ctx, version, supporterToken)

	if err != nil {
		http.ElementNotFound(ctx)
		return
	}

	ctx.Header("Content-Length", strconv.Itoa(len(binary)))

	http.OkZip(ctx, binary)
}
