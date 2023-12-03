package main

import "fmt"

var (
	pennies  = 0
	nickles  = 0
	dimes    = 0
	quarters = 0
)

func calculateCents(dollar float32) (int, int, int, int) {
	cents := dollar * 100
	if cents-25 >= 0 {
		for cents >= 25 {
			cents = cents - 25
			quarters++
		}
	}
	if cents-10 >= 0 {
		for cents >= 10 {
			cents = cents - 10
			dimes++
		}
	}
	if cents-5 >= 0 {
		for cents >= 5 {
			cents = cents - 5
			nickles++
		}
	}
	if cents-1 >= 0 {
		for cents >= 1 {
			cents = cents - 1
			pennies++
		}
	}

	return pennies, nickles, dimes, quarters
}

func main() {
	p, n, d, q := calculateCents(57.89)
	fmt.Println(fmt.Sprintf("You have %v Pennies, %v Nickles, %v Dimes, and %v Quarters.", p, n, d, q))
}
