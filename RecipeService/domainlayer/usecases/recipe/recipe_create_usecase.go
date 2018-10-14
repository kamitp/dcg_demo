package recipe

import (
	"fmt"
	std "dcg_demo/RecipeService/domainlayer/constants"
	repo "dcg_demo/RecipeService/domainlayer/interfaces"
	"dcg_demo/RecipeService/domainlayer/models"
	utils "dcg_demo/RecipeService/utilities"
	"reflect"
)

type Recipe struct {
	Repo repo.RecipeRepoInterface
	Auth repo.TokenAuthInterface
}

func (rc Recipe) CreateRecipe(token string, model *models.CreateRecipeModel) (string, error) {

	// 1. Validate token, if success then extract user name
	if len(token) == 0 {
		return "", fmt.Errorf("Invalid token")
	}
	claims, err := utils.ValidateJWTToken(token, std.JwtSecrete)

	if err != nil {
		return "", fmt.Errorf("Invalid token")
	}

	// 2. Get it varified from auth server
	err = rc.Auth.GetTokenStatusFromServer(token)
	if err != nil {
		return "", err
	}
	//3. retriev user name from token
	if reflect.TypeOf(claims["user"]).Kind() == reflect.String {
		model.UserName = reflect.ValueOf(claims["user"]).String()
	}
	return rc.Repo.CreateRecipe(model)
}
