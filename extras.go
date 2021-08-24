package owl

import (
	"os"
	"reflect"
	"strings"

	"github.com/stretchr/testify/require"
)

// ShellAliases registers a build-shell-aliases subcommand that
// auto-generates shell aliases for all active subcommands.
type ShellAliases struct {
	ShellAliases *shellAliasesCmd `arg:"subcommand:build-shell-aliases" help:"generate shell aliases for all subcommands"`
}

const aliasesPreamble = `#
# Add the following three lines to your ~/.zshrc or ~/.bashrc file:
# if command -v %[1]s > /dev/null; then
#     source <(%[1]s build-shell-aliases)
# fi
#
`

type shellAliasesCmd struct{}

func (c *shellAliasesCmd) Run(o Owl) error {
	binary, err := os.Executable()
	require.NoError(o, err, "cannot find current binary path")
	o.Printf(aliasesPreamble, binary)

	commands := reflect.TypeOf(o).Elem()
	for i := 0; i < commands.NumField(); i++ {
		argTags := commands.Field(i).Tag.Get("arg")
		for _, tag := range strings.Split(argTags, ",") {
			if strings.HasPrefix(tag, "subcommand:") {
				name := strings.TrimPrefix(tag, "subcommand:")
				o.Printf("alias %[1]s='%[2]s %[1]s'\n", name, binary)
			}
		}
	}
	return nil
}
