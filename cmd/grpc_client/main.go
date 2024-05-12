package main

import (
	"context"
	"log"
	"time"

	desc "github.com/encountea/auth/pkg/user_api_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address  = "localhost:50051"
	randomID = 22
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connetc to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewUserApiV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: randomID})
	if err != nil {
		log.Fatalf("Failed to get by id: %v", err)
	}

	log.Printf(color.RedString("User info:\n"), color.GreenString("%+v", r.GetId()))
}
