package repository

import (
	"rianRestApi/internal/models"

	"gorm.io/gorm"
)

type BarangRepository struct {
	DB *gorm.DB
}

func NewBarangRepository(db *gorm.DB) *BarangRepository {
	return &BarangRepository{DB: db}
}

func (rp *BarangRepository) GetAlldata() ([]models.BarangModel, error) {
	var barang []models.BarangModel
	if err := rp.DB.Find(&barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}

func (rp *BarangRepository) Save(barang *models.BarangModel) (models.BarangModel, error) {
	if err := rp.DB.Save(barang).Error; err != nil {

		return models.BarangModel{}, err
	}
	return *barang, nil
}

func (rp *BarangRepository) Show(id int) (*models.BarangModel, error) {
	var barang models.BarangModel
	if err := rp.DB.First(&barang, id).Error; err != nil {
		return nil, err
	}
	return &barang, nil

}

func (rp *BarangRepository) Filtering(page int, pageSize int, search string) ([]models.BarangModel, int64, error) {
	var barang []models.BarangModel
	var totalrow int64

	query := rp.DB.Model(&models.BarangModel{})
	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("nama LIKE ? OR kode LIKE ?", searchTerm, searchTerm)
	}
	//filtering data
	if err := query.Count(&totalrow).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Find(&barang).Error; err != nil {
		return nil, 0, err
	}
	return barang, totalrow, nil
}
