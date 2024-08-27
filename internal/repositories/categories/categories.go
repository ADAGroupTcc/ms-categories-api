package categories

import (
	"context"

	"github.com/ADAGroupTcc/ms-categories-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CATEGORY_COLLECTION = "categories"

type Repository interface {
	GetCategoryById(ctx context.Context, id string) (*domain.Category, error)
	List(ctx context.Context, limit int, offset int) ([]*domain.Category, error)
	ListByCategoryIds(ctx context.Context, categoryIds []primitive.ObjectID, limit int, offset int) ([]*domain.Category, error)
}

type categoriesRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return &categoriesRepository{db}
}

func (r *categoriesRepository) GetCategoryById(ctx context.Context, id string) (*domain.Category, error)
func (r *categoriesRepository) List(ctx context.Context, limit int, offset int) ([]*domain.Category, error)
func (r *categoriesRepository) ListByCategoryIds(ctx context.Context, categoryIds []primitive.ObjectID, limit int, offset int) ([]*domain.Category, error)
