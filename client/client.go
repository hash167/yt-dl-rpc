package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hash167/yt-dl-rpc/api/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:3033"
	defaultName = "https://www.youtube.com/watch?v=Q2YKXiTRcBE"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRPCServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Exec(ctx, &pb.DownloadRequest{
		URL: name, Path: "/downloads", Rename: "hashtag.webm"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r)
}
