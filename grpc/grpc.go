package grpc

import (
	"context"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type Server struct {
	pb.UnimplementedListerServer
}

func (Server) GetAllLists(context.Context, *pb.Empty) (*pb.Lists, error) {
	return nil, nil
}

func (Server) GetList(_ context.Context, _ *pb.List) (*pb.List, error) {
	return nil, nil
}

func (Server) AddList(_ context.Context, _ *pb.List) (*pb.Lists, error) {
	return nil, nil
}
