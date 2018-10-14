package views

import (
	hnd "dcg_demo/LoginService/presentationlayer/handler"

	"github.com/gorilla/mux"
)

func InitEndpoints(router *mux.Router) {
	router.HandleFunc("/signup", hnd.Signup).Methods("POST")
	router.HandleFunc("/login", hnd.Login).Methods("GET")
	router.HandleFunc("/logout", hnd.Logout).Methods("GET")
	router.HandleFunc("/validate", hnd.Validate).Methods("POST")
}
