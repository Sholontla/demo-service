package grpc

import (
	"fmt"
	"io"
	"log"
	"net"

	"service/logger/service/logger"
	pb "service/logger/service/proto/loggs"

	"google.golang.org/grpc"
)

type Server struct {
	pb.LoggsProducerServiceServer
}

func (s *Server) StreamProducer(stream pb.LoggsProducerService_StreamProducerServer) error {

	log.Println("Stream server invoked ....")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		err = stream.Send(&pb.MainLoggerResponse{
			ResponseLogger: "LOG saved",
		})

		l := fmt.Sprintf("%s: %s", req.LoggerId, req.LoggerMessage)
		log.Println(req)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		logger.LogPersonal(l)
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}

var addr string = "logger_backend:50010"

func StreamGrpcServer() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %s\n: ", err)
	}

	log.Printf("GRPC Server Listening on %s\n: ", addr)

	s := grpc.NewServer()

	pb.RegisterLoggsProducerServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve on %s\n: ", err)
	}
}
