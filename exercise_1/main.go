package exercise_1

import "fmt"

func SalaryTaxes(employeeSalary float64) float64 {
	var tax float64

	if employeeSalary > 50000 {
		tax += employeeSalary * 0.17
	}

	if employeeSalary > 150000 {
		tax += employeeSalary * 0.1
	}

	fmt.Println("Tax total:", tax)
	return tax

}
