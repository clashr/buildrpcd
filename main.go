package main

import (
	"net"
	"net/rpc"

	log "github.com/Sirupsen/logrus"

	"github.com/clashr/buildrpcd/api"
	"github.com/clashr/buildrpcd/services"
)

func init() {
	services.Scan()
}

func main() {
	builder := new(api.Build)

	server := rpc.NewServer()
	server.Register(builder)

	//start listening for messages on port 1234
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", err)
	}

	log.Println("Serving RPC handler")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error serving: %s", err)
		}

		go server.ServeConn(conn)
	}

}
