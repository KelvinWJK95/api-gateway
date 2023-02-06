package auth

import (
	"net/http"

	"Api-Gateway/pkg/auth/routes"
)

func RegisterRoutes(mux *http.ServeMux) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	mux.HandleFunc("/register", svc.Register)

	// mux.HandleFunc("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(w http.ResponseWriter, r *http.Request) {
	routes.Register(w, r, svc.Client)
}

// func (svc *ServiceClient) Login(w http.ResponseWriter, r *http.Request) {
// 	routes.Login(r.Body, svc.Client)
// }
