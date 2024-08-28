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
	listResponse, err := s.categoriesRepository.List(ctx, int64(limit), int64(offset))
	if err != nil {
		return nil, err
	}
	return buildCategoriesResponse(listResponse, limit, offset), nil
}

func (s *categoriesService) ListByCategoryIds(ctx context.Context, categoryIds []string, limit int, offset int) (*domain.CategoriesResponse, error) {
	var parsedIds []primitive.ObjectID
	for _, id := range categoryIds {
		parsedId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, exceptions.New(exceptions.ErrInvalidID, err)
		}
		parsedIds = append(parsedIds, parsedId)
	}
	listResponse, err := s.categoriesRepository.ListByCategoryIds(ctx, parsedIds, int64(limit), int64(offset))
	if err != nil {
		return nil, err
	}
	return buildCategoriesResponse(listResponse, limit, offset), nil
}

func buildCategoriesResponse(categories []*domain.Category, limit int, offset int) *domain.CategoriesResponse {
	res := &domain.CategoriesResponse{
		Categories: categories,
	}
	if len(categories) == limit {
		res.NextPage = offset + 1
	}
	return res
}
