package cassandra

import (
	"github.com/gocql/gocql"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type cassandraRepo struct {
	session *gocql.Session
}

func NewCassandraRepo() (*gocql.Session, error) {
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "lister"
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (cassandraRepo) RetrieveAll() (*pb.List, error) {
	return &pb.List{}, nil
}

func (cassandraRepo) Retrieve(list *pb.List) (*pb.List, error) {
	return &pb.List{}, nil
}

func (cassandraRepo) Save(list *pb.List) (*pb.Lists, error) {
	return &pb.Lists{}, nil
}
