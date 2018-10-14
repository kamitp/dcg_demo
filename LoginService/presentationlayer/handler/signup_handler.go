package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	binder "dcg_demo/LoginService/databindinglayer"
	repoProvi "dcg_demo/LoginService/datalayer/repositories"
	"dcg_demo/LoginService/domainlayer/usecases/signup"
	"dcg_demo/LoginService/presentationlayer/models"
	utils "dcg_demo/LoginService/utilities"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	if err != nil {
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	var data models.PLSignupModel
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}

	uc := signup.Signup{Repo: &repoProvi.AuthRepoProvider{}}
	dmdata := binder.BindPLTODMSignup(&data)
	err = uc.Signup(dmdata)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, utils.CreateResponse("error", err.Error()))
		return
	}
	fmt.Fprint(w, utils.CreateResponse("created", "operation success"))
}
