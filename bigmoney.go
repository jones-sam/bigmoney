package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	startingMoney, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	bettingCycles, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	money := startingMoney
	bet := 1
	highestBet := 1
	losingStreak := 0
	highestLosingStreak := 0
	for i := 0; i < bettingCycles; i++ {
		coin := coinFlip()
		fmt.Printf("Current bet: $%v\n", bet)
		fmt.Printf("%v percent complete\n", float64(i)/float64(bettingCycles)*100)

		streakToGoBroke := int(math.Ceil(math.Log2(float64(money))))
		fmt.Printf("Number of times to lose in a row to go broke with current money: %v/%v\n", losingStreak, streakToGoBroke)
		fmt.Printf("Odds of going broke: %v%%\n", math.Pow(0.5, float64(streakToGoBroke-losingStreak))*100)

		if bet > highestBet {
			highestBet = bet
		}

		if coin == true {
			money += bet
			fmt.Printf("Won $%v\nTotal: $%v\n\n", bet, money)
			losingStreak = 0
			bet = 1.0
		} else {
			money -= bet
			fmt.Printf("Lost $%v\nTotal: $%v\n\n", bet, money)
			losingStreak++
			if losingStreak > highestLosingStreak {
				highestLosingStreak = losingStreak
			}
			if money <= 0 || money < bet*2 {
				fmt.Printf("You're broke!")
				break
			}

			bet *= 2

		}
	}

	fmt.Printf("\n\nStats:\n\n")
	fmt.Printf("Starting money: $%v\n", startingMoney)
	fmt.Printf("Final total: $%v\n", money)
	madeMoney := money >= startingMoney
	if madeMoney {
		fmt.Printf("You won $%v\n", money-startingMoney)
	} else {
		fmt.Printf("You lost $%v\n", startingMoney-money)
	}

	fmt.Printf("Highest losing streak: %v\n", highestLosingStreak)
	fmt.Printf(
		"Number of times to lose in a row to go broke from the initial money: %v\n",
		math.Ceil(math.Log2(float64(startingMoney))),
	)
	if madeMoney {
		fmt.Printf(
			"Number of times to lose in a row to go broke from your current money: %v\n",
			math.Ceil(math.Log2(float64(money))),
		)
	}
	fmt.Printf("Highest bet: $%v\n", highestBet)

}

func coinFlip() bool {
	num := rand.Intn(2)
	if num == 0 {
		return false
	} else {
		return true
	}
}
