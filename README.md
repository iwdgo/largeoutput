# Testing when large output is expected

An ordinary exercise may have a large output for which the easy `// Output:`
becomes unreadable.

In that case a convenient structure is an `Example` for a reduced set of data and
a test and its benchmark for larger outputs.

If you need to output to a file, you might have to update your code.
Piping `Stdout` to a file is easier. A helper to compare the file to a referenced file
returns eventually an error usable by the `testing` package.

The original code from `testing/example.go` is using unexported routines.
It has been adapted to ordinary exercises.

## How to ?

The HTML reference file is missing but the other ones are available.
You can edit samples or add your own.
You can use the ordinary `go test` and `go test -bench=.`

## Data files

The `/output` subdirectory avoids to have the data files mixed with source code.
The directory is not created but its existence is checked.
If the working directory (not the temp, nor the executing) is unavailable,
tests will panic.

It is not needed to use the related test helper because a benchmark runs the related test
first. The test sets the working directory correctly.

## Testing package

This package contains testing of module github.com/iwdgo/testingfiles