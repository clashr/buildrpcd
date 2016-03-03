package services

import (
	"fmt"
	"os/exec"
)

func cppcompile(source, dialect string) ([]byte, error) {
	var std string
	switch dialect {
	case "ansi":
		std = "-ansi"
	case "c++98":
		std = "-std=c++98"
	case "c++03":
		std = "-std=c++03"
	case "c++11":
		std = "-std=c++11"
	case "c++14":
		std = "-std=c++14"
	}
	out, err := exec.Command(cpp, std, "-pedantic", "-Werror", "-static",
		"-Wall", "-pipe", "-fPIC", "-o", fmt.Sprintf("%s.out", source),
		source).CombinedOutput()
	return out, err
}
