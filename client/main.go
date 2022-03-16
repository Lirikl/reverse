package main
import (
	"context"
	"log"
	"time"
	"github.com/lirikl/reverse"
	"google.golang.org/grpc"
	)
	func main() {
		log.Println("Client running ...")
		conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock(
		))
		if err != nil {
		log.Fatalln(err)
		}
		defer conn.Close()
		client := reverse.NewReverseClient(conn)
		request := &reverse.Request{Message: "This is a test"}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second
		)
		defer cancel()
		response, err := client.DoReverse(ctx, request)
		if err != nil {
		log.Fatalln(err)
		}
		log.Println("Response:", response.GetMessage())
}