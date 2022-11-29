package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rajatk31/mongo/routes"
)

func main() {
	fmt.Println("Welcome to mongo API")
	r := routes.Routes()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started at 4000")
}
