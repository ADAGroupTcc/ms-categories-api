package categories

import "github.com/ADAGroupTcc/ms-categories-api/internal/repositories/categories"

type Service interface {
}

type categoriesService struct {
	categoriesRepository categories.Repository
}

func New(categoriesRepository categories.Repository) Service {
	return &categoriesService{
		categoriesRepository,
	}
}
