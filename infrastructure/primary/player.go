package main

import (
	"fmt"
	"thegift/usecases"
)

func main() {
	giftPrice, budgets := parseInputs()

	contributions := usecases.DistributeGiftPriceAmongBudgetsUsecase(giftPrice, budgets)

	printAnswer(contributions)
}

func parseInputs() (int, []int) {
	var participantCount int
	fmt.Scan(&participantCount)

	var giftPrice int
	fmt.Scan(&giftPrice)

	budgets := make([]int, participantCount)
	for i := 0; i < participantCount; i++ {
		var budget int
		fmt.Scan(&budget)
		budgets[i] = budget
	}

	return giftPrice, budgets
}

func printAnswer(contributions []int) {
	if len(contributions) == 0 {
		fmt.Println("IMPOSSIBLE")
	}

	for _, contribution := range contributions {
		fmt.Println(contribution)
	}
}
