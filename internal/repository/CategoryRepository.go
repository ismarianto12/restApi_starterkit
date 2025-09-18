package repository

import (
	"rianRestApi/internal/models"

	"gorm.io/gorm"
)

type CateGoryRepository struct {
	DB *gorm.DB
}

func NewCateGoryRepository(db *gorm.DB) *CateGoryRepository {
	return &CateGoryRepository{DB: db}
}

func (ct *CateGoryRepository) GetAll() ([]models.CategoryBarang, error) {
	var data []models.CategoryBarang

	if err := ct.DB.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
