package main

import (
	"fmt"
	"net/http"

	"github.com/BradleySH/stalker/backend/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := ":8080"
	addr := fmt.Sprintf("Listening on %s\n", port)
	fmt.Println(addr)

	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}
