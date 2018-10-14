package models

type CreateRecipeModel struct {
	Name       string // Recipe name
	UserName   string
	PrepTime   uint16 // In minutes
	Difficulty uint8  // Should be between 1-3
	IsVeg      bool   // either true of false
}

type Rating struct {
	RecipeUser string
	Rating     uint8
}

type RecipeModel struct {
	Name       string
	UserName   string
	PrepTime   uint16
	Difficulty uint8
	IsVeg      bool
	UniqueID   string
	Ratings    []Rating
}

//type UpdateRecipeModel CreateRecipeModel
type UpdateRecipeModel map[string]interface{}
type RecipeSearchModel map[string]interface{}
