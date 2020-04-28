package main

import (
	pb "grpc-test/proto"
	"log"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var ch = make(chan int, 10)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSimpleServiceClient(conn)
	t1 := time.Now()

	for i := 0; i < 500; i++ {
		go send(client, i)
	}
	for i := 0; i < 500; i++ {
		<-ch
	}

	log.Printf("Time: %d s", time.Now().Unix()-t1.Unix())

}

func send(client pb.SimpleServiceClient, thread int) {
	for i := 0; i < 2000; i++ {
		r, err := client.Route(context.Background(), &pb.SimpleRequest{Data: strconv.Itoa(i * thread)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Value)
	}
	ch <- thread
}
