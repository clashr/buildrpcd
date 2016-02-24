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
	CCOut, Binary []byte
}

func (b *Build) Compile(args Args, result *Result) (err error) {
	return Compile(args, result)
}

func Compile(args Args, result *Result) (err error) {
	log.Printf("Reached Compile Endpoint\n")
	_, output, err := services.Make(args.Contents, args.Language, args.Dialect)
	*result = Result{output, []byte("binary data")}
	return nil
}
