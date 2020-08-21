package controllers

import (
	"fmt"
	"net/http"
)

// Index :
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Index")
}
