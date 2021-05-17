package main

import (
	"fmt"
	"log"

	greeting "github.com/anitabee/exploring-go/hello-protobuf/greeting"
	"github.com/golang/protobuf/proto"
)

func main() {

	message := &greeting.Message{
		Greet: "Ello there",
	}

	// marshal data
	data, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("Raw protobuf object: %v\n", data)

	// unmarshal byte array into an object
	newMessage := &greeting.Message{}
	err = proto.Unmarshal(data, newMessage)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newMessage.GetGreet())

}
