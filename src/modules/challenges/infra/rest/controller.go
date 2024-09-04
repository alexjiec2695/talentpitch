package restchallenges

import (
	"github.com/gin-gonic/gin"
	"net/http"
	domainchallenge "talentpitch/src/modules/challenges/domain"
)

type Controller interface {
	getChallenges(ctx *gin.Context)
	createChallenges(ctx *gin.Context)
	getChallengesByID(ctx *gin.Context)
	updateChallenges(ctx *gin.Context)
	deleteChallenges(ctx *gin.Context)
}

type controller struct {
	useCase domainchallenge.UseCase
}

func NewController(useCase domainchallenge.UseCase) Controller {
	return &controller{
		useCase: useCase,
	}
}

func (c *controller) getChallenges(ctx *gin.Context) {
	Challenges, err := c.useCase.GetChallenges()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, Challenges)
}

func (c *controller) createChallenges(ctx *gin.Context) {
	var dataRq domainchallenge.Challenges

	if err := ctx.ShouldBindJSON(&dataRq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := c.useCase.CreateChallenges(dataRq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *controller) getChallengesByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	Challenge, err := c.useCase.GetChallengesByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, Challenge)
}

func (c *controller) updateChallenges(ctx *gin.Context) {
	ID := ctx.Param("id")

	var dataRq domainchallenge.Challenges
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

func (c *controller) deleteChallenges(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := c.useCase.DeleteByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
