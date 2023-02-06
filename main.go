package main

import (
	"net/http"

	"Api-Gateway/pkg/auth"
)

func main() {
	mux := http.NewServeMux()

	auth.RegisterRoutes(mux)

	http.ListenAndServe(":8081", mux)
}
