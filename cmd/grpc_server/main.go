package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/brianvoe/gofakeit"
	desc "github.com/encountea/pkg/user_api_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserApiV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("ID: %d", req.GetId())

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     "ggfg@gmail.com",
		Role:      desc.Role_admin,
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserApiV1Server(s, &server{})

	log.Printf("Server is listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
