package main

import (
    "fmt"
    "errors"
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

func getIncomeTax(salary int, startDate time.Time) (incomeTax int, err error) {
    year := startDate.Year()
    if startDate.Month() <= time.June {
       year--
    }
    for _, taxYear := range taxData {
        if taxYear.Year == year {
            return getIncomeTaxFromBrackets(salary, taxYear.Brackets), nil
        }
    }
    return 0, errors.New(
        fmt.Sprintf("Missing tax bracket config for %d/%d", year, year + 1))
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
        applyBasisPointRate(salary - applicableBracket.Income, applicableBracket.Rate)
}

func applyBasisPointRate(value int, basisPointRate int) (int) {
    return int(math.Round(float64(value * basisPointRate) / 10000))
}

func getMonthlyAmount(annualAmount int) (int) {
    return int(math.Round(float64(annualAmount) / 12))
}

func loadTaxData(configFileName string) {
    f := openFile(configFileName)
    loadJSONFile(f, &taxData)
}
