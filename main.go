package main

import "fmt"

//the variables where the coin amount is stored
var (
	pennies  = 0
	nickles  = 0
	dimes    = 0
	quarters = 0
)

func calculateCents(dollar float32) (int, int, int, int) {
	//converts the dollar amount into cents first.
	cents := dollar * 100
	
	/*
	checks to see if the cent amount can be converted into the specific coin
	then it checks how many times it can go into it
	 */
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
	
	//returns the amount of each coin.
	return pennies, nickles, dimes, quarters
}

func main() {
	p, n, d, q := calculateCents(57.89)
	//Output: "You have 4 Pennies, 0 Nickles, 1 Dimes, and 231 Quarters."
	fmt.Println(fmt.Sprintf("You have %v Pennies, %v Nickles, %v Dimes, and %v Quarters.", p, n, d, q))
}
