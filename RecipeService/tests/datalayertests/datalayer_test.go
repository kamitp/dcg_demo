package datalayertests

import (
	repo "dcg_demo/RecipeService/datalayer/repositories"
	"fmt"

	dmmodel "dcg_demo/RecipeService/domainlayer/models"

	"testing"
)

func TestSearchRecipes(t *testing.T) {
	reciRepo := &repo.RecipeRepo{}
	criteria := make(dmmodel.RecipeSearchModel)
	criteria["user_name"] = "robinhood277"

	_, erro := reciRepo.SearchRecipes(criteria)
	if erro != nil {
		fmt.Print(erro.Error())
	}
}
