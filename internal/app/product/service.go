package product

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type ProductService interface {
	ServiceGetProducts(data *[]Product) error
	ServiceGetProduct(data *Product, id string) error
	ServiceCreateProduct(data *Product) error
}

type service struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &service{repo: repo}
}

func (s *service) ServiceGetProducts(data *[]Product) error {
	if err := s.repo.RepositoryGetProducts(data); err != nil {
		return err
	}

	return nil
}

func (s *service) ServiceGetProduct(data *Product, id string) error {
	if id == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Product id is required")
	}

	if err := s.repo.RepositoryGetProduct(data, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ServiceCreateProduct(data *Product) error {
	if data.Name == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Category name is required")
	}

	if data.Price == 0 || data.Price < 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Price is required")
	}

	if data.CategoryName == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Category name is required")
	}

	if err := s.repo.RepositoryCreatProduct(data); err != nil {
		return err
	}

	return nil
}
