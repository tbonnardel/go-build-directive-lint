# `go-build-directive-lint`

A simple go lint to reject files with the old build directive syntax, using `go/analysis`.

## Motivation

Since Go 1.17, [build constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints) syntax changes to the benefit of resilience, robustness and readability.

The old way to define build constraints was `// +build` is now deprecated. Use the new syntax `// go:build` instead.

## Usage

```
gobuilddirectivelint: Check that build directives follow the new syntax introduced in Go 1.17.

Usage: gobuilddirectivelint [-flag] [package]


Flags:
  -V    print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -source
        no effect (deprecated)
  -tags string
        no effect (deprecated)
  -trace string
        write trace log to this file
  -v    no effect (deprecated)
```
