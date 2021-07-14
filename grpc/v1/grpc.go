package grpc

import (
	"context"
	"gitlab.com/insanitywholesale/lister/models"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"gitlab.com/insanitywholesale/lister/repo/cassandra"
	"gitlab.com/insanitywholesale/lister/repo/mock"
	"gitlab.com/insanitywholesale/lister/repo/postgres"
	"log"
	"os"
	"strings"
)

type Server struct {
	pb.UnimplementedListerServer
}

var dbstore models.ListsRepo

func init() {
	cassURL := os.Getenv("CASS_URL")
	if cassURL != "" {
		if cassURL == "test" {
			db, err := cassandra.NewCassandraRepo([]string{"localhost:9042"})
			if err != nil {
				log.Fatalf("error %v", err)
			}
			dbstore = db
		} else {
			cassIPs := strings.Split(cassURL, ",")
			db, err := cassandra.NewCassandraRepo(cassIPs)
			if err != nil {
				log.Fatalf("error %v", err)
			}
			dbstore = db
		}
		return
	}
	pgURL := os.Getenv("PG_URL")
	if pgURL != "" {
		if pgURL == "test" {
			db, err := postgres.NewPostgresRepo("postgresql://tester:Apasswd@localhost:5432?sslmode=disable")
			if err != nil {
				log.Fatalf("error %v", err)
			}
			dbstore = db
		} else {
			db, err := postgres.NewPostgresRepo(pgURL)
			if err != nil {
				log.Fatalf("error %v", err)
			}
			dbstore = db
		}
		return
	}
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
