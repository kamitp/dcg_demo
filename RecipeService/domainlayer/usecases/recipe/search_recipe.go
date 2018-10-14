package recipe

import (
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	"dcg_demo/RecipeService/domainlayer/models"
)

type SearchRecipe struct {
	Repo repo.RecipeRepoInterface
}

func (rc SearchRecipe) SearchRecipe(criteria models.RecipeSearchModel) ([]models.RecipeModel, error) {
	return rc.Repo.SearchRecipes(criteria)
}
