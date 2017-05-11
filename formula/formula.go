package formula

import "math"

func CalculateInterest(name string, amount float64, interestRate float64, termInMonths int64) float64 {
	nominalRate := interestRate / 100
	termInYears := float64(termInMonths) / 12

	var interest float64

	switch name {
	case "bni", "imprebanca", "verso":
		interest = nonCompoundingFormula(amount, nominalRate, termInYears)
	default:
		interest = compoundingFormula(amount, nominalRate, termInYears)
	}
	return interest
}

func compoundingFormula(amount float64, nominalRate float64, termInYears float64) float64 {
	return (math.Pow((1 + nominalRate), termInYears) * amount) - amount
}

func nonCompoundingFormula(amount float64, nominalRate float64, termInYears float64) float64 {
	return amount * (1 + nominalRate * termInYears) - amount
}
