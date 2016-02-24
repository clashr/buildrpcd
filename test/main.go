package main

import (
	"io/ioutil"
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
		Language: "c",
		Dialect:  "ansi",
		Contents: "#include<stdio.h>\nint main() {\nprintf(\"Hello World\");\nreturn 0; }",
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
	log.Printf("Language:%s\n %s\nResult: %d", args.Language, args.Contents, len(result.Binary))

	ioutil.WriteFile("a.out", result.Binary, 0755)
	if err != nil {
		log.Fatalf("could not write result")
	}
	log.Printf("Wrote Binary")
}
