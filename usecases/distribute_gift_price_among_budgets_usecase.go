package usecases

import (
	"log"
	"thegift/domain"
)

func DistributeGiftPriceAmongBudgetsUsecase(giftPrice int, budgets []int) []int {

	allParticipants := domain.NewParticipantFactory().CreateParticipants(budgets)
	log.Printf("Gift price: %d€, number of participants: %d\n", giftPrice, allParticipants.Size())

	if allParticipants.CanAfford(giftPrice) {
		allParticipants.DistributeAmongBudgets(giftPrice)
		return allParticipants.ExtractContributionsInAscendingOrder()
	}

	noContributions := []int{}
	return noContributions
}
