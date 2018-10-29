package payslips

import (
    "testing"
    "time"
)

type PayslipTestData struct {
    in PayslipInput
    expected PayslipOutput
}

func TestGetPaySlipData(t *testing.T) {
    LoadTaxData("tax.json")

    data := []PayslipTestData {
        PayslipTestData {
            in: PayslipInput {
                firstName: "David",
                lastName: "Rudd",
                startDate: time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
                salary: 60050,
                superRate: 900,
            },
            expected: PayslipOutput {
                grossIncome: 5004,
                incomeTax: 922,
                netIncome: 4082,
                super: 450,
            },
        },
        PayslipTestData {
            in: PayslipInput {
                firstName: "Ryan",
                lastName: "Chen",
                startDate: time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
                salary: 120000,
                superRate: 1000,
            },
            expected: PayslipOutput {
                grossIncome: 10000,
                incomeTax: 2669,
                netIncome: 7331,
                super: 1000,
            },
        },
    }

    for _, d := range data {
        actual := GetPayslipData(d.in)
        assert := getIntAsserter(t, d.in.firstName + " " + d.in.lastName)
        assert(actual.grossIncome, d.expected.grossIncome, "Gross Income")
        assert(actual.incomeTax, d.expected.incomeTax, "Income Tax")
        assert(actual.netIncome, d.expected.netIncome, "Net Income")
        assert(actual.super, d.expected.super, "Super")
     }
}