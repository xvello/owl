package must

import (
	"strings"

	"github.com/stretchr/testify/require"
	"github.com/xvello/owl"
)

// Exec wraps execution of an external command. If no arguments are given,
// they are extracted from the command, by splitting it on spaces.
// If the command fails, its output is printed and the command stops.
// If the command succeeds, its output is returned as a string.
func Exec(o owl.Owl, command string, args ...string) string {
	if len(args) == 0 {
		parts := strings.Split(command, " ")
		command = parts[0]
		args = parts[1:]
	}
	out, err := o.ExecCommand(command, args...).CombinedOutput()
	require.NoError(o, err, string(out))
	return strings.TrimSpace(string(out))
}
