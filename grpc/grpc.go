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
	return dbstore.RetrieveAll()
}

func (Server) GetList(_ context.Context, list *pb.List) (*pb.List, error) {
	return dbstore.Retrieve(list)
}

func (Server) AddList(_ context.Context, list *pb.List) (*pb.Lists, error) {
	return dbstore.Save(list)
}
