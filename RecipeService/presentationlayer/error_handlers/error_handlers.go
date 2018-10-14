package error_handlers

import (
	"fmt"
	"net/http"
	utils "dcg_demo/RecipeService/utilities"
)

type AppError struct {
	Error   error
	Message string
	Code    int
}

type RecipeErrorHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn RecipeErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Message, err.Code)
	}
}
func HandleError(w http.ResponseWriter, err error) {

	if err != nil {
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
	}
}
