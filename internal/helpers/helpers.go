package helpers

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func BindQueryParams(c echo.Context, queryParams *QueryParams) error {
	if err := c.Bind(queryParams); err != nil {
		return err
	}
	queryParams.normalize()
	return nil
}

type QueryParams struct {
	RawCategoryIds string `query:"category_ids"`
	CategoryIDs    []string
	Limit          int `query:"limit"`
	Offset         int `query:"next_page"`
}

func (q *QueryParams) normalize() {
	if q.RawCategoryIds != "" {
		q.CategoryIDs = strings.Split(q.RawCategoryIds, ",")
		q.RawCategoryIds = ""
	}
	if q.Limit < 1 {
		q.Limit = 10
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
}
