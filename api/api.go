package api

import (
	"log"

	"github.com/clashr/buildrpcd/services"
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
	binary, output, err := services.Make(args.Contents, args.Language, args.Dialect)
	if err != nil {
		return
	}
	*result = Result{output, binary}
	return nil
}
