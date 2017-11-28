// protobuftest
package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

/*
Create .proto file.
Run protocol buffer compiler -> protoc --go_out=. *.proto
try to open generated .go file,
Output is file with generated API
*/
func testProtobuf() {

	s := "hello"
	t := &Test{
		Label: s,
		Type:  1234,
	}

	fmt.Println("struct is ", t)

	data, err := proto.Marshal(t)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("new struct is", newTest)

	// Now test and newTest contain the same data.
	if t.GetLabel() != t.GetLabel() {
		log.Fatalf("data mismatch %q != %q", t.GetLabel(), t.GetLabel())
	}

}
