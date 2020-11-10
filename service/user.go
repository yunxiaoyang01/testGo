package service

import (
	"context"
	"testMgo/user"
)

type UserService struct {}

func (s *UserService) GetUser(ctx context.Context,req *user.GetUserRequest)(res *user.GetUserResponse,err error)  {
	res =&user.GetUserResponse{
		Name: "lubin",
		Age: "22",
	}
	return res,nil
}