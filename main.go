package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	kirill := &Person{
		Name: "Kirill",
		Age:  24,
		AnotherStruct: &AnotherStruct{
			Field1: "TestString",
			Field2: "TestString2",
		},
	}
	data, err := proto.Marshal(kirill)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}
	fmt.Println(data)

	newKirill := &Person{}
	err = proto.Unmarshal(data, newKirill)
	if err != nil {
		log.Fatal("unmarshall error: ", err)
	}
	fmt.Println(newKirill.GetName())
	fmt.Println(newKirill.GetAge())
	fmt.Println(newKirill.AnotherStruct.GetField1())
	fmt.Println(newKirill.AnotherStruct.GetField2())
}
