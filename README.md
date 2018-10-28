#Set Up



#Assumptions

The coding challenge PDF encapsulates all requirements of the real problem (unlikely!).

First name and last name uniquely identify employee (very unlikely!).

Output 'pay period' calculation specified as PCM, but assumption is the output requires period start date and end date formatted as string.

Employee is assumed to be employed on on payment start date and of continous paid service until the last day of the period i.e. they did not cease employment within the period.

End date is always the last day of the month.  If the start date is not the 1st of the month, the pay period is only from that date, until the end of the month i.e. the pay period my be shorter that one month.

A graphical user interface is not required.

Readability and extensibilty are favoured over code performance. Optimisation for very high throughput is not required.

Super rate is provided at no finer than basis point resolution i.e. 0.01%. Otherwise and error will occur.

Annual salary is a dollar whole figure, greater than zero, and less than $999 billion. 

First name and last name are both provided as UTF8 and are between 1 and 50 characters long.

CSV is UTF8 encoded with unix line endings. Strings are not quoted. Names may not contain commas.

CSV input does not contain headers. Maximum lines in file is 10 million.

The application runs on Linux Unbuntu 16.04.

Only payslips for FY2017-18 can be calculated.  i.e. start-date must be between 01-Jul-2017 and 30-Jun-2018.


#Solution Architecure and Design

CLI application which takes a single argument as the CSV input.  Errors and output are written to standard output / errror.

Implemented on golang.  I wont to learn golang and seems a good fit.

Packages:
- main
- payslip-calculator
- tax-bracket-import
- json
- csv

- unit tests 
  x2 packages

- integration test
  CLI



# Variables

(
	first_name,
	last_name,
	salary_per_year,
	super_rate,
	start_date
) => 
(
	first_name,
	last_name,
	period_start,
	period_end,
	gross_income,
	income_tax,
	net_income,
	super
)
