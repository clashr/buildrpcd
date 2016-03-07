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
	out, err := exec.Command(cc, std, "-pedantic", "-Werror", "-static",
		"-Wall", "-pipe", "-fPIC", "-o", fmt.Sprintf("%s.out", source),
		source).CombinedOutput()
	return out, err
}
