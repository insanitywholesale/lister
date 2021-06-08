package grpc

import (
	pb "gitlab.com/insanitywholesale/lister"
)

type Server struct {
	pb.UnimplementedListerServer
}
