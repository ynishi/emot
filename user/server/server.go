package main

import (
	"context"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"net"

	pb "github.com/ynishi/emot/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"golang.org/x/crypto/bcrypt"
)

type userServiceServer struct {
	db *leveldb.DB
}

func (s *userServiceServer) Auth(ctx context.Context, uar *pb.UserAuthRequest) (*pb.UserResponse, error) {
	if uar == nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "user is required")
	}
	buf, err := s.db.Get([]byte(uar.Name),nil)
	if err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "not found")
	}
	user := pb.User{}
	err = json.Unmarshal(buf, &user)
	if err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "invalid data")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(uar.Password))
	if err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "invalid data")
	}
	return &pb.UserResponse{Name: user.Name}, nil

}

func (s *userServiceServer) Add(ctx context.Context, upr *pb.UserPostRequest) (*pb.UserResponse, error) {
	if upr == nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "user is required")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(upr.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	user := upr.User
	user.Hash = string(hash)
	buf, err := json.Marshal(user)
	if err := s.db.Put([]byte(user.Name),[]byte(buf),nil); err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.UserResponse{Name: user.Name}, nil
}

func (s *userServiceServer) Get(ctx context.Context, ugr *pb.UserGetRequest) (*pb.User, error) {
	if ugr == nil {
		return &pb.User{}, status.Error(codes.InvalidArgument, "user is required")
	}
	buf, err := s.db.Get([]byte(ugr.Name),nil)
	if err != nil {
		return &pb.User{}, status.Error(codes.InvalidArgument, "not found")
	}
	user := pb.User{}
	err = json.Unmarshal(buf, &user)
	if err != nil {
		return &pb.User{}, status.Error(codes.InvalidArgument, "invalid data")
	}
	return &user, nil
}

func (s *userServiceServer) Update(ctx context.Context, upr *pb.UserPostRequest) (*pb.UserResponse, error) {
	if upr == nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "user is required")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(upr.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	user := upr.User
	user.Hash = string(hash)
	buf, err := json.Marshal(user)
	if err := s.db.Put([]byte(user.Name),[]byte(buf),nil); err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.UserResponse{Name: user.Name}, nil
}

func (s *userServiceServer) Delete(ctx context.Context, upr *pb.UserPostRequest) (*pb.UserResponse, error) {
	if upr == nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, "user is required")
	}
	err := s.db.Delete([]byte(upr.User.Name), nil)
	if err != nil {
		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.UserResponse{Name: upr.User.Name}, nil
}

func newServer(db *leveldb.DB) *userServiceServer {
	s := &userServiceServer{
		db: db,
	}
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":1315")
	if err != nil {
		panic(err)
	}

	db, _ := leveldb.OpenFile("user.db",nil)
	defer db.Close()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, newServer(db))
	grpcServer.Serve(lis)
}
