# textcmd

```go
import "github.com/ebarkie/textcmd"
```

Package textcmd implements a simple text command shell and executor.

## Usage

```go
var (
	ErrCmdNotFound = errors.New("command not found")
	ErrCmdQuit     = errors.New("quit command")
)
```
Errors.

#### type Env

```go
type Env struct {
	io.ReadWriter
}
```

Env is the command environment passed to a function.

#### func (Env) Arg

```go
func (e Env) Arg(i int) (s string)
```
Arg returns the argument at index i. 0 is the command , 1 is the first argument,
2 is the second, etc. If an argument does not exist an empty string is returned.

#### type Shell

```go
type Shell struct {
}
```

Shell is a text command shell for which commands can be registered and executed.

#### func (Shell) Exec

```go
func (sh Shell) Exec(rw io.ReadWriter, s string) error
```
Exec attempts to execute the passed string as a command.

#### func (*Shell) Register

```go
func (sh *Shell) Register(f cmdFunc, cmd ...string)
```
Register adds a command to the text command shell. It takes a a command function
and command execution strings.
