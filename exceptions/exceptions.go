package exceptions

import (
	"fmt"
)

const prefix = "categories-api"

var (

	// Errors related to request validation
	ErrInvalidPayload        = fmt.Errorf("%s: invalid payload", prefix)
	ErrInvalidName           = fmt.Errorf("%s: invalid category's name", prefix)
	ErrInvalidDescription    = fmt.Errorf("%s: invalid category's descriptionn", prefix)
	ErrInvalidClassification = fmt.Errorf("%s: invalid category's classification", prefix)
	ErrCategoryAlreadyExists = fmt.Errorf("%s: category already exists", prefix)
	ErrInvalidID             = fmt.Errorf("%s: invalid ID", prefix)

	// Database related errors
	ErrUserNotFound    = fmt.Errorf("%s: user not found", prefix)
	ErrDatabaseFailure = fmt.Errorf("%s: database failure", prefix)
)
