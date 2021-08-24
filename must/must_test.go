package must

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xvello/owl/mocks"
)

func TestExec_Simple(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.OnExecCommand("echo", "one", "two").Run().Once()
	assert.Equal(t, "one two", Exec(mowl, "echo", "one", "two"))
}

func TestExec_SplitArgs(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.OnExecCommand("echo", "a", "b", "c").Run().Once()
	assert.Equal(t, "a b c", Exec(mowl, "echo a b c"))
}

func TestExec_Err(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.OnExecCommand("false").Run().Once()
	mowl.ExpectRequireFailure(t, "exit status 1")
	assert.Panics(t, func() { Exec(mowl, "false") })
}

func TestExec_ErrFromBash(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.OnExecCommand("true").ExecBash("echo mocked; false").Once()
	mowl.ExpectRequireFailure(t, "mocked")
	assert.Panics(t, func() { Exec(mowl, "true") })
}
