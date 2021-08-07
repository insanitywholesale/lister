package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	//"github.com/georgysavva/scany/pgxscan"
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
/*
	var listslice []*pb.List
	pgxscan.Select(ctx, r.client, &listslice, listRetrieveAllQuery)
	return &pb.Lists{Lists: listslice}, nil
*/
	var listslice = []*pb.List{}
	rows, err := r.client.Query(ctx, listRetrieveAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var list = &pb.List{}
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
	if err != nil {
		return nil, err
	}
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

func (r *postgresRepo) Remove(list *pb.List) (*pb.Lists, error) {
	_, err := r.client.Query(ctx, listDeleteQuery, list.Id)
	if err != nil {
		return nil, err
	}
	return r.RetrieveAll()
}
