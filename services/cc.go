package services

import (
	"fmt"
	"os/exec"
)

func ccompile(source, dialect string) ([]byte, error) {
	var std string
	switch dialect {
	case "ansi":
		std = "-ansi"
	case "c89":
		std = "-std=c89"
	case "c90":
		std = "-std=c90"
	case "c99":
		std = "-std=c99"
	case "c11":
		std = "-std=c11"
	}
	out, err = exec.Command(cc, std, "-pedantic", "-Werror", "-static",
		"-Wall", "-pipe", "-fPIC", "-o", fmt.Sprintf("%s.out", source),
		source).CombinedOutput()
	return out, err
}
