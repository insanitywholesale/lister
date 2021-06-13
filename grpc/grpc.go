package grpc

import (
	"context"
	"gitlab.com/insanitywholesale/lister/models"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"gitlab.com/insanitywholesale/lister/repo/mock"
	"gitlab.com/insanitywholesale/lister/repo/postgres"
	"log"
)

type Server struct {
	pb.UnimplementedListerServer
}

var dbstore models.ListsRepo

func init() {
	db, err := postgres.NewPostgresRepo("postgresql://tester:Apasswd@localhost:5432?sslmode=disable")
	if err != nil {
		log.Fatalf("error %v", err)
	}
	dbstore, _ = mock.NewMockRepo()
	dbstore = db
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
