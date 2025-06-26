package util

func IsValidOrderNumber(orderNumber string) bool {
	total := 0
	isSecondDigit := false

	for i := len(orderNumber) - 1; i >= 0; i-- {
		digit := int(orderNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit

		isSecondDigit = !isSecondDigit
	}

	return total%10 == 0
}
