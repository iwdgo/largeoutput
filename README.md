[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/largeoutput.svg)](https://pkg.go.dev/github.com/iwdgo/largeoutput)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/largeoutput)](https://goreportcard.com/report/github.com/iwdgo/largeoutput)
[![codecov](https://codecov.io/gh/iwdgo/largeoutput/branch/master/graph/badge.svg)](https://codecov.io/gh/iwdgo/largeoutput)

[![Build Status](https://api.cirrus-ci.com/github/iwdgo/largeoutput.svg)](https://cirrus-ci.com/github/iwdgo/largeoutput)
[![Build status](https://ci.appveyor.com/api/projects/status/eimlas99romrrro0?svg=true)](https://ci.appveyor.com/project/iwdgo/largeoutput)
[![Go](https://github.com/iwdgo/largeoutput/actions/workflows/go.yml/badge.svg)](https://github.com/iwdgo/largeoutput/actions/workflows/go.yml)

# Some ordinary exercises using modules

- **Binedit** : Produce a copy of a file named "data.bin" that does not fit in memory, while deleting every seventh byte.

- **Modulo37** : Output numbers following some rules. The module outputs sequentially the integers from 1 to 99.

- **Sumofconverts** : Convert a set of strings which contains numbers. It returns the sum of the items representing numbers and skips others.
A recursive and non-recursive version are benchmarked.

More documentation in each package.

# Testing

An ordinary exercise may have a large output for which the easy `// Output:`
becomes unreadable. In this case, output of reduced set of data to a file for which content
can be evaluated is more convenient. It is usable even without access to the original code.

Piping `Stdout` to a file is easier. A helper to compare the file to a reference file.
Comparison is done using module [github.com/iwdgo/testingfiles](https://github.com/iwdgo/testingfiles) which
returns errors usable by the `testing` package.

The original code for piping output is from the commit history of `golang.go/src/testing/example.go` was using
unexported routines. It is adapted to these cases.
