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

	fileName := fmt.Sprintf("source%d", time.Now().Nanosecond)
	err = ioutil.WriteFile(filepath.Join(tmpDir, fileName), []byte(contents), 0644)
	if err != nil {
		return
	}

	log.Printf("Reached Build/Make endpoint.\n")
	log.Printf("Compiling %s source.\n", language)
	switch language {
	case "c":
		ccompile(fileName, dialect)
	case "cpp":
		cppcompile(fileName, dialect)
	case "asm":
		asmcompile(fileName, dialect)
	case "fortran":
		fcompile(fileName, dialect)
	}

	err = os.Chdir(origDir)
	if err != nil {
		return
	}
	return
}
