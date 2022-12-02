package main

import (
	"flag"
	"fmt"
)

func main() {
	var creditCardNumber string
	flag.StringVar(&creditCardNumber, "credit_card_number", "0", "Credit card number")
	flag.Parse()

	validateCreditCardNumber(creditCardNumber)

	fmt.Println("Entered credit card number:", creditCardNumber)

	if checkLuhnAlgorithm(creditCardNumber) {
		fmt.Println("Credit card number is valid")
	} else {
		fmt.Println("Credit card number is invalid")
	}
}

func validateCreditCardNumber(creditCardNumber string) string {
	if creditCardNumber == "" {
		fmt.Println("Please provide a credit card number by running the program with the -credit_card_number flag")
	}

	if len(creditCardNumber) != 16 {
		fmt.Println("Credit card number must be 16 digits")
	}

	return ""
}

func checkLuhnAlgorithm(creditCardNumber string) bool {
	var sum int
	for i, digit := range creditCardNumber {
		if i%2 == 0 {
			sum += int(digit - '0')
		} else {
			sum += 2 * int(digit-'0')
		}
	}

	return sum%10 == 0
}
