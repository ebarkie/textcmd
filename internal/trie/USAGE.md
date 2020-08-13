# trie

```go
import "github.com/ebarkie/textcmd/internal/trie"
```

Package trie implements a simple prefix tree. This is designed to be used for
text command completion and is reasonably efficient in that application.

## Usage

#### type Node

```go
type Node struct {

	// Val is the value of the node if it's terminal, otherwise it will be
	// nil.
	Val interface{}
}
```

Node represents an entire prefix tree or a node within it.

#### func (*Node) Add

```go
func (n *Node) Add(key string, val interface{})
```
Add adds a key and value to the tree. Terminal values must be non-nil.

#### func (*Node) Children

```go
func (n *Node) Children(sorted bool) []*Node
```
Children returns the immediate child nodes optionally sorted in alphabetical
order.

#### func (*Node) Find

```go
func (n *Node) Find(key string, sep rune) (match string, cur *Node)
```
Find returns the node that completes the key as much as possible while remaining
unique. The key is split by sep and each part is completed individually.

#### func (*Node) Get

```go
func (n *Node) Get(key string) *Node
```
Get returns the node of the given key or nil if it's not found.

#### func (Node) Match

```go
func (n Node) Match(key string) <-chan string
```
Match returns all possible completions for the given key.

#### func (Node) String

```go
func (n Node) String() string
```
String retruns a pretty-printed string of the node and all its children.
