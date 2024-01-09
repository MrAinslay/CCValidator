package validator

import (
	"strconv"
	"strings"
)

func seperateNum(i int) []int {
	strNum := strconv.Itoa(i)
	splitStrNum := strings.Split(strNum, "")
	digits := make([]int, len(splitStrNum))
	for i, str := range splitStrNum {
		num, _ := strconv.Atoi(str)
		digits[i] = num
	}
	return digits
}
