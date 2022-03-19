package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Lirikl/mafia/pkg/proto/mafia"
	"google.golang.org/grpc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Client running ...")
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := pb.NewMafiaClient(conn)
	request := &pb.Request{Message: "This is a test"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Do(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	ctx = context.Background()
	stream, err := client.Connect(ctx)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]
	stream.Send(pb.ConnectionUpdate{Name: name, Connect: pb.ConnectionStatus_Connect})
	go func() {
		for {
			in, err2 := stream.Rcv()
			if in.Connect == pb.ConnectionStatus_Start {
				role := in.Role
				game, err2 := client.Game(ctx)
				game.Send(pb.GameCommand{SessionID: in.SessionID}) =
			}
		}
	}()
	log.Println("Response:", response.GetMessage())
}
