package databindinglayer

import (
	dlmodel "dcg_demo/RecipeService/datalayer/models"
	dmmodel "dcg_demo/RecipeService/domainlayer/models"
	plmodel "dcg_demo/RecipeService/presentationlayer/models"
)

/////// DM -> DL /////////////////////////
/////// DM -> PL /////////////////////////
func Bind_DMRecipeModel_To_PLRecipeModel(dm *dmmodel.RecipeModel) plmodel.RecipeModel {
	pl := plmodel.RecipeModel{
		Difficulty: dm.Difficulty,
		IsVeg:      dm.IsVeg,
		Name:       dm.Name,
		PrepTime:   dm.PrepTime,
		UniqueID:   dm.UniqueID,
		UserName:   dm.UserName}
	pl.Ratings = make([]plmodel.Rating, len(dm.Ratings))
	copyRatings2(&pl.Ratings, &dm.Ratings)
	return pl
}

func copyRatings2(dest *[]plmodel.Rating, src *[]dmmodel.Rating) {
	if dest != nil && src != nil && len(*dest) != 0 && len(*src) != 0 {
		for i := 0; i < len(*dest); i++ {
			(*dest)[i].Rating = (*src)[i].Rating
			(*dest)[i].RecipeUser = (*src)[i].RecipeUser
		}
	}
}

func Bind_PLRating_To_DMRating(src *plmodel.Rating) dmmodel.Rating {
	return dmmodel.Rating{Rating: src.Rating, RecipeUser: src.RecipeUser}
}

func Bind_DMRating_To_DLRating(src *dmmodel.Rating) dlmodel.Rating {
	return dlmodel.Rating{Rating: src.Rating, RecipeUser: src.RecipeUser}
}

/////// PL -> DM /////////////////////////

func Bind_PLUpdateRecipeMode_To_DMUpdateRecipeModel(pl *plmodel.UpdateRecipeModel) dmmodel.UpdateRecipeModel {

	dm := make(dmmodel.UpdateRecipeModel)
	for k, v := range *pl {
		dm[k] = v
	}
	return dm
}

func BindPLCreateRecipeModel_To_DMCreateRecipeModel(plmodel *plmodel.CreateRecipeModel) dmmodel.CreateRecipeModel {
	return dmmodel.CreateRecipeModel{Difficulty: plmodel.Difficulty,
		IsVeg:    plmodel.IsVeg,
		Name:     plmodel.Name,
		PrepTime: plmodel.PrepTime}
}

func Bind_DMCreateRecipeModel_To_DLCreateRecipeModel(dm *dmmodel.CreateRecipeModel) dlmodel.RecipeModel {
	return dlmodel.RecipeModel{
		Name:       dm.Name,
		UserName:   dm.UserName,
		PrepTime:   dm.PrepTime,
		Difficulty: dm.Difficulty,
		IsVeg:      dm.IsVeg}
}

// DLmodel to DM model
func Bind_DLRecipeListModel_To_DMRecipeListModel(dl *[]dlmodel.RecipeModel) []dmmodel.RecipeModel {

	if dl != nil && len(*dl) != 0 {
		dm := make([]dmmodel.RecipeModel, len(*dl))
		for i := 0; i < len(*dl); i++ {
			dm[i].UserName = (*dl)[i].UserName
			dm[i].Difficulty = (*dl)[i].Difficulty
			dm[i].IsVeg = (*dl)[i].IsVeg
			dm[i].Name = (*dl)[i].Name
			dm[i].PrepTime = (*dl)[i].PrepTime
			dm[i].Ratings = make([]dmmodel.Rating, len((*dl)[i].Ratings))
			copyRatings(&dm[i].Ratings, &(*dl)[i].Ratings)
			dm[i].UniqueID = (*dl)[i].UniqueID
		}
		return dm
	}
	return nil
}

func Bind_DMRecipeListModel_To_PLRecipeListModel(dm *[]dmmodel.RecipeModel) []plmodel.RecipeModel {

	if dm != nil && len(*dm) != 0 {
		pl := make([]plmodel.RecipeModel, len(*dm))
		for i := 0; i < len(*dm); i++ {
			pl[i].UserName = (*dm)[i].UserName
			pl[i].Difficulty = (*dm)[i].Difficulty
			pl[i].IsVeg = (*dm)[i].IsVeg
			pl[i].Name = (*dm)[i].Name
			pl[i].PrepTime = (*dm)[i].PrepTime
			pl[i].Ratings = make([]plmodel.Rating, len((*dm)[i].Ratings))
			for j := 0; j < len((*dm)[i].Ratings); j++ {
				pl[i].Ratings[j].Rating = (*dm)[i].Ratings[j].Rating
				pl[i].Ratings[j].RecipeUser = (*dm)[i].Ratings[j].RecipeUser
			}
			pl[i].UniqueID = (*dm)[i].UniqueID
		}
		return pl
	}
	return nil
}

func Bind_DLRecipeModel_To_DMRecipeModel(dl *dlmodel.RecipeModel) dmmodel.RecipeModel {

	dm := dmmodel.RecipeModel{
		Difficulty: dl.Difficulty,
		IsVeg:      dl.IsVeg,
		Name:       dl.Name,
		PrepTime:   dl.PrepTime,
		UniqueID:   dl.UniqueID,
		UserName:   dl.UserName}
	dm.Ratings = make([]dmmodel.Rating, len(dl.Ratings))
	copyRatings(&dm.Ratings, &dl.Ratings)

	return dm
}

// DMmodel to DLmodel
func copyRatings(dest *[]dmmodel.Rating, src *[]dlmodel.Rating) {
	if dest != nil && src != nil && len(*dest) != 0 && len(*src) != 0 {
		for i := 0; i < len(*dest); i++ {
			(*dest)[i].Rating = (*src)[i].Rating
			(*dest)[i].RecipeUser = (*src)[i].RecipeUser
		}
	}
}
