package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

	pb "30lesson/pb"
)

func main() {
	user := &pb.User{
		Id:    123,
		Name:  "Konstantin",
		Email: "test@example.com",
	}

	fmt.Printf("original: %+v\n", user)

	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("binary: %v\n", data)
	fmt.Printf("size: %d bytes\n", len(data))

	var decoded pb.User

	if err := proto.Unmarshal(data, &decoded); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("decoded: %+v\n", &decoded)

	fmt.Println("id:", decoded.GetId())
	fmt.Println("name:", decoded.GetName())
	fmt.Println("email:", decoded.GetEmail())
}
