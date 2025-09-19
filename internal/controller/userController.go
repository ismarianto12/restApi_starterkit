package controller

import (
	"net/http"
	"rianRestApi/config"
	"rianRestApi/internal/models"
	"rianRestApi/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repo *repository.UserRepository
}

func NewUserController() *UserController {
	dbcon := config.ConnectDb()
	repoconc := repository.NewUserRepository(dbcon)
	return &UserController{repo: repoconc}
}

func (rp *UserController) GetAll(c *gin.Context) {

	data, err := rp.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data, "messages": "data success fully retrieved",
	})

}

func (rp *UserController) Insert(c *gin.Context) {
	var data models.UserModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": data, "messages": "data success fully retrieved",
		})

	}
	rt, err := rp.repo.SaveData(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": rt, "messages": "data success fully retrieved",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data, "messages": "data success fully retrieved",
	})
}
