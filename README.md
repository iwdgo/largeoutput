[![GoDoc](https://godoc.org/github.com/iWdGo/largeoutput?status.svg)](https://godoc.org/github.com/iWdGo/largeoutput)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/largeoutput)](https://goreportcard.com/report/github.com/iwdgo/largeoutput)
[![codecov](https://codecov.io/gh/iWdGo/largeoutput/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/largeoutput)

[![Build Status](https://travis-ci.com/iWdGo/largeoutput.svg?branch=master)](https://travis-ci.com/iWdGo/largeoutput)
[![Build Status](https://api.cirrus-ci.com/github/iWdGo/largeoutput.svg)](https://cirrus-ci.com/github/iWdGo/largeoutput)
[![Build status](https://ci.appveyor.com/api/projects/status/eimlas99romrrro0?svg=true)](https://ci.appveyor.com/project/iWdGo/largeoutput)
![Build status](https://github.com/iwdgo/largeoutput/workflows/Go/badge.svg)

# Some ordinary exercises

- Editing a large file.
- Outputing numbers following some rules.
- Converting a set of strings which contains numbers.

## Handling output when using a file.

An ordinary exercise may have a large output for which the easy `// Output:`
becomes unreadable.

In that case a convenient structure is an `Example` for a reduced set of data and
a test and its benchmark for larger outputs.

If you need to output to a file, you might have to update your code.
Piping `Stdout` to a file is easier. A helper to compare the file to a referenced file
returns eventually an error usable by the `testing` package.

The original code from `testing/example.go` is using unexported routines.
It has been adapted to ordinary exercises.

## Data files

The `/output` subdirectory avoids to have the data files mixed with source code.
The directory is not created but its existence is checked.
If the working directory (not the temp, nor the executing) is unavailable,
tests will panic.

It is not needed to use the related test helper because a benchmark runs the related test
first. The test sets the working directory correctly.
