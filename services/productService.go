package services

import (
	"github.com/gophers-latam/small-template/internal"
	"github.com/gophers-latam/small-template/models"
	"github.com/gophers-latam/small-template/storage"
)

type ProductService struct {
	repo   storage.Repository
	logger internal.Logger
}

func NewProductService(repo storage.Repository, logger internal.Logger) *ProductService {
	return &ProductService{repo, logger}
}

func (service *ProductService) Get(id int) *models.Product {
	return service.repo.Get(id)
}

func (service *ProductService) Save(product *models.Product) {
	service.repo.Save(product)
	service.logger.Info("product saved")
}

func (service *ProductService) Delete(id int) {
	service.repo.Delete(id)
	service.logger.Info("product deleted")
}
