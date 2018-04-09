package main

import (
	"context"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"

	"net"

	pb "github.com/ynishi/emot/happy"
	"google.golang.org/grpc"
)

type happyServiceServer struct {
	db *leveldb.DB
}

func (s *happyServiceServer) Get(ctx context.Context, req *pb.HappyRequest) (*pb.HappyResponse, error) {

	iter := s.db.NewIterator(util.BytesPrefix([]byte(req.Query)), nil)
	for iter.Next() {
		return &pb.HappyResponse{
			Word: string(iter.Value()),
		}, nil
	}
	iter.Release()
	err := iter.Error()
	return nil, err
}

func newServer(db *leveldb.DB) *happyServiceServer {
	initdb(db)
	s := &happyServiceServer{db: db}
	return s
}

func initdb(db *leveldb.DB) error {
	db.Put([]byte("key"), []byte("value"), nil)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":1314")
	if err != nil {
		panic(err)
	}
	db, err := leveldb.OpenFile("words.db", nil)
	grpcServer := grpc.NewServer()
	pb.RegisterHappyServiceServer(grpcServer, newServer(db))
	grpcServer.Serve(lis)
}
