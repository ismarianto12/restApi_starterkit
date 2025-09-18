package controller

import (
	"net/http"
	"rianRestApi/config"
	"rianRestApi/internal/repository"
	_ "rianRestApi/internal/repository"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	repo *repository.CateGoryRepository
}

func NewCategoryController() *CategoryController {
	db := config.ConnectDb()
	repo := repository.NewCateGoryRepository(db)
	return &CategoryController{repo: repo}
}

func (ct *CategoryController) All(c *gin.Context) {
	data, err := ct.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
