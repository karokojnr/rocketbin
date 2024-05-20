package grpc

import (
	"context"
	"log"
	"net"

	rkt "github.com/karokojnr/rocketbin-protos/rocket/v1"

	"github.com/karokojnr/rocketbin/internal/rocket"
	"google.golang.org/grpc"
)

// RocketService - defines the methods that the handler will
// use to interact with the service
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rocket rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
	rkt.UnimplementedRocketServiceServer
}

// New - creates a new instance of the gRPC Handler
func New(rocketService RocketService) Handler {
	return Handler{RocketService: rocketService}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("failed to listen on port 50051: ", err)
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to register server: %s\n", err)
		return err
	}
	return nil
}

// GetRocket - retrieves a rocket by id and returns the response
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Println("GetRocket request received")

	rocket, err := h.RocketService.GetRocketByID(ctx, req.Id)
	if err != nil {
		log.Println("error getting rocket by id: ", err)
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC endpoint hit")
	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})
	if err != nil {
		log.Print("failed to insert rocket into database")
		return &rkt.AddRocketResponse{}, err
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

// DeleteRocket - handler for deleting a rocket
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("DeleteRocket request received")
	err := h.RocketService.DeleteRocket(ctx, req.Rocket.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "successfully deleted rocket",
	}, nil
}
