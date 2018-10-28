package main

import (
	"testing"
	"time"
)

func assertInts(t *testing.T, actual int, expected int, name string) {
	if actual != expected {
		t.Errorf("%s was incorrect, got: %d, want: %d.", name, actual, expected)
	}
}
func TestGetPaySlipData(t *testing.T) {

	// David Rudd
	startDate := time.Date(2017, time.March, 1, 0, 0, 0, 0, time.UTC)
	salary := 60050
	superRate := 900
	grossIncomeExp := 5004
	incomeTaxExp := 922
	netIncomeExp := 4082
	superExp := 450

	grossIncome, incomeTax, netIncome, super := GetPayslipData(salary, startDate, superRate)

	assertInts(t, grossIncome, grossIncomeExp, "Gross Income")
	assertInts(t, incomeTax, incomeTaxExp, "Income Tax")
	assertInts(t, netIncome, netIncomeExp, "Net Income")
	assertInts(t, super, superExp, "Super")

	// Ryan Chen
	startDate = time.Date(2017, time.March, 1, 0, 0, 0, 0, time.UTC)
	salary = 120000
	superRate = 1000
	grossIncomeExp = 5004
	incomeTaxExp = 922
	netIncomeExp = 4082
	superExp = 450

	grossIncome, incomeTax, netIncome, super = GetPayslipData(salary, startDate, superRate)

	assertInts(t, grossIncome, grossIncomeExp, "Gross Income")
	assertInts(t, incomeTax, incomeTaxExp, "Income Tax")
	assertInts(t, netIncome, netIncomeExp, "Net Income")
	assertInts(t, super, superExp, "Super")
}