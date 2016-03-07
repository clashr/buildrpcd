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
