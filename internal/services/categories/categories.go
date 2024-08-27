package categories

import (
	"context"

	"github.com/ADAGroupTcc/ms-categories-api/exceptions"
	"github.com/ADAGroupTcc/ms-categories-api/internal/domain"
	"github.com/ADAGroupTcc/ms-categories-api/internal/repositories/categories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetCategoryById(ctx context.Context, id string) (*domain.Category, error)
	List(ctx context.Context, limit int, offset int) (*domain.CategoriesResponse, error)
	ListByCategoryIds(ctx context.Context, categoryIds []string, limit int, offset int) (*domain.CategoriesResponse, error)
}

type categoriesService struct {
	categoriesRepository categories.Repository
}

func New(categoriesRepository categories.Repository) Service {
	return &categoriesService{
		categoriesRepository,
	}
}

func (s *categoriesService) GetCategoryById(ctx context.Context, id string) (*domain.Category, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, exceptions.New(exceptions.ErrInvalidID, err)
	}
	return s.categoriesRepository.GetCategoryById(ctx, parsedId)
}

func (s *categoriesService) List(ctx context.Context, limit int, offset int) (*domain.CategoriesResponse, error) {
	return nil, nil
}

func (s *categoriesService) ListByCategoryIds(ctx context.Context, categoryIds []string, limit int, offset int) (*domain.CategoriesResponse, error) {
	return nil, nil
}
