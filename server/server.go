package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"

	pb "github.com/Sean-Miningah/usermanagement-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		// user_list: &pb.UserList{},

	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to lister: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	readBytes, err := os.ReadFile("users.json")
	var users_list *pb.UserList = &pb.UserList{}
	var user_id int32 = rand.Int31n(100)
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found. Creating a new file")
			users_list.Users = append(users_list.Users, created_user)
			jsonBytes, err := protojson.Marshal(users_list)
			if err != nil {
				log.Fatal("JSON Marshalling error: ", err)
			}
			if err := os.WriteFile("users.json", jsonBytes, 0644); err != nil {
				log.Fatalf("Failed to write to file: %v", err)
			}
			return created_user, nil
		} else {
			log.Fatalln("Error reading file: ", err)
		}
	}

	if err := protojson.Unmarshal(readBytes, users_list); err != nil {
		log.Fatalln("Failed to parse user list: ", err)
	}
	users_list.Users = append(users_list.Users, created_user)
	jsonBytes, err := protojson.Marshal(users_list)
	if err != nil {
		log.Fatal("JSON Marshalling error: ", err)
	}
	if err := os.WriteFile("users.json", jsonBytes, 0644); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// s.user_list.Users = append(s.user_list.Users, created_user)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	jsonBytes, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}
	var users_list *pb.UserList = &pb.UserList{}
	if err := protojson.Unmarshal(jsonBytes, users_list); err != nil {
		log.Fatalf("Unmarshaling failed: %v", err)
	}
	// return s.user_list, nil

	return users_list, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
