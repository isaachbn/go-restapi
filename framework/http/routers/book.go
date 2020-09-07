package routers

import (
	"github.com/gorilla/mux"
	"restapi/framework/http/hadlers"
)

var router *mux.Router

func init()  {
	//Init Router
	router = mux.NewRouter()
}

//Get all routers by book endpoint
func GetRouter() *mux.Router {
	//Init Router
	router := mux.NewRouter()

	//Router Handlers and Endpoints
	router.HandleFunc("/api/books", hadlers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", hadlers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", hadlers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", hadlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", hadlers.DeleteBook).Methods("DELETE")

	return router
}
