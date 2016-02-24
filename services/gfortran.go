package services

import (
	"fmt"
	"os/exec"
)

func fcompile(source, dialect string) (out []byte, err error) {
	var standard string
	switch dialect {
	case "legacy":
		standard = "-std=legacy"
	case "f95":
		standard = "-std=f95"
	case "f2003":
		standard = "-std=f2003"
	case "f2008":
		standard = "-std=f2003"
	case "c11":
		standard = "-std=c11"
	}
	cflags := fmt.Sprintf("%s -pedantic -Werror -Wall -mtune=i386 -o %s.out",
		standard, source)
	out, err = exec.Command(cc, cflags).CombinedOutput()
	return
}
