package payslips

import (
	"math"
    "time"
)

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

func GetIncomeTax(salary int, startDate time.Time) (incomeTax int) {
	year := startDate.Year()
	if startDate.Month() <= time.June {
       year--
	}
	for _, taxYear := range taxData {
		if taxYear.Year == year {
			return getIncomeTaxFromBrackets(salary, taxYear.Brackets)
		}
	}
	return 0 // TBI	- raise error if no year data found
}

func getIncomeTaxFromBrackets(salary int, brackets []TaxBracket) (incomeTax int) {
	applicableBracket := brackets[0]
	for _, bracket := range brackets {
		if salary <= bracket.Income {
			break
		}
		applicableBracket = bracket
	}
	return applicableBracket.Base +
			ApplyBasisPointRate(salary - applicableBracket.Income, applicableBracket.Rate)
}

func ApplyBasisPointRate(value int, basisPointRate int) (int) {
	return int(math.Round(float64(value * basisPointRate) / 10000))
}

func GetMonthlyAmount(annualAmount int, startDate time.Time) (int) {
	return int(math.Round(float64(annualAmount) / 12))
}

func LoadTaxData(configFileName string) {
	f := openFile(configFileName)
	loadJSONFile(f, &taxData)
}
