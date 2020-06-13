package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func Init()  {
	fmt.Println("--- Init router")
	Router = mux.NewRouter()
}

func Add(path string, f func(w http.ResponseWriter, r *http.Request), method string)  {
	fmt.Println("--- Router add: " + method + " " + path)
	Router.HandleFunc(path, f).Methods(method)
}