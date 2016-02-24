package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/clashr/buildrpcd/api"
	"github.com/clashr/buildrpcd/services"
)

func init() {
	services.Scan()
}

func main() {
	builder := new(api.Build)
	err := rpc.Register(builder)
	if err != nil {
		log.Fatalf("Format of service builder isn't correct. %s", err)
	}
	rpc.HandleHTTP()
	//start listening for messages on port 1234
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", err)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
