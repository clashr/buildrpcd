//
// Copyright (c) 2016 Dennis Chen
//
// This file is part of Clashr.
//
// Clashr is free software: you can redistribute it and/or modify it under the
// terms of the GNU Affero General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// Clashr is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License for
// more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Clashr.  If not, see <http://www.gnu.org/licenses/>.
//

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
