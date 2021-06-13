package cassandra

import (
	"context"
	"github.com/gocql/gocql"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type cassandraRepo struct {
	session *gocql.Session
}

var ctx = context.Background()

func newCassandraSession(hosts []string) (*gocql.Session, error) {
	cluster := gocql.NewCluster()
	cluster.Hosts = hosts
	cluster.Keyspace = "lister"
	cluster.Consistency = gocql.Quorum
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

func (r *cassandraRepo) RetrieveAll() (*pb.Lists, error) {
	var id uint32
	var title string
	var items []string
	var lists []*pb.List

	scanner := r.session.Query(listRetrieveAllQuery).WithContext(ctx).Iter().Scanner()
	for scanner.Next() {
		err := scanner.Scan(&id, &title, &items)
		if err != nil {
			return nil, err
		}
		list := &pb.List{
			Id:    id,
			Title: title,
			Items: items,
		}
		lists = append(lists, list)
	}

	return &pb.Lists{Lists: lists}, nil
}

func (r *cassandraRepo) Retrieve(list *pb.List) (*pb.List, error) {
	var id uint32
	var title string
	var items []string

	err := r.session.Query(listRetrievalQuery).WithContext(ctx).Consistency(gocql.One).Scan(&id, &title, &items)
	if err != nil {
		return nil, err
	}

	return &pb.List{Id: id, Title: title, Items: items}, nil
}

func (r *cassandraRepo) Save(list *pb.List) (*pb.Lists, error) {
	var id uint8
	err := r.session.Query(listMaxIdQuery).WithContext(ctx).Consistency(gocql.One).Scan(&id)
	if err != nil {
		return nil, err
	}
	id = id + 1
	err = r.session.Query(listInsertQuery, id, list.Title, list.Items).WithContext(ctx).Exec()
	if err != nil {
		return nil, err
	}
	return r.RetrieveAll()
}
