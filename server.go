package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	pb "student_pb" // import the generated protobuf code
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudentDetails(ctx context.Context, in *pb.StudentRequest) (*pb.StudentResponse, error) {
	return &pb.StudentResponse{Name: "John Doe", Age: 20, Subject: "Mathematics"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, &server{})

	fmt.Println("Server is running on port 50051.")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
