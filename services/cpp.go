package services

import (
	"fmt"
	"os/exec"
)

func cppcompile(source, dialect string) (out []byte, err error) {
	var standard string
	switch dialect {
	case "ansi":
		standard = "-ansi"
	case "c++98":
		standard = "-std=c++98"
	case "c++03":
		standard = "-std=c++03"
	case "c++11":
		standard = "-std=c++11"
	case "c++14":
		standard = "-std=c++14"
	}
	cflags := fmt.Sprintf("%s -pedantic -Werror -Wall -mtune=i386 -o %s.out",
		standard, source)
	out, err = exec.Command(cpp, cflags).CombinedOutput()
	return
}
