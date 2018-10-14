package handler

import (
	"fmt"
	"net/http"

	repoProvi "dcg_demo/LoginService/datalayer/repositories"
	loginuc "dcg_demo/LoginService/domainlayer/usecases/login"
	utils "dcg_demo/LoginService/utilities"
)

func Login(w http.ResponseWriter, r *http.Request) {

	authKey, ok := r.Header["Authorization"]
	w.Header().Set("Content-type", "application/json")
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", "authorization key not found"))
		return
	}

	ps, err := utils.DecodeBasicAuthorizationKey(authKey[0])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", "user name or password did not match"))
		return
	}

	uc := loginuc.Login{Repo: &repoProvi.AuthRepoProvider{}}
	jwtToken, err := uc.Login(ps[0], ps[1])

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, utils.CreateResponse("error", "user name or password did not match"))
		return
	}
	fmt.Fprint(w, utils.CreateResponse("success", jwtToken))
}
