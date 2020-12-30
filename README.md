[![Go Reference](https://pkg.go.dev/badge/iwdgo/largeoutput.svg)](https://pkg.go.dev/iwdgo/largeoutput)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/largeoutput)](https://goreportcard.com/report/github.com/iwdgo/largeoutput)
[![codecov](https://codecov.io/gh/iWdGo/largeoutput/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/largeoutput)

[![Build Status](https://travis-ci.com/iWdGo/largeoutput.svg?branch=master)](https://travis-ci.com/iWdGo/largeoutput)
[![Build Status](https://api.cirrus-ci.com/github/iWdGo/largeoutput.svg)](https://cirrus-ci.com/github/iWdGo/largeoutput)
[![Build status](https://ci.appveyor.com/api/projects/status/eimlas99romrrro0?svg=true)](https://ci.appveyor.com/project/iWdGo/largeoutput)
![Build status](https://github.com/iwdgo/largeoutput/workflows/Go/badge.svg)

# Some ordinary exercises in modules

`go get github.com/iwdgo/largeouput/binedit`
- Edit a huge file named "data.bin" that does not fit in memory and delete every 7th byte of it.

`go get github.com/iwdgo/largeouput/modulo37`
- Output numbers following some rules. The module outputs sequentially the integers from 1 to 99
 but on some conditions prints a string instead. Implemented rules are:
  - when the integer is a multiple of 3 print “Open” instead of the number,
  - when it is a multiple of 7 print “Source” instead of the number,
  - when it is a multiple of both 3 and 7 print “OpenSource” instead of the number.

`go get github.com/iwdgo/largeouput/sumofconverts`
- Convert a set of strings which contains numbers. It returns the sum of the list items and skips others.
A recursive and non-recursive version are benchmarked.

## Handling output when using a file.

An ordinary exercise may have a large output for which the easy `// Output:`
becomes unreadable.

In that case a convenient structure is an `Example` for a reduced set of data, and
a test and its benchmark for larger outputs.

If you need to output to a file, you might have to update your code.
Piping `Stdout` to a file is easier. A helper to compare the file to a referenced file
returns eventually an error usable by the `testing` package.

The original code from `testing/example.go` is using unexported routines.
It has been adapted to ordinary exercises.

## Data files

The `/output` subdirectory avoids having the data files mixed with source code.
Module checks the existence of the directory but does not create it.
If the working directory (not the temp, nor the executing) is unavailable,
tests will panic.

The output directory is created by the tests that always run before benchmarks.
