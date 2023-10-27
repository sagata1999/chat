package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/sagata1999/chat/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

// Create
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create request with: %v", req.GetUsernames())

	// while there's nowhere to save -> just yield back received info
	return &desc.CreateResponse{
		Id: int64(gofakeit.Number(0, 100)),
	}, nil
}

// Delete
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	// while there's nothing to delete -> just yield back received info
	log.Printf("Delete request with: id=%d", req.GetId())

	return &emptypb.Empty{}, nil
}

// SendMessage
func (s *server) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	// while there's nothing to delete -> just yield back received info
	log.Printf("Delete request with: from=%s text=%s timestamp=%v", req.GetFrom(), req.GetText(), req.GetTimestamp())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
