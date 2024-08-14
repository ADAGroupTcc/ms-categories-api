package categories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const CATEGORY_COLLECTION = "categories"

type Repository interface {
}

type categoriesRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return &categoriesRepository{db}
}
