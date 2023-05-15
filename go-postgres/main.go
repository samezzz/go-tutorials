package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/samezzz/go-postgres/api"
	"github.com/samezzz/go-postgres/storage"
)

func main() {
	listenAddress := flag.String("listenAddress", ":8000", "the server address")
	flag.Parse()

	godotenv.Load()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddress, store)
	fmt.Println("server running on port", *listenAddress)
	log.Fatal(server.Start())
}
