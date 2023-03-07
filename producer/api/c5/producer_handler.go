package c5

import (
	"fmt"
	"log"
	"sync"

	"service/producer/demo/config"
	c5 "service/producer/demo/internal/grpc/c5"
	"service/producer/demo/internal/models"
	pb "service/producer/demo/internal/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var wg sync.WaitGroup

func SendDataServiceConsumer5(ctxF *fiber.Ctx) error {

	grpc_host, grpc_port, grpc_tls := config.GRPCConfig5()
	add := fmt.Sprintf("%s:%s", grpc_host, grpc_port)
	cert_file := config.GRPCPathsConfig()

	opts := []grpc.DialOption{}

	if grpc_tls {
		creds, err := credentials.NewClientTLSFromFile(cert_file, add)
		if err != nil {
			log.Fatalf("Error While loading client cert file: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	var request models.MainProducer

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	//conn, err := grpc.Dial(add, opts...)
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect Client: %v\n", err)
	}
	defer conn.Close()

	log.Printf("GRPC Client Listening on %s\n: ", grpc_port)

	c := pb.NewProducerServiceClient(conn)

	requestChan := make(chan models.MainProducer)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c5.StreamConsumerClient5(c, requestChan)
	}()
	requestChan <- request
	close(requestChan)

	wg.Wait()

	return ctxF.JSON(fiber.Map{"store": request})
}
