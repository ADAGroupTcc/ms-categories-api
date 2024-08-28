package categories

import (
	"context"
	"errors"

	"github.com/ADAGroupTcc/ms-categories-api/exceptions"
	"github.com/ADAGroupTcc/ms-categories-api/internal/domain"
	"github.com/ADAGroupTcc/ms-categories-api/pkg/mongorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CATEGORY_COLLECTION = "categories"

type Repository interface {
	GetCategoryById(ctx context.Context, id primitive.ObjectID) (*domain.Category, error)
	List(ctx context.Context, limit int64, offset int64) ([]*domain.Category, error)
	ListByCategoryIds(ctx context.Context, categoryIds []primitive.ObjectID, limit int64, offset int64) ([]*domain.Category, error)
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

func (r *categoriesRepository) List(ctx context.Context, limit int64, offset int64) ([]*domain.Category, error) {
	var categories []*domain.Category = make([]*domain.Category, 0)
	var filter bson.M
	err := mongorm.List(ctx, r.db, CATEGORY_COLLECTION, filter, &categories, options.Find().SetLimit(limit).SetSkip(offset*limit))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return categories, nil
		}
		return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return categories, nil
}

func (r *categoriesRepository) ListByCategoryIds(ctx context.Context, categoryIds []primitive.ObjectID, limit int64, offset int64) ([]*domain.Category, error) {
	var categories []*domain.Category = make([]*domain.Category, 0)
	filter := bson.M{"_id": bson.M{"$in": categoryIds}}
	err := mongorm.List(ctx, r.db, CATEGORY_COLLECTION, filter, &categories, options.Find().SetLimit(limit).SetSkip(offset*limit))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return categories, nil
		}
		return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return categories, nil
}
