package owl

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellAliases(t *testing.T) {
	executable, _ := os.Executable()
	expected := fmt.Sprintf(aliasesPreamble+"alias simple='%[1]s simple'\nalias another='%[1]s another'\n", executable)
	var stdout strings.Builder
	var stderr strings.Builder
	c := &struct {
		Base
		ShellAliases
		Simple   *simpleSub   `arg:"subcommand:simple"`
		Advanced *fallibleSub `arg:"subcommand:another"`
	}{
		Base: Base{
			stdout: &stdout,
			stderr: &stderr,
		},
	}

	os.Args = []string{"owl", "build-shell-aliases"}
	RunOwl(c)
	assert.Empty(t, stderr.String())
	assert.Equal(t, expected, stdout.String())
}
