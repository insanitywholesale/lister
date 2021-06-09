package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"log"
)

type postgresRepo struct {
	client *sql.DB
	pgURL  string
}

func newPostgresClient(url string) (*sql.DB, error) {
	client, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createListTableQuery)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewPostgresRepo(url string) (*postgresRepo, error) {
	pgclient, err := newPostgresClient(url)
	if err != nil {
		return nil, err
	}
	repo := &postgresRepo{
		pgURL:  url,
		client: pgclient,
	}
	return repo, nil
}

func(r *postgresRepo) RetrieveAll() (*pb.Lists, error) {
	var list = &pb.List{}
	var listslice []*pb.List

	rows, err := r.client.Query(listRetrieveAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&list.Id,
			&list.Title,
			&list.Items,
		)
		if err != nil {
			return nil, err
		}
		listslice = append(listslice, list)
	}
	return &pb.Lists{Lists: listslice}, nil
}

func(r *postgresRepo) Retrieve(l *pb.List) (*pb.List, error) {
	row, err := r.client.Query(listRetrievalQuery, l.Id)
	if err != nil {
		return nil, err
	}
	err = row.Scan(
		&list.Id,
		&list.Title,
		&list.Items
	)
	return list, nil
}

func(r *postgresRepo) Save(list *pb.List) (*pb.Lists, error) {
	var id int
	err := r.client.QueryRow(listInsertQuery,
		list.Title,
		list.Items,
	).Scan(&id)
	list.Id = id
	return r.RetrieveAll()
}
