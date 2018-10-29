package payslips

import (
	"fmt"
    "testing"
    "time"
)

type PayslipTestData struct {
    in PayslipInput
    expected PayslipOutput
}

func TestGetPaySlipData(t *testing.T) {
    loadTaxData("tax.json")

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

func TestParseDate(t *testing.T) {
    s1 := "01 June – 31 March"
    s2 := "18 July – 31 March"
    s3 := "18 Oranges – 31 March"
    e1 := time.Date(2018, 6, 1, 0, 0, 0, 0, time.UTC)
    e2 := time.Date(2017, 7, 18, 0, 0, 0, 0, time.UTC)
    a1, err1 := parseDate(s1)
    a2, err2 := parseDate(s2)
    _, err3 := parseDate(s3)
    if err1 != nil {
        t.Error(err1)
    }
    if err2 != nil {
        t.Error(err2)
    }
    if err3 == nil {
        t.Error("Oranges should error")
    }
    if e1 != a1 || e2 != a2 {
        t.Error("Dates not parsed")
    }
}

func TestFormatDate(t *testing.T) {
    d1 := time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC)
    d2 := time.Date(2017, 7, 18, 0, 0, 0, 0, time.UTC)
    e1 := "01 March – 31 March"
    e2 := "18 July – 17 August"
    a1 := formatDate(d1)
    a2 := formatDate(d2)
    fmt.Println(a2)
    if e1 != a1 || e2 != a2 {
        t.Error("Date format was wrong")
    }
}