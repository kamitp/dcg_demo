package handler

import (
	"fmt"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout API called!")
}
