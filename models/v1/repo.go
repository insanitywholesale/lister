package models

import (
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type ListsRepo interface {
	RetrieveAll() (*pb.Lists, error)
	Retrieve(*pb.List) (*pb.List, error)
	Save(*pb.List) (*pb.Lists, error)
	Remove(*pb.List) (*pb.Lists, error)
}
