package mocks

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ExpectAssertFailure configures the mock to expect a testify/assert failure containing the given string.
func (_m *Owl) ExpectAssertFailure(t *testing.T, contains string) {
	_m.On("Errorf", mock.MatchedBy(func(format string) bool {
		return format == "\n%s"
	}), mock.MatchedBy(func(arg string) bool {
		return assert.Contains(t, arg, contains)
	})).Return()
}

// ExpectRequireFailure configures the mock to expect a testify/require failure containing the given string.
func (_m *Owl) ExpectRequireFailure(t *testing.T, contains string) {
	_m.ExpectAssertFailure(t, contains)
	_m.On("FailNow").Panic("FailNow called by subcommand")
}

// ExecCommandCall is returned by OnExecCommand, see its methods to set the return value.
type ExecCommandCall struct {
	call *mock.Call
	name string
	arg  []string
}

// OnExecCommand configures the mock to expect a call to ExecCommand and returns a
// ExecCommandCall that can be used to define what to return.
func (_m *Owl) OnExecCommand(name string, arg ...string) *ExecCommandCall {
	var parts []interface{}
	parts = append(parts, name)
	for _, a := range arg {
		parts = append(parts, a)
	}
	return &ExecCommandCall{
		call: _m.On("ExecCommand", parts...),
		name: name,
		arg:  arg,
	}
}

// Run configures the mock to actually run the specified command. Do this if the invoked
// command is a build dependency of your project and you control the side effets. Otherwise,
// see ExecBash instead.
func (c *ExecCommandCall) Run() *mock.Call {
	return c.call.Return(exec.Command(c.name, c.arg...))
}

// ExecBash configures the mock to run a bash one-liner instead of the specified command.
// Do this to avoid failing if the command is not available, or you want to be resilient to
// its behaviour changing.
func (c *ExecCommandCall) ExecBash(shell string) *mock.Call {
	return c.call.Return(exec.Command("bash", "-c", shell))
}
