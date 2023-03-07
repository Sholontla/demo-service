package grpc

import (
	pb "service/consumer1/demo/internal/proto"
)

type Server struct {
	pb.ProducerServiceServer
}
