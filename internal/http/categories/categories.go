package categories

import (
	"github.com/ADAGroupTcc/ms-categories-api/internal/services/categories"
)

type Handler interface {
}

type categoriesHandler struct {
	categoriesService categories.Service
}

func New(categoriesService categories.Service) Handler {
	return &categoriesHandler{
		categoriesService,
	}
}
