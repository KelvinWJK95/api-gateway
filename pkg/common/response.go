package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, str string, code int) {

	//console output
	fmt.Println("Response: ", str)

	//response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if code < 300 {
		json.NewEncoder(w).Encode(map[string]any{"data": str, "status": code})
	} else {
		json.NewEncoder(w).Encode(map[string]any{"message": str, "status": code, "error": http.StatusText(code)})
	}

}
