package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
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

	// TODO: Return Compiled File Data

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
