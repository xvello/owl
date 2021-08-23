package must

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xvello/owl/mocks"
)

func TestExec_Simple(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.On("ExecCommand", "echo", "one", "two").
		Return(exec.Command("echo", "one", "two")).Once()
	assert.Equal(t, "one two", Exec(mowl, "echo", "one", "two"))
}

func TestExec_SplitArgs(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.On("ExecCommand", "echo", "a", "b", "c").
		Return(exec.Command("echo", "a", "b", "c")).Once()
	assert.Equal(t, "a b c", Exec(mowl, "echo a b c"))
}

func TestExec_Err(t *testing.T) {
	mowl := new(mocks.Owl)
	mowl.On("ExecCommand", "false").Return(exec.Command("false")).Once()
	mowl.ExpectRequireFailure(t, "exit status 1")
	assert.Panics(t, func() { Exec(mowl, "false") })
}
