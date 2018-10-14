package models

type Rating struct {
	RecipeUser string `bson:"recipe_user"`
	Rating     uint8  `bson:"rating"`
}

type RecipeModel struct {
	Name       string   `bson:"name"`
	UserName   string   `bson:"user_name"`
	PrepTime   uint16   `bson:"prep_time"`
	Difficulty uint8    `bson:"difficulty"`
	IsVeg      bool     `bson:"is_veg"`
	UniqueID   string   `bson:"unique_id"`
	Ratings    []Rating `bson:"ratings"`
}

// type RecipeUpdateModel struct {
// 	Name       string `bson:"name"`
// 	UserName   string `bson:"user_name"`
// 	PrepTime   uint16 `bson:"prep_time"`
// 	Difficulty uint8  `bson:"difficulty"`
// 	IsVeg      bool   `bson:"is_veg"`
// }

type UpdateRecipeModel map[string]interface{}
type RecipeSearchModel map[string]interface{}
