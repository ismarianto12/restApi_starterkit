package repository

import (
	"rianRestApi/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (rp *UserRepository) GetAll() ([]models.UserModel, error) {
	var user []models.UserModel
	if err := rp.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (rp *UserRepository) FindById(id int) (models.UserModel, error) {
	var user models.UserModel
	if err := rp.DB.Find(&user, id).Error; err != nil {
		return models.UserModel{}, err
	}
	return user, nil

}

func (rp *UserRepository) SaveData(data *models.UserModel) (models.UserModel, error) {
	if err := rp.DB.Save(data).Error; err != nil {
		return models.UserModel{}, nil
	}
	return models.UserModel{}, nil

}
