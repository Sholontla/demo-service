package grpc

import (
	pb "service/consumer3/demo/internal/proto"
)

type Server struct {
	pb.ProducerServiceServer
}
