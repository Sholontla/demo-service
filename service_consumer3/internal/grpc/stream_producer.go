package grpc

import (
	"fmt"
	"io"
	"log"
	pb "service/consumer3/demo/internal/proto"
	"time"

	"github.com/google/uuid"
)

func (s *Server) StreamProducer(stream pb.ProducerService_StreamProducerServer) error {
	log.Println("Stream server invoked ....")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		fmt.Println("Comming from StreamInventory req: ", req)

		err = stream.Send(&pb.MainProducerResponse{
			ProducerId:       uuid.New().String(),
			HistoryMessage:   "client 123",
			HistoryCreatedAt: time.Now().Local().String(),
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}

}
