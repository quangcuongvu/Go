package repository

import (
	"go-restful-api/api/models"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	Save(*models.Product) (*models.Product, error)
}

type productsRepositoryImpl struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsRepositoryImpl{
	return &productsRepositoryImpl{db}
}

func (r *productsRepositoryImpl) Save(product *models.Product)(*models.Product,error){
	tx:=r.db.Begin()
	err:=tx.Debug().Model(&models.Product{}).Create(product).Error
	if err!=nil{
		tx.Rollback()
		return nil,err
	}
	return product, tx.Commit().Error
}