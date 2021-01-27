package day01

import "errors"

const TotalSum = 2020

func ReportRepair(expenses []int) (int, error) {
	first, second := check(expenses)

	if first == -1 && second == -1 {
		return 0, errors.New("no expenses resulting 2020 exist")
	}

	return expenses[first] * expenses[second], nil
}

func check(expenses []int) (int, int) {

	var firstPosition = -1
	var secondPosition = -1

	for i, firstValue := range expenses {
		for j, secondValue := range expenses[(i + 1):] {

			if sum2020(firstValue, secondValue) {
				firstPosition = i
				secondPosition = (i + 1) + j

				return firstPosition, secondPosition
			}

		}
	}
	return -1, -1
}

func sum2020(firstValue int, secondValue int) bool {
	return (firstValue + secondValue) == TotalSum
}
