package main

import (
	"io"
	"bufio"
    "encoding/csv"
    "flag"
    "fmt"
    "log"
    "math"
    "os"
    "regexp"
    "strconv"
    "time"
)

var inputFileName, outputFileName, taxConfigFileName string

func init() {
    const (
        usageInput = "Input CSV data file with payee details"
        usageOutput = "Output file for payslip details"
        usageTax = "Tax bracket JSON data"
    )
    flag.StringVar(&inputFileName, "i", "input.csv", usageInput)
    flag.StringVar(&outputFileName, "o", "output.csv", usageOutput)
    flag.StringVar(&taxConfigFileName, "t", "tax.json", usageTax)
    flag.Parse()
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -i input.csv\n", os.Args[0])
        flag.PrintDefaults()
    }
}

const MAX_NAME_LENGTH = 50 
const MAX_SALARY = 1000000000
const MAX_SUPER_BASIS_POINTS = 5000 
const DATE_LAYOUT = "02 January" // This is a weird golang thing
const DEFAULT_YEAR = 2017 

const CSV_INPUT_COLUMNS = 5

type PayslipInput struct {
    firstName   string 
    lastName    string 
    startDate   time.Time
    salary 		int
    superRate   int
}

type PayslipOutput struct {
    name        string
    startDate   time.Time
    endDate     time.Time
    grossIncome int
    incomeTax   int
    netIncome   int
    super       int
}

// These really should have tests
func validate(in PayslipInput) (bool) {
    errs := []string{}

    if len(in.firstName) > MAX_NAME_LENGTH {
        errs = append(errs, fmt.Sprintf(
            "First Name must not be greater than %d", MAX_NAME_LENGTH))
    }

    if len(in.lastName) > MAX_NAME_LENGTH {
        errs = append(errs, fmt.Sprintf(
            "Last Name must not be greater than %d", MAX_NAME_LENGTH))
    }

    if in.salary < 0 {
        errs = append(errs, 
            "Salary must not be negative")
    }

    if in.salary > MAX_SALARY {
        errs = append(errs, fmt.Sprintf(
            "Salary must not be greater than %d", MAX_SALARY))
    }

    if in.superRate < 0 {
        errs = append(errs, fmt.Sprintf(
            "Super rate basis points must not be negative"))
    }

    if in.superRate > MAX_SUPER_BASIS_POINTS {
        errs = append(errs, fmt.Sprintf(
            "Super rate basis points must not be greater than %d", MAX_SUPER_BASIS_POINTS))
    }

    for _, validationError := range errs {
        fmt.Println(validationError)
    }
    return len(errs) == 0
}

func GetPayslipData(in PayslipInput) (result PayslipOutput) {
    annualIncomeTax, err := getIncomeTax(in.salary, in.startDate)
    if err != nil {
        log.Fatal(err)
    }
    result.name        = in.firstName + " " + in.lastName
    result.grossIncome = getMonthlyAmount(in.salary)
    result.incomeTax   = getMonthlyAmount(annualIncomeTax)
    result.netIncome   = result.grossIncome - result.incomeTax
    result.super       = applyBasisPointRate(result.grossIncome, in.superRate)
    result.startDate   = in.startDate
    result.endDate     = in.startDate.AddDate(0, 1, -1)

    return result
}

// I probably should have chosen my own formats!
func parseDate(in string) (time.Time, error){
    re := regexp.MustCompile("^\\d+\\s+[a-zA-Z]+")
    match := re.FindString(in)
    t, err := time.Parse(DATE_LAYOUT, match)
    year := DEFAULT_YEAR
    if t.Month() < time.July {
        year++
    } 
    return t.AddDate(year, 0, 0), err
}

func formatDate(in time.Time) (string) {
    endOfMonth := in.AddDate(0, 1, -1)
    return in.Format(DATE_LAYOUT) + " – " + endOfMonth.Format(DATE_LAYOUT)
}

func parsePercentAsBasisPoints(in string) (int, error) {
    re := regexp.MustCompile("\\d+\\.?\\d*")
    match := re.FindString(in)
    float, err := strconv.ParseFloat(match, 32)
    return int(math.Round(float * 100)), err
}

func main() {
    if inputFileName == "" || outputFileName == "" {
        flag.Usage()
        os.Exit(1)	
    }
    loadTaxData(taxConfigFileName)

    file := openFile(inputFileName)
    defer file.Close()
    reader := csv.NewReader(bufio.NewReader(file))

    writeFile := createFile(outputFileName)
    defer writeFile.Close()
    writer := csv.NewWriter(writeFile)
    defer writer.Flush()

    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        if len(line) != CSV_INPUT_COLUMNS {
            log.Fatal("Input file format is incorrect")
        }
        salary, err := strconv.ParseInt(line[2], 10, 32)
        if err != nil {
            log.Fatal(err)
        }
        superRate, err := parsePercentAsBasisPoints(line[3])
        if err != nil {
            log.Fatal(err)
        }
        startDate, err := parseDate(line[4])
        if err != nil {
            log.Fatal(err)
        }
        input := PayslipInput{
            firstName: line[0],
            lastName: line[1],
            salary: int(salary),
            superRate: superRate,
            startDate: startDate,
        }
        if validate(input) {
            output := GetPayslipData(input)
            csvOutput := []string {
                output.name,
                formatDate(output.startDate),
                fmt.Sprintf("%d", output.grossIncome),
                fmt.Sprintf("%d", output.incomeTax),
                fmt.Sprintf("%d", output.netIncome),
                fmt.Sprintf("%d", output.super),
            }
            err := writer.Write(csvOutput)
            if err != nil {
                log.Fatal(err)
            }
        }
    }
}