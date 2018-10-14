package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	repoProvi "dcg_demo/LoginService/datalayer/repositories"
	"dcg_demo/LoginService/domainlayer/usecases/validate"
	utils "dcg_demo/LoginService/utilities"
)

type ValidationModel struct {
	Token string `json:"jwt_token"`
}

func Validate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	if err != nil {
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	var data ValidationModel
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	uc := validate.Validate{Repo: &repoProvi.AuthRepoProvider{}}
	err = uc.Validate(data.Token)
	if err != nil {
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	fmt.Fprint(w, utils.CreateResponse("success", "valid"))
}
