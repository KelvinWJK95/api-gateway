package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, code int, str string) {

	//console output
	fmt.Println("Response: ", str)

	//response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]any{"status": code, "error": str})
}
