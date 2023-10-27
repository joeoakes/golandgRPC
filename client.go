package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "student_pb" // import the generated protobuf code
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStudentServiceClient(conn)

	response, err := c.GetStudentDetails(context.Background(), &pb.StudentRequest{Id: 1})
	if err != nil {
		log.Fatalf("Could not get student details: %v", err)
	}

	fmt.Printf("Student Details: Name: %s, Age: %d, Subject: %s\n", response.Name, response.Age, response.Subject)
}
