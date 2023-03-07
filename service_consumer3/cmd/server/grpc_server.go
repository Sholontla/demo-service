package server

import (
	"service/consumer3/demo/internal/grpc"
	"sync"
)

var wg sync.WaitGroup

func GRPCProducerServer() {

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpc.GrpcProductServer()
	}()
	wg.Wait()
}
