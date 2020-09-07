package main

import (
	"log"
	"net/http"
	"restapi/framework/http/routers"
)

func main() {
	log.Println("Application is available...")
	log.Fatal(http.ListenAndServe(":8000", routers.GetRouter()))
}
