package createphonenumber

import (
	"fmt"
	"strconv"
)

func printArray(slice []uint) string {
	var res string
	for _, value := range slice {
		res += strconv.Itoa(int(value))
	}
	return res
}

func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", printArray(numbers[0:3]), printArray(numbers[3:6]), printArray(numbers[6:]))

}
