// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package textcmd

import "strings"

// History tracks command history with readline-style cursor
// navigation.  It is per-session state and is not tied to a
// Shell.
type History struct {
	entries []string
	max     int
	cursor  int
}

// NewHistory creates a new History.  max limits the number of
// entries kept; max <= 0 means unbounded.
func NewHistory(max int) *History {
	return &History{max: max}
}

// Add appends a command to the history and resets the cursor
// to the end.  Empty/whitespace-only strings and consecutive
// duplicates are skipped.
func (h *History) Add(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	if len(h.entries) > 0 && h.entries[len(h.entries)-1] == s {
		h.cursor = len(h.entries)
		return
	}

	h.entries = append(h.entries, s)
	if h.max > 0 && len(h.entries) > h.max {
		h.entries = h.entries[len(h.entries)-h.max:]
	}
	h.cursor = len(h.entries)
}

// Prev moves the cursor back and returns the entry.  Returns
// "" if history is empty.  Stays at the first entry if already
// at the beginning.
func (h *History) Prev() string {
	if len(h.entries) == 0 {
		return ""
	}

	if h.cursor > 0 {
		h.cursor--
	}

	return h.entries[h.cursor]
}

// Next moves the cursor forward and returns the entry.  Returns
// "" when moving past the last entry.
func (h *History) Next() string {
	if len(h.entries) == 0 || h.cursor >= len(h.entries) {
		return ""
	}

	h.cursor++
	if h.cursor >= len(h.entries) {
		return ""
	}

	return h.entries[h.cursor]
}

// Reset moves the cursor to the end (past the last entry).
func (h *History) Reset() {
	h.cursor = len(h.entries)
}
