package cassandra

import (
	"github.com/gocql/gocql"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type cassandraRepo struct {
	session *gocql.Session
}

func newCassandraSession(hosts []string) (*gocql.Session, error) {
	cluster := gocql.NewCluster()
	cluster.Hosts = hosts
	cluster.Keyspace = "lister"
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func NewCassandraRepo(hosts []string) (*cassandraRepo, error) {
	cassandraSession, err := newCassandraSession(hosts)
	if err != nil {
		return nil, err
	}
	repo := &cassandraRepo{
		session: cassandraSession,
	}
	return repo, nil
}

//TODO: implement
func (cassandraRepo) RetrieveAll() (*pb.Lists, error) {
	return &pb.Lists{}, nil
}

//TODO: implement
func (cassandraRepo) Retrieve(list *pb.List) (*pb.List, error) {
	return &pb.List{}, nil
}

//TODO: implement
func (cassandraRepo) Save(list *pb.List) (*pb.Lists, error) {
	return &pb.Lists{}, nil
}
