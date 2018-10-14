package domaintests

import (
	dm "dcg_demo/RecipeService/domainlayer/interfaces"
	uc "dcg_demo/RecipeService/domainlayer/usecases/recipe"
	"testing"
)

type sampleMode struct {
	dm.RecipeRepoInterface
}

func TestCreateMetho(t *testing.T) {
	testuc := uc.Recipe{}
	//CreateRecipe(token string, model *models.DMRecipeModel)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTUzODI0NzA2MywiaXNzIjoiSGVsbG8gRnJlc2giLCJzdWIiOiJUZXN0IEFzc2lnbm1uZXQiLCJ1c2VyIjoiZ2FuZGhpIn0.WX6lsQJufTnsQzguzCD_msEzbAMbNZX_VDbLpZsZ4Q8"
	testuc.CreateRecipe(token, nil)
}
