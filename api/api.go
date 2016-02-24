package api

import (
	"github.com/clashr/buildrpcd/services"
	"log"
)

type Args struct {
	Language, Dialect, Contents string
}

type Build struct {
}

type Result struct {
	CCOut, Binary string
}

func (b *Build) Compile(args Args, result *Result) (err error) {
	return Compile(args, result)
}

func Compile(args Args, result *Result) (err error) {
	log.Printf("Reached Compile Endpoint\n")
	services.Nothing()
	*result = Result{"not yet implemented", "binary data"}
	return nil
}
