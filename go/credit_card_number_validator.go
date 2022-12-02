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
	var alternate bool

	for i := len(creditCardNumber) - 1; i >= 0; i-- {
		var n int = int(creditCardNumber[i] - '0')
		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}
		sum += n
		alternate = !alternate
	}

	return sum%10 == 0
}
