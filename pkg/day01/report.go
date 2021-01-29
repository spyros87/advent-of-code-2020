package day01

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/spyros87/advent-of-code-2020/pkg/input"
)

const TotalSum = 2020

var NoNumbersFound = errors.New("no expenses resulting 2020 exist")

func ReportRepair() {
	fmt.Println("Day One Starts")
	sort.Ints(input.Day01())

	partOne(input.Day01())
	partTwo(input.Day01())
}

func partOne(expenses []int) {
	fmt.Println("Day One - Part One Starts")
	first, second, err := searchForTwoEntriesEqualTo(expenses, TotalSum)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("[%d] + [%d] = %d\n", first, second, first+second)
	fmt.Printf("[%d] * [%d] = %d\n", first, second, first*second)
}

func partTwo(expenses []int) {
	fmt.Println("Day One - Part Two Starts")
	first, second, third, err := searchForThreeEntriesEqualTo(expenses, TotalSum)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("[%d] + [%d] + [%d] = %d\n", first, second, third, first+second+third)
	fmt.Printf("[%d] * [%d] * [%d] = %d\n", first, second, third, first*second*third)
}

func searchForTwoEntriesEqualTo(expenses []int, expectedSum int) (first int, second int, er error) {
	for _, value := range expenses {

		rest := expectedSum - value
		index := sort.SearchInts(expenses, rest)

		if index < len(expenses) && expenses[index] == rest {
			return value, rest, nil
		}
	}
	return 0, 0, NoNumbersFound
}

func searchForThreeEntriesEqualTo(expenses []int, expectedSum int) (first int, second int, third int, er error) {
	for _, value := range expenses {

		rest := expectedSum - value
		second, third, _ := searchForTwoEntriesEqualTo(expenses, rest)

		if second != 0 && third != 0 {
			return value, second, third, nil
		}
	}
	return 0, 0, 0, NoNumbersFound
}
