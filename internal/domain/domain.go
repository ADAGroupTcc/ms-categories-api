package domain

import (
	"github.com/ADAGroupTcc/ms-categories-api/exceptions"
	"github.com/ADAGroupTcc/ms-categories-api/pkg/mongorm"
	"go.mongodb.org/mongo-driver/bson"
)

type Category struct {
	mongorm.Model  `json:",inline" bson:",inline"`
	Name           string `json:"name" bson:"name"`
	Description    string `json:"description" bson:"description"`
	Classification int    `json:"classification" bson:"classification"`
}

func (c *Category) ToBsonM() bson.M {
	return bson.M{"$set": bson.M{
		"name":           c.Name,
		"description":    c.Description,
		"classification": c.Classification,
	},
	}
}

type CategoriesResponse struct {
	Categories []*Category `json:"categories"`
	NextPage   int         `json:"next_page,omitempty"`
}

type CategoryRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Classification int    `json:"classification"`
}

func (c *CategoryRequest) ToCategory() *Category {
	return &Category{
		Name:           c.Name,
		Description:    c.Description,
		Classification: c.Classification,
	}
}

func (c *CategoryRequest) Validate() error {
	if c.Name == "" || len(c.Name) < 3 {
		return exceptions.New(exceptions.ErrInvalidName, nil)
	}
	if c.Description == "" || len(c.Description) < 3 {
		return exceptions.New(exceptions.ErrInvalidDescription, nil)
	}

	if c.Classification <= 0 || c.Classification > 5 {
		return exceptions.New(exceptions.ErrInvalidClassification, nil)
	}

	return nil
}
