package validator

import (
	"strconv"
	"strings"
)

func seperateNum(s string) []int {
	splitStrNum := strings.Split(s, "")
	digits := make([]int, len(splitStrNum))
	for i, str := range splitStrNum {
		num, _ := strconv.Atoi(str)
		digits[i] = num
	}
	return digits
}
