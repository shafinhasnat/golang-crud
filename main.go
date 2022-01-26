package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shafinhasnat/kloudlab/domains/auth"
)

func main() {
	r := mux.NewRouter()
	auth.RegisterRoutes(r)
	fmt.Println("LISTENING ON PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
