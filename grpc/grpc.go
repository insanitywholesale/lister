package grpc

import (
	"context"
	"gitlab.com/insanitywholesale/lister/models"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"gitlab.com/insanitywholesale/lister/repo/mock"
)

type Server struct {
	pb.UnimplementedListerServer
}

var dbstore models.ListsRepo

func init() {
	dbstore, _ = mock.NewMockRepo()
}

func (Server) GetAllLists(context.Context, *pb.Empty) (*pb.Lists, error) {
	lists, err := dbstore.RetrieveAll()
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (Server) GetList(_ context.Context, list *pb.List) (*pb.List, error) {
	list, err := dbstore.Retrieve(list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (Server) AddList(_ context.Context, _ *pb.List) (*pb.Lists, error) {
	return &pb.Lists{}, nil
}
