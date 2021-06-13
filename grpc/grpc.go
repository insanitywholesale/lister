package grpc

import (
	"context"
	"gitlab.com/insanitywholesale/lister/models"
	"os"
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
	if os.Getenv("PG_URL") != "" {
		pgURL := os.Getenv("PG_URL")
		if pgURL == "test" {
			db, err := postgres.NewPostgresRepo("postgresql://tester:Apasswd@localhost:5432?sslmode=disable")
			if err != nil {
				log.Fatalf("error %v", err)
			}
		}
	}
	else {
		db, err := postgres.NewPostgresRepo(pgURL)
		if err != nil {
			log.Fatalf("error %v", err)
		}
	}
	dbstore = db
	dbstore, _ = mock.NewMockRepo()
	return
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
