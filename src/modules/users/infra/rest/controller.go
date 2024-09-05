package restuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"talentpitch/src/modules/users/domain"
)

type Controller interface {
	getUsers(ctx *gin.Context)
	createUser(ctx *gin.Context)
	getUserByID(ctx *gin.Context)
	updateUser(ctx *gin.Context)
	deleteUser(ctx *gin.Context)
}

type controller struct {
	useCase domainuser.UseCase
}

func NewController(useCase domainuser.UseCase) Controller {
	return &controller{
		useCase: useCase,
	}
}

func (c *controller) getUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize := 10
	offset := (page - 1) * pageSize

	users, err := c.useCase.GetUsers(pageSize, offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *controller) createUser(ctx *gin.Context) {
	var dataRq domainuser.User

	if err := ctx.ShouldBindJSON(&dataRq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := c.useCase.CreateUser(dataRq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *controller) getUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	user, err := c.useCase.GetUserByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *controller) updateUser(ctx *gin.Context) {
	ID := ctx.Param("id")

	var dataRq domainuser.User
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

func (c *controller) deleteUser(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := c.useCase.DeleteByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
