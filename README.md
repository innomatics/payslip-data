Response to coding challenge done over a few hours and to teach myself some Go.

# Usage

The program run with no arguments will process the default file `input.csv`.

The confiuration file `tax.json` constains tax bracket data and must be in the same directory as the executable.

```
# process input.csv in same directory
./payslip-data
```

Output will be written by default to `output.csv` (over-writing if exists).

Input or output files may be specified with the `-i` or `-o` flags respectively.

For a full list of options see:
```
./payslip-data --help
```

# Building

This program is written in Go. Install the binary release for your OS from https://golang.org/dl/ and follow installation instructions there to install.

## Compilation
```
go build
```

## Run Tests
```
go test
```
