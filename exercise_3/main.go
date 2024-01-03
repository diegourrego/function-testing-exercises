package exercise_3

func CalculateSalary(hoursPerMonth float32, category string) (salary float32, err string) {
	switch category {
	case "C":
		salary = 1000 * hoursPerMonth
	case "B":
		salary = (1500 * hoursPerMonth) + calculateMonthlyBonus(1500, 0.2)
	case "A":
		salary = (3000 * hoursPerMonth) + calculateMonthlyBonus(3000, 0.5)
	default:
		err = "Unexpected category. Check it and try again."
		salary = 0.0

	}
	return
}

func calculateMonthlyBonus(baseSalary float32, percentage float32) (salaryBonus float32) {
	salaryBonus = (baseSalary * 8 * 30) * percentage
	return
}
