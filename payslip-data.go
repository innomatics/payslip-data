package payslips

import (
	"log"
    "flag"
    "fmt"
    "os"
    "time"
)

var inputFileName, taxConfigFileName string

func init() {
    const (
        usageInput = "Input CSV data file with payee details"
        usageTax = "Tax bracket JSON data"
    )
    flag.StringVar(&inputFileName, "i", "", usageInput)
    flag.StringVar(&taxConfigFileName, "t", "tax.json", usageTax)
    flag.Parse()
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -i payee-details.csv\n", os.Args[0])
        flag.PrintDefaults()
    }
}

const MAX_NAME_LENGTH = 50 
const MAX_SALARY = 999000000000 
const MAX_SUPER_BASIS_POINTS = 5000 

type PayslipInput struct {
    firstName   string 
    lastName    string 
    startDate   time.Time
    salary 		int
    superRate   int
}

type PayslipOutput struct {
    endDate     time.Time
    grossIncome int
    incomeTax   int
    netIncome   int
    super       int
}

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
    annualIncomeTax, err := GetIncomeTax(in.salary, in.startDate)
    if err != nil {
        log.Fatal(err)
    }
    result.grossIncome = GetMonthlyAmount(in.salary)
    result.incomeTax   = GetMonthlyAmount(annualIncomeTax)
    result.netIncome   = result.grossIncome - result.incomeTax
    result.super       = ApplyBasisPointRate(result.grossIncome, in.superRate)
    result.endDate     = in.startDate.AddDate(0, 1, -1)

    return result
}

func main() {
    if inputFileName=="" {
        flag.Usage()
        os.Exit(1)	
    }
    LoadTaxData(taxConfigFileName)

    openFile(inputFileName)

    //openFile(outputFile)

    // For line in CSV file, output to stdout or output file

}