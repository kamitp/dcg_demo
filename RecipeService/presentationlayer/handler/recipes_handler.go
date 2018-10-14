package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	binders "dcg_demo/RecipeService/databindinglayer"
	repo "dcg_demo/RecipeService/datalayer/repositories"
	recpuc "dcg_demo/RecipeService/domainlayer/usecases/recipe"
	models "dcg_demo/RecipeService/presentationlayer/models"
	utils "dcg_demo/RecipeService/utilities"

	"github.com/gorilla/mux"
)

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	//1. read JWT Token from header
	token, ok := r.Header["Token"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", "Invalid token"))
		return
	}

	//2. Validate it, if not expired then move forward otherwise resopnse an error
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	var model models.CreateRecipeModel
	json.Unmarshal(body, &model)

	err = model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	// 3. Now create instance of create recipe usecase
	uc := recpuc.Recipe{Auth: &repo.TokenAuthRepo{}, Repo: &repo.RecipeRepo{}}
	dmmodel := binders.BindPLCreateRecipeModel_To_DMCreateRecipeModel(&model)
	recipeId, err := uc.CreateRecipe(token[0], &dmmodel)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json_bytes, _ := json.Marshal(map[string]interface{}{"status": "created", "recipe_id": recipeId})
	fmt.Fprint(w, string(json_bytes))
}

func ListRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	uc := recpuc.RecipeList{Repo: &repo.RecipeRepo{}}
	recipeList, err := uc.ListRecipes()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json_bytes, _ := json.Marshal(map[string]interface{}{"status": "error", "message": err.Error()})
		fmt.Fprint(w, string(json_bytes))
		return
	}
	pl := binders.Bind_DMRecipeListModel_To_PLRecipeListModel(&recipeList)
	json_bytes, _ := json.Marshal(map[string]interface{}{"status": "success", "recipes": pl})
	fmt.Fprint(w, string(json_bytes))
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	id := urlParams["id"]
	w.Header().Set("Content-type", "application/json")
	uc := recpuc.GetRecipe{Repo: &repo.RecipeRepo{}}
	recipe, err := uc.GetRecipe(id)
	plrecipe := binders.Bind_DMRecipeModel_To_PLRecipeModel(&recipe)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json_bytes, _ := json.Marshal(map[string]interface{}{"status": "error", "message": err.Error()})
		fmt.Fprint(w, string(json_bytes))
		return
	}
	json_bytes, _ := json.Marshal(map[string]interface{}{"status": "success", "recipe": plrecipe})
	fmt.Fprint(w, string(json_bytes))
}

func Rating(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	ratingModel := models.Rating{}
	err = json.Unmarshal(body, &ratingModel)

	urlParams := mux.Vars(r)
	id := urlParams["id"]

	uc := recpuc.RateRecipe{Repo: &repo.RecipeRepo{}}
	dmRating := binders.Bind_PLRating_To_DMRating(&ratingModel)
	err = uc.Rate(id, dmRating)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, utils.CreateResponse("success", "rating applied"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	//1. read JWT Token from header
	token, ok := r.Header["Token"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", "Invalid token"))
		return
	}

	//2. Validate it, if not expired then move forward otherwise resopnse an error
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	var model models.UpdateRecipeModel

	json.Unmarshal(body, &model)
	err = model.Validate()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	// 3. Now create instance of create recipe usecase
	uc := recpuc.UpdateRecipe{Auth: &repo.TokenAuthRepo{}, Repo: &repo.RecipeRepo{}}
	dmmodel := binders.Bind_PLUpdateRecipeMode_To_DMUpdateRecipeModel(&model)
	err = uc.UpdateRecipe(token[0], mux.Vars(r)["id"], &dmmodel)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	fmt.Fprint(w, utils.CreateResponse("updated", "recipe updated successfully"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Header["Token"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", "Invalid token"))
		return
	}
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	uc := recpuc.DeleteRecipe{Auth: &repo.TokenAuthRepo{}, Repo: &repo.RecipeRepo{}}
	err := uc.DeleteRecipe(token[0], id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
	}
	fmt.Fprint(w, utils.CreateResponse("success", "recipe deleted successfully"))
}

func Search(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	var criteria models.RecipeSearchModel
	json.Unmarshal(body, &criteria)
	uc := recpuc.SearchRecipe{Repo: &repo.RecipeRepo{}}
	recipe, err := uc.SearchRecipe(criteria.DMModel())
	plrecipe := binders.Bind_DMRecipeListModel_To_PLRecipeListModel(&recipe)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	fmt.Fprint(w, utils.CreateResponse("success", plrecipe))
}
