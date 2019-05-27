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

# Assumptions

The coding challenge PDF encapsulates all requirements of the real problem (unlikely!).

First name and last name uniquely identify employee (very unlikely!).

Output 'pay period' calculation specified as PCM, but assumption is the output requires period start date and end date formatted as string.

Employee is assumed to be employed on payment start date and of continuous paid service until the last day of the period i.e. they did not cease employment within the period.

End date is always the last day of the month.  Start date is always the first of the month i.e. payee are always paid for entire month.  Initially, I was not going to make this assumption and calculate part-month payments but limited scope to get this done :).

A graphical user interface is not required.

Readability and extensibilty are favoured over code performance. Optimisation for very high throughput is not required.

A maxium 50% super rate is allowed.

Annual salary is a dollar whole figure, greater than zero, and less than $1 billion. 

First name and last name are both provided as UTF8 and are between 1 and 50 characters long.

CSV is UTF8 encoded with unix line endings. Strings are not quoted. Names may not contain commas.

CSV input does not contain headers. Maximum lines in file is 10 million.

The application runs on Linux Unbuntu 16.04.

Only payslips for FY2017-18 can be calculated.  i.e. start-date must be between 01-Jul-2017 and 01-Jun-2018.


# Solution Architecture and Design

CLI application which takes a single argument as the CSV input.  Errors and output are written to standard output / errror.

Implemented on golang.  I want to learn golang and seems a good fit to get the job done fast.

Packages: (not real go packages - TODO - make them so!)
- main
- file 
- tax-calc

- unit tests 
  x2 packages

- integration test / volume tests
  via CLI (TODO!)

- tax band config file JSON
