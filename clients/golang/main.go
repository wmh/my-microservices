package main

import (
	"fmt"
	"log"

	"github.com/wmh/my-microservices/pb/users"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":2333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	// 实例化 UserInfoService 微服务的客户端
	client := users.NewUserInfoServiceClient(conn)

	// 调用服务
	req := new(users.UserRequest)
	req.Name = "wmh"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}
	fmt.Printf("Recevied: %v\n", resp)
}
