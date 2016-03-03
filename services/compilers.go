package services

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

var cc string
var cpp string

var nasm string
var gas string

var gfortran string

func Scan() {
	var err error

	cc, err = exec.LookPath("cc")
	if err != nil {
		log.Fatal("Could not find C Compiler.")
	}
	log.Printf("Found C Compiler at %s\n", cc)

	cpp, err = exec.LookPath("c++")
	if err != nil {
		log.Fatal("Could not find C++ Compiler.")
	}
	log.Printf("Found C++ Compiler at %s\n", cpp)

	nasm, err = exec.LookPath("nasm")
	if err != nil {
		log.Fatal("Could not find NASM Assembler.")
	}
	log.Printf("Found NASM Assembler at %s\n", nasm)

	gas, err = exec.LookPath("as")
	if err != nil {
		log.Fatal("Could not find GNU Assembler.")
	}
	log.Printf("Found GNU Assembler at %s\n", cc)

	gfortran, err = exec.LookPath("gfortran")
	if err != nil {
		log.Fatal("Could not find GNU Fortran Compiler.")
	}
	log.Printf("Found GNU Fortran Compiler at %s\n", gfortran)
}
