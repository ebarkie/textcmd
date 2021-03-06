// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package textcmd

import (
	"net"
	"strings"
)

// Env is the command environment passed to a function.
type Env struct {
	net.Conn
	args []string
}

// Arg returns the argument at index i.  0 is the command , 1
// is the first argument, 2 is the second, etc.  If an argument
// does not exist an empty string is returned.
func (e Env) Arg(i int) (s string) {
	if len(e.args) > i {
		s = strings.ToLower(e.args[i])
	}

	return
}
