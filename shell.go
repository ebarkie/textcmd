// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package textcmd implements a simple text command shell and
// executor.
package textcmd

import (
	"errors"
	"net"
	"strings"

	"github.com/ebarkie/textcmd/internal/trie"
)

// Errors.
var (
	ErrCmdNotFound = errors.New("command not found")
	ErrCmdQuit     = errors.New("quit command")
)

// cmdFunc is the function that will be called when a command is
// executed.
type cmdFunc func(Env) error

// Shell is a text command shell for which commands can be
// registered and executed.
type Shell struct {
	cmds trie.Node
}

// Exec attempts to execute the passed string as a command.
func (sh Shell) Exec(conn net.Conn, s string) error {
	tokens := strings.Split(s, " ")
	for i := 0; i < len(tokens); i++ {
		cmd := strings.Join(tokens[:i+1], " ")

		if match, cur := sh.cmds.Find(cmd); cur != nil && cur.Val() != nil {
			return cur.Val().(cmdFunc)(Env{
				Conn: conn,
				args: append([]string{match}, tokens[i+1:]...)})
		}
	}

	return ErrCmdNotFound
}

// Register adds a command to the text command shell.  It takes a
// a command function and command execution strings.
func (sh *Shell) Register(f cmdFunc, cmd ...string) {
	for _, c := range cmd {
		sh.cmds.Add(c, f)
	}
}
