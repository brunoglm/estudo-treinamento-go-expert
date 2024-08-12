package tax

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := calculateTax(amount)
	return repository.SaveTax(tax)
}

func calculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}
