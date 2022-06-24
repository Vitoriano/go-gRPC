package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Vitoriano/go-gRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server? %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receve the msg: %v", err)
		}

		fmt.Println("Status:", stream.Status, "-", stream.GetUser())
	}

}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "Vitoriano",
			Email: "vitor@com.br",
		},
		&pb.User{
			Id:    "w2",
			Name:  "Vitoriano2",
			Email: "vitor@com.br",
		},
		&pb.User{
			Id:    "w3",
			Name:  "Vitoriano3",
			Email: "vitor@com.br",
		},
		&pb.User{
			Id:    "w4",
			Name:  "Vitoriano4",
			Email: "vitor@com.br",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receving response %v", err)
	}

	fmt.Println(res)

}
