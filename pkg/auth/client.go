package auth

import (
	"fmt"

	"Api-Gateway/pkg/auth/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

// need config for url
func InitServiceClient() pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	//c.AuthSvcUrl replace hardcoded string
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
