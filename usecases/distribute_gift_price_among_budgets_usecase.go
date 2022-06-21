package usecases

import (
	"log"
	"thegift/domain"
)

func DistributeGiftPriceAmongBudgetsUsecase(giftPrice int, budgets []int) []int {

	if giftPrice > calculateSum(budgets) {
		return []int{}
	}

	allParticipants := domain.NewParticipantFactory().CreateParticipants(budgets)
	log.Printf("Gift price: %d€, number of participants: %d\n", giftPrice, len(allParticipants.Content))

	averageContribution := allParticipants.CalculateAverageContribution(giftPrice)
	log.Printf("Average contribution (lower bound) : %d€\n", averageContribution)
	log.Printf("Trying to make everyone contribute %d€\n", averageContribution)
	for index := range allParticipants.Content {
		allParticipants.Content[index].ContributeAsMuchAsPossible(averageContribution)
	}

	totalContribution := allParticipants.CalculateTotalContribution()
	log.Printf("Current total contribution achieved : %d/%d€\n", totalContribution, giftPrice)

	for totalContribution < giftPrice {
		participants := allParticipants.FilterOutThoseWhoGaveEverything()
		participantWithSmallestBudget := participants.Content[0]
		nextContribution := participantWithSmallestBudget.CalculateBudgetLeft()
		log.Printf("Everyone can give at least %d€\n", nextContribution)

		estimatedTotalContribution := totalContribution + nextContribution*len(participants.Content)
		log.Printf("If everyone would give %d€, we could reach %d€\n", nextContribution, estimatedTotalContribution)

		if estimatedTotalContribution <= giftPrice {
			for index := range participants.Content {
				participant := &participants.Content[index]
				participant.Contribute(nextContribution)
			}
			totalContribution = estimatedTotalContribution
		} else {
			shareContributionLeftAmongParticipantsLeft(giftPrice, totalContribution, nextContribution, participants)
			totalContribution = giftPrice
		}
	}

	return allParticipants.ExtractContributionsInAscendingOrder()
}

func shareContributionLeftAmongParticipantsLeft(giftPrice, totalContribution int, nextContribution int, participants *domain.Participants) {
	amountNeeded := giftPrice - totalContribution
	amountPerParticipant := amountNeeded / len(participants.Content)
	log.Println("totalContribution")
	log.Println(totalContribution)
	log.Println("amountNeeded")
	log.Println(amountNeeded)
	log.Println("amountPerParticipant")
	log.Println(amountPerParticipant)

	for index := range participants.Content {
		participants.Content[index].Contribute(amountPerParticipant)
		log.Println(participants.Content[index].Contribution)
	}
	participantsLeft := participants.FilterOutThoseWhoGaveEverything()
	amountLeft := amountNeeded % len(participants.Content)

	for i := 0; i < amountLeft; i++ {
		participantsLeft.Content[i].Contribute(1)
	}
}

func calculateSum(array []int) int {
	sum := 0

	for _, element := range array {
		sum += element
	}

	return sum
}
