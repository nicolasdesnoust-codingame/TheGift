package main

import (
	"fmt"
	"thegift/usecases"
)

func main() {
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

	contributions := usecases.DistributeGiftPriceAmongBudgetsUsecase(giftPrice, budgets)

	if len(contributions) == 0 {
		fmt.Println("IMPOSSIBLE")
	}

	for _, contribution := range contributions {
		fmt.Println(contribution)
	}
}
