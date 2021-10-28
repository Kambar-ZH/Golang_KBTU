package main

import (
	"context"
	"google.golang.org/grpc"
	"example/api"
	"log"
	"time"
)

const (
	port = ":8080"
)

func main() {
	ctx := context.Background()

	connStartTime := time.Now()
	conn, err := grpc.Dial("localhost" + port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", port, err)
	}
	log.Printf("connected in %d microsec", time.Since(connStartTime))

	userRepositoryClient := api.NewUserServiceClient(conn)
	users, err := userRepositoryClient.GetAll(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Printf("got list of users: %v", users.Users)

	validId, invalidId := 1, 3
	user, err := userRepositoryClient.Get(ctx, &api.UserRequestId{Id: int64(validId)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("got user with id %d: %v", validId, user)

	_, err = userRepositoryClient.Get(ctx, &api.UserRequestId{Id: int64(invalidId)})
	if err != nil {
		log.Printf("got error: %v", err)
	}

	newUser := &api.User{
		Id: 2,      
		Handle: "Yergeldi", 
		Country: "Kazakhstan",
		City: "Ust Kamenogorsk",
		Rating: 2500,
		MaxRating: 3000,
		Avatar: "url-link",
	}

	insertResponseUser, err := userRepositoryClient.Insert(ctx, newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user successfully inserted to db: %v", insertResponseUser)

	updatedUser := &api.User{
		Id: 1,      
		Handle: "Kambar_Z", 
		Country: "Kazakhstan",
		City: "Atyrau",
		Rating: 1800,
		MaxRating: 1800,
		Avatar: "url-link",
	}
	updateResponseUser, err := userRepositoryClient.Update(ctx, updatedUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user successfully updated: %v", updateResponseUser)

	_, err = userRepositoryClient.Remove(ctx, &api.UserRequestId{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
}