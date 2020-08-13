// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package trie implements a simple prefix tree.  This is designed to be used
// for text command completion and is reasonably efficient in that application.
package trie

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// Node represents an entire prefix tree or a node within it.
type Node struct {
	// char is the node character.  For the root node this will be null
	// (\x00).
	char rune

	// val is the value of the node if it's terminal, otherwise it will be
	// nil.
	val interface{}

	// children is a map of child nodes keyed by the char.
	children map[rune]*Node
}

// Add adds a key and value to the tree.  The value must be non-nil for it to be
// considered terminal.
func (n *Node) Add(key string, val interface{}) {
	// Walk the nodes for each key character and add any missing nodes along
	// the way.
	cur := n
	for _, c := range key {
		child, exists := cur.children[c]
		if !exists {
			child = &Node{char: c}
			if cur.children == nil {
				cur.children = make(map[rune]*Node)
			}
			cur.children[c] = child
		}
		cur = child
	}

	// The current node is the end of the key so set the value.
	cur.val = val
}

// Children returns the immediate child nodes.
func (n *Node) Children(sorted bool) []*Node {
	children := make([]*Node, 0, len(n.children))
	for _, child := range n.children {
		children = append(children, child)
	}

	if sorted {
		sort.Slice(children, func(i, j int) bool {
			return children[i].char < children[j].char
		})
	}

	return children
}

// dump walks the node and all its children and pretty-prints to the specified
// buffer.
func (n Node) dump(buf *bytes.Buffer, charTrail []rune, branches []bool, moreTwigs bool) {
	// Write branch indents.
	branches = append(branches, moreTwigs)
	for i := 0; i < len(branches)-1; i++ {
		if branches[i] {
			buf.WriteString(" |  ")
		} else {
			buf.WriteString("    ")
		}
	}

	// Write the node and include value if it's terminal.
	buf.WriteString(" `--(")
	if n.char == '\x00' { // Root
		buf.WriteRune('*')
	} else {
		buf.WriteRune(n.char)
		charTrail = append(charTrail, n.char)
	}
	buf.WriteRune(')')
	if n.val != nil {
		fmt.Fprintf(buf, " => \"%s\": %s", string(charTrail), n.val)
	}

	buf.WriteString("\n")

	// Walk the child nodes.
	children := n.Children(true)
	for i, child := range children {
		// If this is the last child node then indicate that there are
		// no more twigs.
		child.dump(buf, charTrail, branches, i+1 < len(children))
	}
}

// Find returns the node that completes the key as much as possible while
// remaining unique.
func (n *Node) Find(key string) (match string, cur *Node) {
	cur = n

	// Split the key into words and complete each one.
	for _, w := range strings.Split(key, " ") {
		cur = cur.Get(w)
		if cur == nil {
			return
		}

		var m string
		cur.walk(w, true, true, func(key string, n *Node) bool {
			m = key
			cur = n

			return n.char != ' '
		})

		match = match + m
	}

	// Complete anything after the last word.
	cur.walk(match, true, false, func(key string, n *Node) bool {
		match = key
		cur = n

		return true
	})

	return
}

// Get returns the node of the given key or nil if it's not found.
func (n *Node) Get(key string) *Node {
	// Walk the nodes for each key character until we reach a missing node
	// or complete the key.
	cur := n
	for _, c := range key {
		child, exists := cur.children[c]
		if !exists {
			return nil
		}
		cur = child
	}

	return cur
}

// Match returns all possible completions for the given key.
func (n Node) Match(key string) <-chan string {
	matches := make(chan string)

	go func() {
		defer close(matches)

		cur := n.Get(key)
		if cur == nil {
			return
		}

		cur.walk(key, false, false, func(key string, n *Node) bool {
			matches <- key

			return true
		})
	}()

	return matches
}

// String retruns a pretty-print string of the node and all its children.
func (n Node) String() string {
	var buf bytes.Buffer
	n.dump(&buf, []rune{}, []bool{}, false)
	return buf.String()
}

// Val returns the node value.
func (n Node) Val() interface{} {
	return n.val
}

type walkFunc func(string, *Node) bool

func (n *Node) walk(key string, onlyUniq bool, fAll bool, f walkFunc) {
	// If we are only looking for unique matches and there is more than one
	// child then stop walking.
	if onlyUniq && len(n.children) > 1 {
		f(key, n)
		return
	}

	// Call user function if the node is terminal or we were asked to for
	// all walks.
	if fAll || n.val != nil {
		if !f(key, n) {
			return
		}
	}

	// Walk the child nodes.
	for _, child := range n.Children(true) {
		child.walk(key+string(child.char), onlyUniq, fAll, f)
	}
}
