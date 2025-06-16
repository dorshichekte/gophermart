package entity

func NewBalance(current, Withdrawn float64) Balance {
	return Balance{
		Current:   current,
		Withdrawn: Withdrawn,
	}
}
