# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Owl provides a `ExecCommand` function, wrapping `exec.Command` to ease unit testing
  of commands calling external programs. `must.Exec` uses this wrapper.

### Changed

- [BREAKING] The `build-bash-aliases` extra command now works with zsh and is renamed
  to `build-shell-aliases`. It is provided via the `ShellAliases` embed, sunsetting
  the too broad `Extras` embed.

### 

## [0.1.0] - 2021-08-23

Initial release