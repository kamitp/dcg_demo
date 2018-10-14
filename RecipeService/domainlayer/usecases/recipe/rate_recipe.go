package recipe

import (
	"fmt"
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	"dcg_demo/RecipeService/domainlayer/models"
)

type RateRecipe struct {
	Repo repo.RecipeRepoInterface
}

func (rc RateRecipe) Rate(recipeId string, rating models.Rating) error {
	if len(recipeId) == 0 {
		return fmt.Errorf("invalid recipe id")
	}

	if rating.Rating < 1 || rating.Rating > 5 {
		return fmt.Errorf("rating shoudl be within 1 to 5")
	}

	return rc.Repo.RateRecipe(recipeId, rating)
}
