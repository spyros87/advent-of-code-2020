package day02

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spyros87/advent-of-code-2020/pkg/input"
)

func PasswordPhilosophy() {
	fmt.Println("Day Two Starts")

	partOne(input.Day02())
}

func partOne(passwordPolicies []string) {
	fmt.Println("Day Two - Part One Starts")

	validPasswords := findValidPasswordsFrom(passwordPolicies)

	fmt.Printf("Valid Passwords: [%d]\n", validPasswords)
}

func findValidPasswordsFrom(passwordPolicies []string) int {
	regex := regexp.MustCompile("^(\\d+)-(\\d+)\\s(\\w+):\\s(\\w+)$")
	validPasswords := 0

	for _, passwordPolicy := range passwordPolicies {
		matchingBlocks := regex.FindStringSubmatch(passwordPolicy)

		min := asInt(matchingBlocks[1])
		max := asInt(matchingBlocks[2])

		count := strings.Count(matchingBlocks[4], matchingBlocks[3])

		if count >= min && count <= max {
			validPasswords++
		}
	}
	return validPasswords
}

func asInt(numberAsString string) int {
	number, err := strconv.Atoi(numberAsString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return number
}
