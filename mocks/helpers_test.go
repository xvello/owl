package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOnExecCommand_Run(t *testing.T) {
	mowl := new(Owl)
	mowl.OnExecCommand("echo", "true").Run().Once()
	cmd := mowl.ExecCommand("echo", "true")
	assert.Contains(t, cmd.Path, "echo")
	assert.Equal(t, []string{"echo", "true"}, cmd.Args)
	mock.AssertExpectationsForObjects(t, mowl)
}

func TestOnExecCommand_ExecBash(t *testing.T) {
	mowl := new(Owl)
	mowl.OnExecCommand("echo", "true").ExecBash("false").Once()
	cmd := mowl.ExecCommand("echo", "true")
	assert.Contains(t, cmd.Path, "bash")
	assert.Equal(t, []string{"bash", "-c", "false"}, cmd.Args)
	mock.AssertExpectationsForObjects(t, mowl)
}
