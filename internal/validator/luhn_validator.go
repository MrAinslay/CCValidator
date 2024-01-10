package validator

import "fmt"

func ValidateNum(s string) bool {
	digits := seperateNum(s)

	var sum int
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i]
		if i > 0 {
			i--
			mult := digits[i] * 2

			if mult > 9 {
				spratedNum := seperateNum(fmt.Sprint(mult))
				sum += spratedNum[0] + spratedNum[1]
			} else {
				sum += mult
			}
		}
	}
	return sum%10 == 0
}
