package api

import (
	log "github.com/Sirupsen/logrus"

	"github.com/clashr/buildrpcd/services"
)

func compile(args Args, result *Result) error {
	log.Printf("Reached Compile Endpoint\n")
	binary, output, err := services.Make(args.Contents, args.Language, args.Dialect)
	if err != nil {
		return err
	}
	*result = Result{output, binary}
	return nil
}
