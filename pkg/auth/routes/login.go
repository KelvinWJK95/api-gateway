package routes

import (
	"Api-Gateway/pkg/auth/pb"
	"Api-Gateway/pkg/common"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	body := LoginRequestBody{}

	body.Email = r.FormValue("email")
	body.Password = r.FormValue("password")

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	var code int
	var str string

	if err != nil {
		code = http.StatusBadGateway
		str = err.Error()

		common.ErrorResponse(w, code, str)
		return
	}

	code = int(res.Status)
	str = res.String()

	//console output
	fmt.Println("Response: ", str)

	//response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(res)
}
