[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/largeoutput.svg)](https://pkg.go.dev/github.com/iwdgo/largeoutput)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/largeoutput)](https://goreportcard.com/report/github.com/iwdgo/largeoutput)
[![codecov](https://codecov.io/gh/iwdgo/largeoutput/branch/master/graph/badge.svg)](https://codecov.io/gh/iwdgo/largeoutput)

[![Build Status](https://app.travis-ci.com/iwdgo/largeoutput.svg?branch=master)](https://travis-ci.com/iwdgo/largeoutput)
[![Build Status](https://api.cirrus-ci.com/github/iwdgo/largeoutput.svg)](https://cirrus-ci.com/github/iwdgo/largeoutput)
[![Build status](https://ci.appveyor.com/api/projects/status/eimlas99romrrro0?svg=true)](https://ci.appveyor.com/project/iwdgo/largeoutput)
![Build status](https://github.com/iwdgo/largeoutput/workflows/Go/badge.svg)

# Some ordinary exercises using modules

## Binedit

Produce of a copy of a file named "data.bin" that does not fit in memory, while deleting every seventh byte.

## Module37

Output numbers following some rules. The module outputs sequentially the integers from 1 to 99
 but on some conditions prints a string instead. Implemented rules are:   
  - when the integer is a multiple of 3 print “Open” instead of the number,
  - when it is a multiple of 7 print “Source” instead of the number,
  - when it is a multiple of both 3 and 7 print “OpenSource” instead of the number.

## Sumofconverts

Convert a set of strings which contains numbers. It returns the sum of the list items and skips others.
A recursive and non-recursive version are benchmarked.

## Testing of modulo37

An ordinary exercise may have a large output for which the easy `// Output:`
becomes unreadable.

In this case a convenient structure is an `Example` for a reduced set of data, and
a test and its benchmark for larger outputs.

If you need to output to a file, you might have to update your code.
Piping `Stdout` to a file is easier. A helper to compare the file to a referenced file
returns eventually an error usable by the `testing` package.

The original code is from the commit history of `golang.go/src/testing/example.go` was using
unexported routines. It is adapted to these cases.

## Start

Download with `go get github.com/iwdgo/largeouput/<module name>`,
then `go test -v` for instance.

These commands may change with [module](https://github.com/golang/go/wiki/Modules) configuration.