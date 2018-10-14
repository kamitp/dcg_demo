package views

import (
	hnd "dcg_demo/RecipeService/presentationlayer/handler"

	"github.com/gorilla/mux"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

func InitEndpoints(router *mux.Router) {

	router.HandleFunc("/recipes", hnd.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipes", hnd.ListRecipes).Methods("GET")
	router.HandleFunc("/recipes/{id}", hnd.GetRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}/rating", hnd.Rating).Methods("POST")
	router.HandleFunc("/recipes/{id}", hnd.Update).Methods("PUT")
	router.HandleFunc("/recipes/{id}", hnd.Delete).Methods("DELETE")
	router.HandleFunc("/search", hnd.Search).Methods("POST")

}

// type appHandler func(http.ResponseWriter, *http.Request) *appError

// func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
// 		http.Error(w, e.Message, e.Code)
// 	}
// }

// func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
// 	return nil
// }
