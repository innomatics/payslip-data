package main

import (
	"encoding/json"
    "flag"
	"fmt"
	"io/ioutil"
	"log"
    "os"
    "time"
)

var inputFile, taxFile string

type TaxBracket struct {
	Income int
	Base int
	Rate int
}

type TaxYear struct {
	Year int
	Brackets []TaxBracket
}
var taxData []TaxYear

func init() {
	const (
		usageInput = "Input CSV data file with payee details"
		usageTax = "Tax bracket JSON data"
	)
	flag.StringVar(&inputFile, "input", "", usageInput)
	flag.StringVar(&inputFile, "i", "", usageInput + " (shorthand)")
	flag.StringVar(&taxFile, "tax", "tax.json", usageTax)
	flag.Parse()
	flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -i payee-details.csv\n", os.Args[0])
        flag.PrintDefaults()
	}
}

func openFile(fileName string) (file *os.File) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func getIncomeTax(salary int, year int) (incomeTax int) {
	return 0 // TBI	
}

func GetPayslipData(
	salary 		int,
	startDate   time.Time,
	super_rate  int,
	) (
	grossIncome int,
	incomeTax   int,
	netIncome   int,
	super       int,
) {
	return 0, 0, 0, 0 // TBI
}

func main() {
	if inputFile=="" {
		flag.Usage()
		os.Exit(1)	
	}
	f := openFile(taxFile)
	taxFileBytes, _ := ioutil.ReadAll(f)

	taxFileErr := json.Unmarshal(taxFileBytes, &taxData)
	if taxFileErr != nil {
		log.Fatal(taxFileErr)
	}
	fmt.Println(taxData)

	openFile(inputFile)

	//openFile(outputFile)

	// For line in CSV file, output to stdout or output file

}