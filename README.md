# Testing when large output is expected

Some ordinary exercises which are tested using intermediate files.
The output is compared against a reference using a test helper which returns an error message.
This error message conforms to testing package.

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
