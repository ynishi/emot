package main

import (
	"context"
	"fmt"
	pb "github.com/ynishi/emot/user"
	"google.golang.org/grpc"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, _ := grpc.Dial("127.0.0.1:1315",opts...)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	testAuth := &pb.UserAuthRequest{
		Name: "test1",
		Password: "pass1",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, _ := client.Auth(ctx, testAuth)
	fmt.Println(res)

}
