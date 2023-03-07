package grpc

import (
	pb "service/consumer5/demo/internal/proto"
)

type Server struct {
	pb.ProducerServiceServer
}
