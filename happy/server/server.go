package main

import (
	"context"
	"math/rand"
	"net"

	pb "github.com/ynishi/emot/happy"
	"google.golang.org/grpc"
)

type happyServiceServer struct {
	happies []*pb.Happy
}

func (s *happyServiceServer) GetHappyWord(context.Context, *pb.HappyRequest) (*pb.HappyResponse, error) {
	return &pb.HappyResponse{
		"success",
		s.happies[rand.Intn(len(s.happies))],
	}, nil
}

func newServer() *happyServiceServer {
	s := &happyServiceServer{happies: make([]*pb.Happy, 0)}
	s.happies = append(s.happies, &pb.Happy{"happy1"})
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":1315")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHappyServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
