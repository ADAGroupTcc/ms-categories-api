package helpers

import (
	"github.com/ADAGroupTcc/ms-categories-api/exceptions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseStringToObjectId(id string) (*primitive.ObjectID, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, exceptions.New(exceptions.ErrInvalidID, err)
	}
	return &parsedId, nil
}
