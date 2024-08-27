package categories

import (
	"context"
	"errors"

	"github.com/ADAGroupTcc/ms-categories-api/exceptions"
	"github.com/ADAGroupTcc/ms-categories-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CATEGORY_COLLECTION = "categories"

type Repository interface {
	GetCategoryById(ctx context.Context, id primitive.ObjectID) (*domain.Category, error)
	List(ctx context.Context, limit int, offset int) ([]*domain.Category, error)
	ListByCategoryIds(ctx context.Context, categoryIds []primitive.ObjectID, limit int, offset int) ([]*domain.Category, error)
}

type categoriesRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return &categoriesRepository{db}
}

func (r *categoriesRepository) GetCategoryById(ctx context.Context, id primitive.ObjectID) (*domain.Category, error) {
	var category domain.Category = domain.Category{}
	err := category.Read(ctx, r.db, CATEGORY_COLLECTION, bson.M{"_id": id}, &category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, exceptions.New(exceptions.ErrCategoryNotFound, err)
		}
		return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return &category, nil
}

func (r *categoriesRepository) List(ctx context.Context, limit int, offset int) ([]*domain.Category, error) {
	return nil, nil
}

func (r *categoriesRepository) ListByCategoryIds(ctx context.Context, categoryIds []primitive.ObjectID, limit int, offset int) ([]*domain.Category, error) {
	return nil, nil
}
