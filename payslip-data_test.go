package payslips

import (
	"testing"
	"time"
)

func TestGetPaySlipData(t *testing.T) {
    LoadTaxData("tax.json")

    testName := "David Rudd"
    startDate := time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC)
    salary := 60050
    superRate := 900
    grossIncomeExp := 5004
    incomeTaxExp := 922
    netIncomeExp := 4082
    superExp := 450

    grossIncome, incomeTax, netIncome, super := GetPayslipData(salary, startDate, superRate)

    assert := getIntAsserter(t, testName)

    assert(grossIncome, grossIncomeExp, "Gross Income")
    assert(incomeTax, incomeTaxExp, "Income Tax")
    assert(netIncome, netIncomeExp, "Net Income")
    assert(super, superExp, "Super")

    testName = "Ryan Chen"
    startDate = time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC)
    salary = 120000
    superRate = 1000
    grossIncomeExp = 10000 
    incomeTaxExp = 2669 
    netIncomeExp = 7331 
    superExp = 1000 

    grossIncome, incomeTax, netIncome, super = GetPayslipData(salary, startDate, superRate)

    assert(grossIncome, grossIncomeExp, "Gross Income")
    assert(incomeTax, incomeTaxExp, "Income Tax")
    assert(netIncome, netIncomeExp, "Net Income")
    assert(super, superExp, "Super")
}