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
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
)

func Nothing() {
	log.Printf("Reached services.Nothing endpoint.\n")
	log.Printf("By the way, you can find a C/C++ Compiler here: %s\n", cc)
}

func Make(contents, language, dialect string) (bindata, ccout []byte, err error) {
	origDir, err := os.Getwd()
	if err != nil {
		return
	}
	tmpDir := os.TempDir()

	err = os.Chdir(tmpDir)
	if err != nil {
		return
	}

	log.Printf("Reached Build/Make endpoint.\n")
	log.Printf("Compiling %s source.\n", language)

	var fileName string
	switch language {
	case "c":
		fileName, err = writeSource(contents, tmpDir, ".c")
		if err != nil {
			break
		}
		ccout, err = ccompile(fileName, dialect)
	case "cpp":
		fileName, err = writeSource(contents, tmpDir, ".cpp")
		if err != nil {
			break
		}
		ccout, err = cppcompile(fileName, dialect)
	case "asm":
		fileName, err = writeSource(contents, tmpDir, ".asm")
		if err != nil {
			break
		}
		ccout, err = asmcompile(fileName, dialect)
	case "fortran":
		fileName, err = writeSource(contents, tmpDir, ".f")
		if err != nil {
			break
		}
		ccout, err = fcompile(fileName, dialect)
	}
	if err != nil {
		return
	}

	binary := fmt.Sprintf("%s.out", fileName)

	strip(binary)

	bindata, err = ioutil.ReadFile(binary)
	if err != nil {
		return
	}

	err = os.Remove(fileName)
	if err != nil {
		return
	}
	err = os.Remove(binary)
	if err != nil {
		return
	}
	err = os.Chdir(origDir)
	if err != nil {
		return
	}
	return
}

func writeSource(source, location, ext string) (fileName string, err error) {
	fileName = fmt.Sprintf("source%d%s", time.Now().Nanosecond, ext)
	err = ioutil.WriteFile(filepath.Join(location, fileName), []byte(source), 0644)
	if err != nil {
		return
	}
	return
}

func strip(file string) error {
	err := exec.Command("strip", "-s", file).Run()
	return err
}
