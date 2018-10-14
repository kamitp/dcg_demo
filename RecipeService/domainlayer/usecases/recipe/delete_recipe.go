package recipe

import (
	"fmt"
	std "dcg_demo/RecipeService/domainlayer/constants"
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	utils "dcg_demo/RecipeService/utilities"
)

type DeleteRecipe struct {
	Repo repo.RecipeRepoInterface
	Auth repo.TokenAuthInterface
}

func (rc DeleteRecipe) DeleteRecipe(token string, recipeId string) error {
	// 1. Validate token, if success then extract user name
	if len(token) == 0 {
		return fmt.Errorf("invalid token")
	}
	_, err := utils.ValidateJWTToken(token, std.JwtSecrete)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	// 2. Get it varified from auth server
	err = rc.Auth.GetTokenStatusFromServer(token)
	if err != nil {
		return err
	}

	return rc.Repo.DeleteRecipe(recipeId)
}
