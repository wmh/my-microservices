package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wmh/my-microservices/pb/users"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// UserInfoService -
type UserInfoService struct{}

var u = UserInfoService{}

// GetUserInfo -
func (u *UserInfoService) GetUserInfo(ctx context.Context, req *users.UserRequest) (resp *users.UserResponse, err error) {

	name := req.Name
	if name == "wmh" {
		resp = &users.UserResponse{
			Id:    233,
			Name:  name,
			Age:   20,
			Title: []string{"Gopher", "PHPer", "Strongest Avenger"},
		}
	}

	err = nil
	return
}

func main() {
	port := ":2333"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()

	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	users.RegisterUserInfoServiceServer(s, &u)
	s.Serve(l)
}
