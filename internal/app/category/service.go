package category

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CategoryService interface {
	ServiceGetCategories(data *[]Category) error
	ServiceGetCategory(data *Category, id string) error
	ServiceCreateCategory(data *Category) error
}

type service struct {
	repo Repository
}

func NewCategoryService(repo Repository) CategoryService {
	return &service{repo: repo}
}

func (s *service) ServiceGetCategory(data *Category, id string) error {
	if err := s.repo.RepositoryGetCategory(data, id); err != nil {
		return err
	}

	return nil
}

func (s *service) ServiceGetCategories(data *[]Category) error {
	if err := s.repo.RepositoryGetCategories(data); err != nil {
		return err
	}

	return nil
}

func (s *service) ServiceCreateCategory(data *Category) error {
	if data.Name == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Category name is required")
	}

	if err := s.repo.RepositoryCreatCategory(data); err != nil {
		return err
	}

	return nil
}
