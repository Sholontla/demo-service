package grpc

import (
	pb "service/consumer2/demo/internal/proto"
)

type Server struct {
	pb.ProducerServiceServer
}
