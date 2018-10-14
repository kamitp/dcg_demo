package repositories

import (
	"fmt"
	binders "dcg_demo/RecipeService/databindinglayer"
	mongo "dcg_demo/RecipeService/datalayer/datasources/databases/mongodb"
	std "dcg_demo/RecipeService/domainlayer/constants"
	auth "dcg_demo/RecipeService/domainlayer/interfaces"

	dlmodels "dcg_demo/RecipeService/datalayer/models"
	dmmodels "dcg_demo/RecipeService/domainlayer/models"
	"sync"

	"gopkg.in/mgo.v2/bson"
)

type RecipeRepo struct {
	auth.RecipeRepoInterface
	mutex sync.Mutex
	rr    string
}

func (rc RecipeRepo) CreateRecipe(model *dmmodels.CreateRecipeModel) (string, error) {

	dlmodel := binders.Bind_DMCreateRecipeModel_To_DLCreateRecipeModel(model)
	rc.mutex.Lock()
	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)
	rcpid := rc.getNextId()
	if len(rcpid) == 0 {
		return "", fmt.Errorf("recipe cound not be created")
	}
	dlmodel.UniqueID = rcpid
	err := col.Insert(&dlmodel)
	if err != nil {
		return "", err
	}
	rc.mutex.Unlock()
	return rcpid, err
}

func (rc RecipeRepo) ListRecipe() []dmmodels.RecipeModel {
	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)

	var recipeList []dlmodels.RecipeModel
	err := col.Find(bson.M{}).All(&recipeList)
	if err != nil {
		return nil
	}
	dmrecipeList := binders.Bind_DLRecipeListModel_To_DMRecipeListModel(&recipeList) // []models.DMListRecipeModel
	fmt.Print(recipeList)
	return dmrecipeList
}

func (rc RecipeRepo) GetRecipe(recipeID string) (dmmodels.RecipeModel, error) {

	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)

	var recipe dlmodels.RecipeModel
	err := col.Find(bson.M{"unique_id": recipeID}).One(&recipe)
	if err != nil {
		return dmmodels.RecipeModel{}, err
	}
	dmrecipeList := binders.Bind_DLRecipeModel_To_DMRecipeModel(&recipe) // []models.DMListRecipeModel
	return dmrecipeList, nil
}

func (rc RecipeRepo) RateRecipe(recipeID string, rating dmmodels.Rating) error {

	dl := binders.Bind_DMRating_To_DLRating(&rating)
	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)
	err := col.Update(bson.M{"unique_id": recipeID}, bson.M{"$push": bson.M{"ratings": dl}})
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func (rc RecipeRepo) Update(recipeID string, model *dmmodels.UpdateRecipeModel) error {

	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)

	err := col.Update(bson.M{"unique_id": recipeID}, bson.M{"$set": model})
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}

func (rc RecipeRepo) SearchRecipes(criteria dmmodels.RecipeSearchModel) ([]dmmodels.RecipeModel, error) {

	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)
	var recipes []dlmodels.RecipeModel
	err := col.Find(criteria).All(&recipes)

	if err != nil {
		return nil, err
	}
	dm := binders.Bind_DLRecipeListModel_To_DMRecipeListModel(&recipes)
	return dm, nil
}

////////////////////////////////////////////////////////// Private data member/////////////////////////////////////////////////
type recipeIdModel struct {
	RecipeId uint32 `bson:"sequence_value"`
}

func (rc RecipeRepo) DeleteRecipe(recipeID string) error {
	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, std.RecipesCollection)
	return col.Remove(bson.M{"unique_id": recipeID})
}

func (rc RecipeRepo) getNextId() string {
	col := mongo.GetCollectionFromDB(std.RecipeServiceDB, "idgenetor")
	var model recipeIdModel
	err := col.Update(bson.M{"_id": "recipeid"}, bson.M{"$inc": bson.M{"sequence_value": 1}})
	if err != nil {
		fmt.Print(err.Error())
	}

	err = col.Find(bson.M{"_id": "recipeid"}).One(&model)
	if err != nil {
		fmt.Print(err.Error())
		return ""
	}
	return fmt.Sprint(model.RecipeId)
}
