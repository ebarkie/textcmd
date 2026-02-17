// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package textcmd

import "testing"

func TestHistoryPrevNext(t *testing.T) {
	h := NewHistory(0)
	h.Add("first")
	h.Add("second")
	h.Add("third")

	// Walk backward.
	if got := h.Prev(); got != "third" {
		t.Errorf("Prev() = %q, want %q", got, "third")
	}
	if got := h.Prev(); got != "second" {
		t.Errorf("Prev() = %q, want %q", got, "second")
	}
	if got := h.Prev(); got != "first" {
		t.Errorf("Prev() = %q, want %q", got, "first")
	}

	// Walk forward.
	if got := h.Next(); got != "second" {
		t.Errorf("Next() = %q, want %q", got, "second")
	}
	if got := h.Next(); got != "third" {
		t.Errorf("Next() = %q, want %q", got, "third")
	}

	// Past the end returns "".
	if got := h.Next(); got != "" {
		t.Errorf("Next() past end = %q, want %q", got, "")
	}
}

func TestHistoryBounds(t *testing.T) {
	h := NewHistory(0)

	// Empty history.
	if got := h.Prev(); got != "" {
		t.Errorf("Prev() on empty = %q, want %q", got, "")
	}
	if got := h.Next(); got != "" {
		t.Errorf("Next() on empty = %q, want %q", got, "")
	}

	h.Add("only")

	// Stays at first entry.
	if got := h.Prev(); got != "only" {
		t.Errorf("Prev() = %q, want %q", got, "only")
	}
	if got := h.Prev(); got != "only" {
		t.Errorf("Prev() at beginning = %q, want %q", got, "only")
	}
}

func TestHistoryMaxEviction(t *testing.T) {
	h := NewHistory(3)
	h.Add("a")
	h.Add("b")
	h.Add("c")
	h.Add("d")

	// "a" should be evicted.
	var got []string
	for {
		s := h.Prev()
		if len(got) > 0 && s == got[len(got)-1] {
			break
		}
		got = append(got, s)
	}

	want := []string{"d", "c", "b"}
	if len(got) != len(want) {
		t.Fatalf("entries = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("entry[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestHistoryDuplicateSkip(t *testing.T) {
	h := NewHistory(0)
	h.Add("cmd")
	h.Add("cmd")
	h.Add("cmd")

	if got := h.Prev(); got != "cmd" {
		t.Errorf("Prev() = %q, want %q", got, "cmd")
	}
	// Should stay at same entry since there's only one.
	if got := h.Prev(); got != "cmd" {
		t.Errorf("Prev() at beginning = %q, want %q", got, "cmd")
	}
}

func TestHistoryEmptyWhitespaceSkip(t *testing.T) {
	h := NewHistory(0)
	h.Add("")
	h.Add("   ")
	h.Add("\t\n")

	if got := h.Prev(); got != "" {
		t.Errorf("Prev() = %q, want %q", got, "")
	}
}

func TestHistoryTrimWhitespace(t *testing.T) {
	h := NewHistory(0)
	h.Add("  hello  ")

	if got := h.Prev(); got != "hello" {
		t.Errorf("Prev() = %q, want %q", got, "hello")
	}
}

func TestHistoryReset(t *testing.T) {
	h := NewHistory(0)
	h.Add("first")
	h.Add("second")

	h.Prev() // "second"
	h.Prev() // "first"
	h.Reset()

	// After reset, Prev returns last entry again.
	if got := h.Prev(); got != "second" {
		t.Errorf("Prev() after Reset = %q, want %q", got, "second")
	}
}
