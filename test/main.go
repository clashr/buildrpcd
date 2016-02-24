package main

import (
	"log"
	"net/rpc"

	"github.com/clashr/buildrpcd/api"
)

func main() {
	//make connection to rpc server
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &api.Args{
		Language: "c++",
		Contents: "#include<iostream>\nusing namespace std;\nint main() { return 0; }",
	}
	//this will store returned result
	var result api.Result
	//call remote procedure with args
	log.Printf("%s ", args.Language)
	err = client.Call("Build.Compile", args, &result)
	if err != nil {
		log.Fatalf("error in Build", err)
	}
	//we got our result in result
	log.Printf("Language:%s\n %s\n Result: %s", args.Language, args.Contents, result)
}
