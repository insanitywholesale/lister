package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type postgresRepo struct {
	client *pgx.Conn
	pgURL  string
}

var ctx = context.Background()

func newPostgresClient(url string) (*pgx.Conn, error) {
	client, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(ctx, createListTableQuery)
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

func (r *postgresRepo) RetrieveAll() (*pb.Lists, error) {
	var list = &pb.List{}
	var listslice []*pb.List

	rows, err := r.client.Query(ctx, listRetrieveAllQuery)
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

func (r *postgresRepo) Retrieve(list *pb.List) (*pb.List, error) {
	row, err := r.client.Query(ctx, listRetrievalQuery, list.Id)
	if err != nil {
		return nil, err
	}
	err = row.Scan(
		&list.Id,
		&list.Title,
		&list.Items,
	)
	return list, nil
}

func (r *postgresRepo) Save(list *pb.List) (*pb.Lists, error) {
	var id uint32
	err := r.client.QueryRow(ctx, listInsertQuery,
		list.Title,
		list.Items,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	list.Id = id
	return r.RetrieveAll()
}
