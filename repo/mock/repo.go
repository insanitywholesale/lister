package mock

import (
	"errors"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"strconv"
)

var testlist = &pb.List{
	Id:    1,
	Title: "Future Optiplexes",
	Items: []string{
		"http://www.proshop.gr/index.php?route=product/product&product_id=143",
		"http://www.proshop.gr/index.php?route=product/product&product_id=62",
	},
}

type listrepo []*pb.List

var testlists listrepo = []*pb.List{
	testlist,
	&pb.List{
		Id:    2,
		Title: "Git forges",
		Items: []string{
			"https://gitlab.com/insanitywholesale",
			"https://github.com/insanitywholesale",
			"http://gitnas.hell:3000/inherently",
		},
	},
}

var listId uint32 = 3

func NewMockRepo() (listrepo, error) {
	return testlists, nil
}

func (listrepo) RetrieveAll() (*pb.Lists, error) {
	return &pb.Lists{Lists: testlists}, nil
}

func (listrepo) Retrieve(list *pb.List) (*pb.List, error) {
	id := list.Id
	for _, l := range testlists {
		if id == l.Id {
			return l, nil
		}
	}
	return nil, errors.New("no list with ID " + strconv.Itoa(int(id)) + " was found")
}

func (listrepo) Save(list *pb.List) (*pb.Lists, error) {
	list.Id = listId
	listId = listId + 1
	testlists = append(testlists, list)
	return &pb.Lists{Lists: testlists}, nil
}
