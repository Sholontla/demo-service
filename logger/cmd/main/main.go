package main

import (
	"time"

	"service/logger/service/cmd/grpc"
	shcommands "service/logger/service/cmd/sh_commands"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				shcommands.ShScriptLogUpdate()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	grpc.StreamGrpcServer()
}
