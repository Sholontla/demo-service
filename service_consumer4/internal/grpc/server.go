package grpc

import (
	pb "service/consumer4/demo/internal/proto"
)

type Server struct {
	pb.ProducerServiceServer
}
