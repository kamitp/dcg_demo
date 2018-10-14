package models

import (
	"fmt"
	dmmodel "dcg_demo/RecipeService/domainlayer/models"
	"reflect"
)

type CreateRecipeModel struct {
	Name       string `json:"name"`
	PrepTime   uint16 `json:"prep_time"`
	Difficulty uint8  `json:"difficulty"`
	IsVeg      bool   `json:"is_veg"`
}

func (md CreateRecipeModel) Validate() error {

	if len(md.Name) == 0 {
		return fmt.Errorf("recipe name is empty")
	} else if md.PrepTime < 1 {
		return fmt.Errorf("wrong preparation time")
	} else if md.Difficulty < 1 || md.Difficulty > 3 {
		return fmt.Errorf("difficulty is not within range")
	}

	return nil
}

func (md CreateRecipeModel) DomainCreateRecipeModel() dmmodel.CreateRecipeModel {
	return dmmodel.CreateRecipeModel{Difficulty: md.Difficulty,
		IsVeg:    md.IsVeg,
		Name:     md.Name,
		PrepTime: md.PrepTime}
}

// type PLRecipeListModel_ struct {
// 	RecipeList []PLListRecipeModel
// }

type Rating struct {
	RecipeUser string `json:"recipe_user"`
	Rating     uint8  `json:"rating"`
}

type RecipeModel struct {
	Name       string   `json:"name"`
	UserName   string   `json:"user_name"`
	PrepTime   uint16   `json:"prep_time"`
	Difficulty uint8    `json:"difficulty"`
	IsVeg      bool     `json:"is_veg"`
	UniqueID   string   `json:"unique_id"`
	Ratings    []Rating `json:"ratings"`
}

type UpdateRecipeModel map[string]interface{}

func (md UpdateRecipeModel) Validate() error {
	return validateRecipeModel(md)
}

func validateRecipeModel(recipeMode map[string]interface{}) error {
	for k, v := range recipeMode {
		if k == "name" {
			if reflect.ValueOf(v).Kind() == reflect.String {
				if len(reflect.ValueOf(v).String()) == 0 {
					return fmt.Errorf("name is empty")
				}
			}
		} else if k == "prep_time" {
			if reflect.ValueOf(v).Kind() == reflect.Float64 {
				pt := reflect.ValueOf(v).Float()
				if pt < 1 {
					return fmt.Errorf("invalid preparation time")
				}
			}
		} else if k == "difficulty" {
			if reflect.ValueOf(v).Kind() == reflect.Float64 {
				diff := reflect.ValueOf(v).Float()
				if diff < 1 || diff > 3 {
					return fmt.Errorf("invalid difficulty")
				}
			}
		} else {
			return fmt.Errorf(k + " can't be updated")
		}
	}
	return nil
}

type RecipeSearchModel map[string]interface{}

func (rc RecipeSearchModel) DMModel() dmmodel.RecipeSearchModel {
	dm := make(dmmodel.RecipeSearchModel)
	for k, v := range rc {
		dm[k] = v
	}
	return dm
}
