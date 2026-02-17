// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package textcmd

import (
	"slices"
	"testing"
)

func testShell() Shell {
	var sh Shell
	for _, cmd := range []string{
		"?",
		"archive",
		"conditions",
		"date",
		"exit",
		"help",
		"health",
		"lamps off",
		"lamps on",
		"logout",
		"loop",
		"quit",
		"time",
		"trend",
		"uname",
		"uptime",
		"version",
		"watch conditions",
		"watch log debug",
		"watch log trace",
		"watch loops",
		"whoami",
	} {
		sh.Register(func(Env) error { return nil }, cmd)
	}

	return sh
}

func TestComplete(t *testing.T) {
	sh := testShell()

	for _, test := range []struct {
		name       string
		input      string
		completion string
		matches    []string
	}{
		{
			name:       "unique",
			input:      "d",
			completion: "date",
			matches:    []string{"date"},
		},
		{
			name:       "ambiguous",
			input:      "he",
			completion: "he",
			matches:    []string{"health", "help"},
		},
		{
			name:       "multi-word",
			input:      "wa l",
			completion: "watch lo",
			matches:    []string{"watch log debug", "watch log trace", "watch loops"},
		},
		{
			name:       "no match",
			input:      "xyz",
			completion: "xyz",
			matches:    []string{},
		},
		{
			name:       "empty input",
			input:      "",
			completion: "",
			matches: []string{
				"?",
				"archive",
				"conditions",
				"date",
				"exit",
				"health",
				"help",
				"lamps off",
				"lamps on",
				"logout",
				"loop",
				"quit",
				"time",
				"trend",
				"uname",
				"uptime",
				"version",
				"watch conditions",
				"watch log debug",
				"watch log trace",
				"watch loops",
				"whoami",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			completion, matches := sh.Complete(test.input)

			if completion != test.completion {
				t.Errorf("completion = %q, want %q", completion, test.completion)
			}

			got := slices.Collect(matches)
			if len(got) == 0 && len(test.matches) == 0 {
				return
			}
			if !slices.Equal(got, test.matches) {
				t.Errorf("matches = %v, want %v", got, test.matches)
			}
		})
	}
}
