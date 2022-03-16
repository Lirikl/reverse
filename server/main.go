package main


import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/lirikl/reverse"
	"google.golang.org/grpc"
)

type server struct {
reverse.UnimplementedReverseServer
}
func (s *server) DoReverse(ctx context.Context, request *reverse.Request
) (*reverse.Response, error) {
log.Println(fmt.Sprintf("Request: %s", request.GetMessage()))
return &reverse.Response{Message: fmt.Sprintf("Reversed string %s",
stringReverse(request.GetMessage()))}, nil
}

func main() {
lis, err := net.Listen("tcp", ":9000")
if err != nil {
log.Fatalf("failed to listen: %v", err)
}
srv := grpc.NewServer()
reverse.RegisterReverseServer(srv, &server{})
log.Fatalln(srv.Serve(lis))
}
func stringReverse(str string) (result string) {
for _, v := range str {
result = string(v) + result
}
return
}
		