package main

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"

	"net"

	pb "github.com/ynishi/emot/happy"
	"google.golang.org/grpc"
)

type happyServiceServer struct {
	db *leveldb.DB
	m  sync.Mutex
}

func (s *happyServiceServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {

	iter := s.db.NewIterator(util.BytesPrefix([]byte(req.Query)), nil)
	for iter.Next() {
		return &pb.GetResponse{
			Word: string(iter.Value()),
		}, nil
	}
	iter.Release()
	err := iter.Error()
	return nil, err
}

func (s *happyServiceServer) Post(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	err := s.db.Put([]byte(req.Key), []byte(req.Word), nil)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{
		Key: req.Key,
	}, nil
}

func (s *happyServiceServer) ListTag(req *pb.ListTagRequest, stream pb.HappyService_ListTagServer) error {
	s.m.Lock()
	defer s.m.Unlock()
	iter := s.db.NewIterator(util.BytesPrefix([]byte(req.Query)), nil)
	tmp := make(map[string]bool)
	for iter.Valid() {
		resp := pb.ListTagResponse{}
		if err := json.Unmarshal(iter.Value(), &resp); err != nil {
			return err
		}
		if _, ok := tmp[resp.Tag]; !ok {
			if err := stream.Send(&resp); err != nil {
				return err
			}
			tmp[resp.Tag] = true
		}
		iter.Next()
	}
	iter.Release()
	err := iter.Error()
	return err
}

func newServer(db *leveldb.DB) *happyServiceServer {
	initdb(db)
	s := &happyServiceServer{db: db}
	return s
}

func initdb(db *leveldb.DB) error {
	o := &pb.ListTagResponse{
		Tag: "tag1",
		Num: 1,
	}
	v, _ := json.Marshal(o)
	db.Put([]byte("key"), []byte(v), nil)
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
