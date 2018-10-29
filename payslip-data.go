package payslips

import (
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

func GetPayslipData(
		salary 		int,
		startDate   time.Time,
		superRate   int,
	) (
		grossIncome int,
		incomeTax   int,
		netIncome   int,
		super       int,
	) {
	grossIncome = GetMonthlyAmount(salary, startDate)
	incomeTax   = GetMonthlyAmount(GetIncomeTax(salary, startDate), startDate)
	netIncome   = grossIncome - incomeTax
	super       = ApplyBasisPointRate(grossIncome, superRate)

	return grossIncome, incomeTax, netIncome, super
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