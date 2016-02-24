package services

import (
	"fmt"
	"os/exec"
)

func ccompile(source, dialect string) (out []byte, err error) {
	var standard string
	switch dialect {
	case "ansi":
		standard = "-ansi"
	case "c89":
		standard = "-std=c89"
	case "c90":
		standard = "-std=c90"
	case "c99":
		standard = "-std=c99"
	case "c11":
		standard = "-std=c11"
	}
	out, err = exec.Command(cc, standard, "-pedantic", "-Werror", "-Wall",
		"-pipe", "-fPIC", fmt.Sprintf("-o %s.out", source),
		source).CombinedOutput()
	return
}
