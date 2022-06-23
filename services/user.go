package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Vitoriano/go-gRPC/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultsStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultsStream{
		Status: "Insert Database",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultsStream{
		Status: "User has been save",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultsStream{
		Status: "Complete",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}
