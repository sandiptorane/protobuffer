package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main(){
	sandip := &Person{
		Name: "sandip",
		Age:  22,
		SocialFollowers: &SocialFollowers{
			Youtube:       1000,
			Twitter:       500,
		},
	}
	//marshal sandip it returns []byte result
	data,err := proto.Marshal(sandip)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newSandip := &Person{}
	err = proto.Unmarshal(data, newSandip)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our  object
	// for good measure
	fmt.Println("age:",newSandip.GetAge())
	fmt.Println("name:",newSandip.GetName())
	fmt.Println("youtube subscribers:",newSandip.SocialFollowers.GetYoutube())
	fmt.Println("twitter followers:",newSandip.SocialFollowers.GetTwitter())
	newSandip.SocialFollowers.Reset()
	fmt.Println("after resetting  socialFollowers=",newSandip)
	newSandip.Reset()
	fmt.Println("after resetting  person=",newSandip)

}
