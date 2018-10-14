package recipe

import (
	"fmt"
	std "dcg_demo/RecipeService/domainlayer/constants"
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	"dcg_demo/RecipeService/domainlayer/models"
	utils "dcg_demo/RecipeService/utilities"
)

type UpdateRecipe struct {
	Repo repo.RecipeRepoInterface
	Auth repo.TokenAuthInterface
}

func (rc UpdateRecipe) UpdateRecipe(token string, recipeID string, dm *models.UpdateRecipeModel) error {

	// 1. Validate token, if success then extract user name
	if len(token) == 0 {
		return fmt.Errorf("Invalid token")
	}
	_, err := utils.ValidateJWTToken(token, std.JwtSecrete)
	if err != nil {
		return fmt.Errorf("Invalid token")
	}
	// 2. Get it varified from auth server
	err = rc.Auth.GetTokenStatusFromServer(token)
	if err != nil {
		return err
	}
	return rc.Repo.Update(recipeID, dm)
}
