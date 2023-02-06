package routes

import (
	"Api-Gateway/pkg/auth/pb"
	"Api-Gateway/pkg/common"
	"context"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	//body := RegisterRequestBody{}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    "a", //body.Email,
		Password: "b", //body.Password,
	})

	if err != nil {
		common.Response(w, err.Error(), http.StatusBadGateway)
		return
	}

	common.Response(w, res.String(), int(res.Status))
}
