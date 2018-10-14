package interfaces

import (
	models "dcg_demo/RecipeService/domainlayer/models"
)

type RecipeRepoInterface interface {
	CreateRecipe(model *models.CreateRecipeModel) (string, error)
	ListRecipe() []models.RecipeModel
	GetRecipe(recipeID string) (models.RecipeModel, error)
	RateRecipe(recipeId string, rating models.Rating) error
	Update(recipeID string, model *models.UpdateRecipeModel) error
	DeleteRecipe(recipeId string) error
	SearchRecipes(criteria models.RecipeSearchModel) ([]models.RecipeModel, error)
}
