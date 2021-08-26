![logo](.github/logo.svg?raw=true)

# Owl - Only Write (your) Logic

[![CI](https://github.com/xvello/owl/actions/workflows/go.yml/badge.svg)](https://github.com/xvello/owl/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/xvello/owl)](https://goreportcard.com/report/github.com/xvello/owl)
[![Coverage Status](https://coveralls.io/repos/github/xvello/owl/badge.svg?branch=main)](https://coveralls.io/github/xvello/owl?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/xvello/owl.svg)](https://pkg.go.dev/github.com/xvello/owl)

Owl is a micro-framework that builds on top of thethe excellent [go-arg](https://github.com/alexflint/go-arg) and
[testify](https://github.com/stretchr/testify) libraries to allow developers to write their CLI tools in Golang,
instead of maintaining brittle Bash / Python scripts.

All engineering teams will accumulate an (official or unofficial) collection of scripts to automate dev and ops tasks.
Eventually, an outage will happen because of bad error handling, or a subtle bug that a static typing system would
have caught. The goal of this project is to allow teams to maintain "scripts" as a Go project:

- borrowing from Busybox, all commands are compiled in a single binary for easy distribution
- input and current state can be checked with [testify/require](https://pkg.go.dev/github.com/stretchr/testify/require)
before running commands. No need to google for the bash test syntax anymore!
- common actions are more concise thanks to helpers in the [must package](https://pkg.go.dev/github.com/xvello/owl/must)
- commands can be extensively unit-tested, with a [mocked owl](https://pkg.go.dev/github.com/xvello/owl/mocks)
allowing to intercept and mock calls to external commands
- you can leverage all the power of the standard library, and any Go library you already work with

## Minimal example

```go
// rootCommand holds the list of active commands, and embeds owl.Base for common helpers.
type rootCommand struct {
    owl.Base
    Lower   *lowerCmd `arg:"subcommand:lower" help:"return the text in lowercase"`
}

// lowerCmd holds the arguments for this command.
type lowerCmd struct {
    Text string `arg:"positional" help:"text to lowercase"`
}

// Run is the entrypoint for your command, argument values are set in the struct fields.
func (c *lowerCmd) Run(o owl.Owl) {
    out := strings.ToLower(c.Text)
    require.NotEqual(o, "viper", out, "snakes not allowed here")
    o.Println(out)
}

// main just calls owl.RunOwl to parse arguments and start the selected command.
func main() {
    owl.RunOwl(new(rootCommand))
}
```
