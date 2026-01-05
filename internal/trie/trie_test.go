// Copyright (c) 2020 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package trie

import (
	"fmt"
	"testing"
)

var testKeys = [...]string{
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
}

func testVal(key string) any {
	return "<value for \"" + key + "\">"
}

func testTree() *Node {
	n := &Node{}
	for _, k := range testKeys {
		n.Add(k, testVal(k))
	}

	return n
}

func BenchmarkNode_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testTree()
	}
}

func BenchmarkNode_String(b *testing.B) {
	n := testTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = n.String()
	}
}

func ExampleNode_String() {
	n := testTree()
	fmt.Println(n.String())
	// Output:
	// `--(*)
	//      `--(?) => "?": <value for "?">
	//      `--(a)
	//      |   `--(r)
	//      |       `--(c)
	//      |           `--(h)
	//      |               `--(i)
	//      |                   `--(v)
	//      |                       `--(e) => "archive": <value for "archive">
	//      `--(c)
	//      |   `--(o)
	//      |       `--(n)
	//      |           `--(d)
	//      |               `--(i)
	//      |                   `--(t)
	//      |                       `--(i)
	//      |                           `--(o)
	//      |                               `--(n)
	//      |                                   `--(s) => "conditions": <value for "conditions">
	//      `--(d)
	//      |   `--(a)
	//      |       `--(t)
	//      |           `--(e) => "date": <value for "date">
	//      `--(e)
	//      |   `--(x)
	//      |       `--(i)
	//      |           `--(t) => "exit": <value for "exit">
	//      `--(h)
	//      |   `--(e)
	//      |       `--(a)
	//      |       |   `--(l)
	//      |       |       `--(t)
	//      |       |           `--(h) => "health": <value for "health">
	//      |       `--(l)
	//      |           `--(p) => "help": <value for "help">
	//      `--(l)
	//      |   `--(a)
	//      |   |   `--(m)
	//      |   |       `--(p)
	//      |   |           `--(s)
	//      |   |               `--( )
	//      |   |                   `--(o)
	//      |   |                       `--(f)
	//      |   |                       |   `--(f) => "lamps off": <value for "lamps off">
	//      |   |                       `--(n) => "lamps on": <value for "lamps on">
	//      |   `--(o)
	//      |       `--(g)
	//      |       |   `--(o)
	//      |       |       `--(u)
	//      |       |           `--(t) => "logout": <value for "logout">
	//      |       `--(o)
	//      |           `--(p) => "loop": <value for "loop">
	//      `--(q)
	//      |   `--(u)
	//      |       `--(i)
	//      |           `--(t) => "quit": <value for "quit">
	//      `--(t)
	//      |   `--(i)
	//      |   |   `--(m)
	//      |   |       `--(e) => "time": <value for "time">
	//      |   `--(r)
	//      |       `--(e)
	//      |           `--(n)
	//      |               `--(d) => "trend": <value for "trend">
	//      `--(u)
	//      |   `--(n)
	//      |   |   `--(a)
	//      |   |       `--(m)
	//      |   |           `--(e) => "uname": <value for "uname">
	//      |   `--(p)
	//      |       `--(t)
	//      |           `--(i)
	//      |               `--(m)
	//      |                   `--(e) => "uptime": <value for "uptime">
	//      `--(v)
	//      |   `--(e)
	//      |       `--(r)
	//      |           `--(s)
	//      |               `--(i)
	//      |                   `--(o)
	//      |                       `--(n) => "version": <value for "version">
	//      `--(w)
	//          `--(a)
	//          |   `--(t)
	//          |       `--(c)
	//          |           `--(h)
	//          |               `--( )
	//          |                   `--(c)
	//          |                   |   `--(o)
	//          |                   |       `--(n)
	//          |                   |           `--(d)
	//          |                   |               `--(i)
	//          |                   |                   `--(t)
	//          |                   |                       `--(i)
	//          |                   |                           `--(o)
	//          |                   |                               `--(n)
	//          |                   |                                   `--(s) => "watch conditions": <value for "watch conditions">
	//          |                   `--(l)
	//          |                       `--(o)
	//          |                           `--(g)
	//          |                           |   `--( )
	//          |                           |       `--(d)
	//          |                           |       |   `--(e)
	//          |                           |       |       `--(b)
	//          |                           |       |           `--(u)
	//          |                           |       |               `--(g) => "watch log debug": <value for "watch log debug">
	//          |                           |       `--(t)
	//          |                           |           `--(r)
	//          |                           |               `--(a)
	//          |                           |                   `--(c)
	//          |                           |                       `--(e) => "watch log trace": <value for "watch log trace">
	//          |                           `--(o)
	//          |                               `--(p)
	//          |                                   `--(s) => "watch loops": <value for "watch loops">
	//          `--(h)
	//              `--(o)
	//                  `--(a)
	//                      `--(m)
	//                          `--(i) => "whoami": <value for "whoami">

}

func TestNode_Add(t *testing.T) {
	testTree()
}

func TestNode_Find(t *testing.T) {
	n := testTree()
	for _, test := range []struct {
		key   string
		match string
		val   any
	}{
		{"d", "date", testVal("date")},
		{"l", "l", nil},
		{"l off", "l", nil},
		{"la", "lamps o", nil},
		{"la on", "lamps on", testVal("lamps on")},
		{"lamps off", "lamps off", testVal("lamps off")},
		{"log", "logout", testVal("logout")},
		{"up", "uptime", testVal("uptime")},
		{"version", "version", testVal("version")},
		{"w c", "w", nil},
		{"wa c", "watch conditions", testVal("watch conditions")},
		{"wa z", "watch ", nil},
		{"wa l", "watch lo", nil},
		{"wa lo", "watch lo", nil},
		{"wa log", "watch log ", nil},
		{"wa log d", "watch log debug", testVal("watch log debug")},
		{"wa loo", "watch loops", testVal("watch loops")},
		{"wh", "whoami", testVal("whoami")},
	} {
		t.Run(test.key, func(t *testing.T) {
			match, cur := n.Find(test.key, ' ')
			if match != test.match {
				t.Errorf("%q match is %q but expected %q", test.key, match, test.match)
			}

			if cur == nil && test.val != nil {
				t.Errorf("%q value is <nil> but expected %v", test.key, test.val)
			} else if cur != nil && cur.Val != nil && cur.Val != test.val {
				t.Errorf("%q value is %v but expected %v", test.key, cur.Val, test.val)
			}
		})
	}
}

func TestNode_Get(t *testing.T) {
	n := testTree()
	for _, test := range []struct {
		key string
		val any
	}{
		{"", nil},
		{"?", testVal("?")},
		{"a", nil},
		{"archiv", nil},
		{"archive", testVal("archive")},
		{"archive ", nil},
		{"Archive", nil},
		{"foo", nil},
	} {
		t.Run(test.key, func(t *testing.T) {
			cur := n.Get(test.key)
			if cur == nil && test.val != nil {
				t.Errorf("%q value is <nil> but expected %v", test.key, test.val)
			} else if cur != nil && cur.Val != nil && cur.Val != test.val {
				t.Errorf("%q value is %v but expected %v", test.key, cur.Val, test.val)
			}
		})
	}
}

func TestNode_Match(t *testing.T) {
	n := testTree()
	for _, test := range []struct {
		key  string
		keys []string
	}{
		{"lo", []string{"logout", "loop"}},
		{"wa", []string{
			"watch conditions",
			"watch log debug",
			"watch log trace",
			"watch loops",
		}},
		{"z", []string{}},
	} {
		t.Run(test.key, func(t *testing.T) {
			i := 0
			for k := range n.Match(test.key) {
				if i >= len(test.keys) {
					t.Errorf("%q: unexpected index %d %q", test.key, i, k)
				} else if k != test.keys[i] {
					t.Errorf("%q: mismatched index %d %q != %q", test.key, i, k, test.keys[i])
				}

				i++
			}
			if i < len(test.keys) {
				t.Errorf("%q: missing %q", test.key, test.keys[i-1:])
			}
		})
	}
}
