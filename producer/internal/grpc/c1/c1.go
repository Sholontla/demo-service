package c1

import (
	"context"
	"fmt"
	"io"
	"log"
	"service/producer/demo/internal/models"
	pb "service/producer/demo/internal/proto"
	"time"
)

func StreamConsumerClient1(c pb.ProducerServiceClient, requestChan <-chan models.MainProducer) {
	log.Println("gRPC Service 1 Client invoked ...")

	stream, err := c.StreamProducer(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}
	waitc := make(chan struct{})

	go func() {
		for request := range requestChan {
			req := &pb.MainProducerRequest{
				ProducerId:           request.ProducerID,
				ProducerServiceArea:  request.ProducerServiceArea,
				ProducerCreatedAt:    request.ProducerCreatedAt,
				ProducerMessageId:    request.ProducerMessage.ProducerMessageID,
				Host:                 request.ProducerMessage.Host,
				Client:               request.ProducerMessage.Client,
				Ip:                   request.ProducerMessage.Ip,
				Port:                 request.ProducerMessage.Port,
				InformationCreatedAt: request.ProducerMessage.CreatedAt,
				DataProducerId:       request.ProducerMessage.DataProducer.DataProducerID,
				Product:              request.ProducerMessage.DataProducer.Product,
				Name:                 request.ProducerMessage.DataProducer.Name,
				Category:             request.ProducerMessage.DataProducer.Category,
				SubCategory:          request.ProducerMessage.DataProducer.SubCategory,
				Price:                request.ProducerMessage.DataProducer.Price,
				Quantity:             request.ProducerMessage.DataProducer.Quantity,
				Supplier:             request.ProducerMessage.DataProducer.Supplier,
				Description:          request.ProducerMessage.DataProducer.Description,
				Gender:               request.ProducerMessage.DataProducer.Gender,
			}

			start := time.Now()
			log.Printf("Send Request: %v\n", req)
			stream.Send(req)
			elapsed := time.Since(start)
			log.Printf("Address Channel took %s", elapsed)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while reciving: %v\n", err)
				break
			}

			p := res

			fmt.Println("Recived: ", p)
		}

		close(waitc)
	}()

	<-waitc
}
