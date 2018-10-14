package recipe

import (
	"fmt"
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	"dcg_demo/RecipeService/domainlayer/models"
)

type RecipeList struct {
	Repo repo.RecipeRepoInterface
}

func (rc RecipeList) ListRecipes() ([]models.RecipeModel, error) {
	list := rc.Repo.ListRecipe()
	if list != nil && len(list) != 0 {
		return list, nil
	}

	return nil, fmt.Errorf("recipe not found")
}
