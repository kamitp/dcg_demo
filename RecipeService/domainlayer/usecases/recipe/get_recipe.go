package recipe

import (
	"fmt"
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	"dcg_demo/RecipeService/domainlayer/models"
)

type GetRecipe struct {
	Repo repo.RecipeRepoInterface
}

func (rc GetRecipe) GetRecipe(recipeId string) (models.RecipeModel, error) {

	if len(recipeId) == 0 {
		return models.RecipeModel{}, fmt.Errorf("invalid recipe id")
	}
	return rc.Repo.GetRecipe(recipeId)
}
