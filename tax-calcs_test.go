package payslips

import (
	"fmt"
    "testing"
    "time"
)

func TestGetIncomeTax(t *testing.T) {
    loadTaxData("tax.json")

    type TestData struct {
        in int
        expected int
    }
    date := time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC)
    date2016 := time.Date(2017, time.June, 30, 0, 0, 0, 0, time.UTC)
    date2017a := time.Date(2017, time.July, 1, 0, 0, 0, 0, time.UTC)
    date2017b := time.Date(2018, time.June, 30, 0, 0, 0, 0, time.UTC)
    date2018 := time.Date(2018, time.July, 1, 0, 0, 0, 0, time.UTC)
    data := []TestData {
        {
            0, 0,
        },
        {
            18201, 0,
        },
        {
            18202, 0,
        },
        {
            18203, 1,
        },
        {
            20000, 342,
        },
        {
            37001, 3572,
        },
        {
            37002, 3573,
        },
        {
            37003, 3573,
        },
        {
            50000, 7797,
        },
        {
            87001, 19822,
        },
        {
            87002, 19823,
        },
        {
            87003, 19823,
        },
        {
            100000, 24632,
        },
        {
            180000, 54232,
        },
        {
            180001, 54232,
        },
        {
            180002, 54233,
        },
        {
            400000, 153232,
        },
    }

    for _, d := range data {
        actual, err := getIncomeTax(d.in, date)
        if err != nil {
            t.Error(err)
        }
        assert := getIntAsserter(t, fmt.Sprintf("$%d", d.in))
        assert(actual, d.expected, "Income Tax")
     }

     _, err := getIncomeTax(50000, date2016)
     if err == nil {
         t.Error("2016 date should be an error")
     } 
     _, err = getIncomeTax(50000, date2017a)
     if err != nil {
         t.Error("2017 date should not be an error")
     } 
     _, err = getIncomeTax(50000, date2017b)
     if err != nil {
         t.Error("2017 date should not be an error")
     } 
     _, err = getIncomeTax(50000, date2018)
     if err == nil {
         t.Error("2018 date should be an error")
     } 

}