// rian reposotory

package repository

import (
	"rianRestApi/internal/models"

	"gorm.io/gorm"
)

type PurchasingRepository struct {
	DB *gorm.DB
}

func NewPurchasingRepository(db *gorm.DB) *PurchasingRepository {
	return &PurchasingRepository{DB: db}
}

func (rp *PurchasingRepository) IndexAll() ([]models.PurcashingModel, error) {
	var data []models.PurcashingModel

	if err := rp.DB.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (rp *PurchasingRepository) Insert(data models.PurcashingModel) (models.PurcashingModel, error) {

	if err := rp.DB.Save(&data).Error; err != nil {
		return models.PurcashingModel{}, err
	}
	return data, nil

}

func (repo *PurchasingRepository) UpdateData(id int) (models.PurcashingModel, error) {
	var modelupdate models.PurcashingModel
	err := repo.DB.First(&modelupdate, id)
	if err != nil {
		return modelupdate, nil
	}

	if err := repo.DB.Save(&modelupdate).Error; err != nil {
		return modelupdate, err
	}
	return modelupdate, nil
}

func (rep *PurchasingRepository) ShowByID(id int) (models.PurcashingModel, error) {
	var modelnya models.PurcashingModel

	if err := rep.DB.First(&modelnya, id).Error; err != nil {
		return modelnya, nil
	}
	return modelnya, nil
}

func (rep *PurchasingRepository) DeleteById(id int) error {
	if err := rep.DB.Delete(id).Error; err != nil {
		return nil
	}
	return nil
}
