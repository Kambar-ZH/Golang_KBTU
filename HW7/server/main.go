package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"example/api"
	"log"
	"net"
)

const (
	port = ":8080"
)

type UserRepository struct {
	users map[int64]*api.User
	api.UnimplementedUserServiceServer
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[int64]*api.User{
			1: {
				Id: 1,      
				Handle: "Kambar_Z", 
				Country: "Kazakhstan",
				City: "Atyrau",
				Rating: 1703,
				MaxRating: 1703,
				Avatar: "url-link",
			},
		},
	} 
}

func (ur UserRepository) GetAll(ctx context.Context, req *api.Empty) (*api.UserList, error) {
	res := []*api.User{}
	for _, user := range (ur.users) {
		res = append(res, user)
	}
	ans := api.UserList{Users: res}
	return &ans, nil 
}

func (ur UserRepository) Get(ctx context.Context, req *api.UserRequestId) (*api.User, error) {
	if user, ok := ur.users[req.Id]; ok {
		return user, nil
	}

	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("user with id %d does not exist", req.Id))
}

func (ur UserRepository) Insert(ctx context.Context, req *api.User) (*api.User, error) {
	user := &api.User{
		Id: req.Id,
		Handle: req.Handle,
		Country: req.Country,
		City: req.City,
		Rating: req.Rating,
		MaxRating: req.MaxRating,
		Avatar: req.Avatar,
	}

	ur.users[req.Id] = req
	
	return user, nil
}

func (ur UserRepository) Update(ctx context.Context, req *api.User) (*api.User, error) {
	user := &api.User{
		Id: req.Id,
		Handle: req.Handle,
		Country: req.Country,
		City: req.City,
		Rating: req.Rating,
		MaxRating: req.MaxRating,
		Avatar: req.Avatar,
	}

	ur.users[req.Id] = req
	
	return user, nil
}

func (ur UserRepository) Remove(ctx context.Context, req *api.UserRequestId) (*api.Empty, error) {
	if _, ok := ur.users[req.Id]; !ok {
		return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
	}
	delete(ur.users, req.Id)
	return &api.Empty{}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	userRepository := NewUserRepository()

	api.RegisterUserServiceServer(grpcServer, userRepository)

	log.Printf("Serving on %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve on %v: %v", listener.Addr(), err)
	}
}