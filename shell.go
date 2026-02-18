// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package textcmd implements a simple text command shell and
// executor.
package textcmd

import (
	"context"
	"errors"
	"io"
	"iter"
	"strings"

	"github.com/ebarkie/textcmd/internal/trie"
)

// Errors.
var (
	ErrCmdNotFound = errors.New("command not found")
	ErrCmdQuit     = errors.New("quit command")
)

// CmdFunc is the function signature for command handlers.
type CmdFunc func(context.Context, io.ReadWriter, ...string) error

// Shell is a text command shell for which commands can be
// registered and executed.
type Shell struct {
	cmds trie.Node
}

// Exec attempts to execute the passed string as a command.
func (sh Shell) Exec(ctx context.Context, rw io.ReadWriter, s string) error {
	tokens := strings.Fields(s)
	for i := range tokens {
		cmd := strings.Join(tokens[:i+1], " ")

		if _, cur := sh.cmds.Find(cmd, ' '); cur != nil && cur.Val != nil {
			return cur.Val.(CmdFunc)(ctx, rw, tokens[i+1:]...)
		}
	}

	return ErrCmdNotFound
}

// Complete returns the input expanded as far as possible and all possible full
// command strings.
func (sh Shell) Complete(s string) (completion string, matches iter.Seq[string]) {
	completion, _ = sh.cmds.Find(s, ' ')
	if completion == "" {
		completion = s
	}

	matches = sh.cmds.Match(completion)

	return
}

// Register adds a command to the text command shell.  It takes a
// command function and command execution strings.
func (sh *Shell) Register(f CmdFunc, cmd ...string) {
	for _, c := range cmd {
		sh.cmds.Add(c, f)
	}
}
