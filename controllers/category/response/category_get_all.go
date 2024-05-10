package response

import "github.com/irvansn/go-find-helpers/entities"

type CategoryGetAll struct {
	Categories []CategoryDetailResponse `json:"categories"`
}

func SliceFromUseCase(categories *[]entities.Category) *CategoryGetAll {
	allCategories := make([]CategoryDetailResponse, len(*categories))
	for i, _category := range *categories {
		allCategories[i] = CategoryDetailResponse{
			ID:   _category.ID,
			Name: _category.Name,
		}
	}

	return &CategoryGetAll{
		Categories: allCategories,
	}
}
