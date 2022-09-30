package main

import (
	grpcclients "api/grpcClients"
	"api/server"
	"log"
)

func main() {
	grpcclients.Init()
	server.Init()
	defer grpcclients.GetAccountServiceClient().Conn.Close()

	log.Fatal(server.Run())
}
