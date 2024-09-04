package restvideo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	domainvideos "talentpitch/src/modules/videos/domain"
)

type Controller interface {
	getVideo(ctx *gin.Context)
	createVideo(ctx *gin.Context)
	getVideoByID(ctx *gin.Context)
	updateVideo(ctx *gin.Context)
	deleteVideo(ctx *gin.Context)
}

type controller struct {
	useCase domainvideos.UseCase
}

func NewController(useCase domainvideos.UseCase) Controller {
	return &controller{
		useCase: useCase,
	}
}

func (c *controller) getVideo(ctx *gin.Context) {
	Videos, err := c.useCase.GetVideos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, Videos)
}

func (c *controller) createVideo(ctx *gin.Context) {
	var dataRq domainvideos.Videos

	if err := ctx.ShouldBindJSON(&dataRq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := c.useCase.CreateVideo(dataRq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *controller) getVideoByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	Video, err := c.useCase.GetVideoByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, Video)
}

func (c *controller) updateVideo(ctx *gin.Context) {
	ID := ctx.Param("id")

	var dataRq domainvideos.Videos
	if err := ctx.ShouldBindJSON(&dataRq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	dataRq.ID = ID
	err := c.useCase.Update(dataRq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *controller) deleteVideo(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := c.useCase.DeleteByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
