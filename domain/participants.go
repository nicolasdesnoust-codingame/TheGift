package domain

import (
	"log"
	"sort"
)

type Participants struct {
	content []Participant
}

func NewParticipants(participant []Participant) *Participants {
	return &Participants{participant}
}

func (allParticipants *Participants) DistributeAmongBudgets(giftPrice int) {
	averageContribution := allParticipants.CalculateAverageContribution(giftPrice)
	log.Printf("Average contribution (lower bound) : %d€\n", averageContribution)

	log.Printf("Trying to make everyone contribute %d€\n", averageContribution)
	for index := range allParticipants.content {
		allParticipants.content[index].ContributeAsMuchAsPossible(averageContribution)
	}

	totalContribution := allParticipants.CalculateTotalContribution()
	log.Printf("Current total contribution achieved : %d/%d€\n", totalContribution, giftPrice)

	for totalContribution < giftPrice {
		participants := allParticipants.FilterOutThoseWhoGaveEverything()
		participantWithSmallestBudget := participants.content[0]
		nextContribution := participantWithSmallestBudget.CalculateBudgetLeft()
		log.Printf("Everyone can give at least %d€\n", nextContribution)

		estimatedTotalContribution := totalContribution + nextContribution*len(participants.content)
		log.Printf("If everyone would give %d€, we could reach %d€\n", nextContribution, estimatedTotalContribution)

		if estimatedTotalContribution <= giftPrice {
			for index := range participants.content {
				participant := &participants.content[index]
				participant.Contribute(nextContribution)
			}
			totalContribution = estimatedTotalContribution
		} else {
			participants.shareContributionLeftAmongParticipantsLeft(giftPrice, totalContribution, nextContribution)
			totalContribution = giftPrice
		}
	}
}

func (participants *Participants) shareContributionLeftAmongParticipantsLeft(giftPrice, totalContribution int, nextContribution int) {
	amountNeeded := giftPrice - totalContribution
	amountPerParticipant := amountNeeded / len(participants.content)

	for index := range participants.content {
		participants.content[index].Contribute(amountPerParticipant)
		log.Println(participants.content[index].Contribution)
	}
	participantsLeft := participants.FilterOutThoseWhoGaveEverything()
	amountLeft := amountNeeded % len(participants.content)

	for i := 0; i < amountLeft; i++ {
		participantsLeft.content[i].Contribute(1)
	}
}

func (participants *Participants) CanAfford(giftPrice int) bool {
	totalBudget := 0

	for _, participant := range participants.content {
		totalBudget += participant.Budget
	}

	return totalBudget >= giftPrice
}

func (participants *Participants) CalculateAverageContribution(giftPrice int) int {
	return giftPrice / len(participants.content)
}

func (participants *Participants) HasContributedEnoughFor(giftPrice int) bool {
	totalContribution := participants.CalculateTotalContribution()

	return totalContribution == giftPrice
}

func (participants *Participants) CalculateTotalContribution() int {
	totalContribution := 0

	for _, participant := range participants.content {
		totalContribution += participant.Contribution
	}

	return totalContribution
}

func (participants *Participants) ExtractContributionsInAscendingOrder() []int {
	contributions := make([]int, len(participants.content))

	for index := range participants.content {
		contributions[index] = participants.content[index].Contribution
	}

	orderedContributions := contributions[:]
	sort.Ints(orderedContributions)
	return orderedContributions
}

func (participants *Participants) FilterOutThoseWhoGaveEverything() *Participants {
	firstIndex := 0
	lastIndex := len(participants.content)

	for _, participant := range participants.content {
		if participant.CanContribute() {
			break
		}
		firstIndex++
	}

	return NewParticipants(participants.content[firstIndex:lastIndex])
}

func (participants *Participants) Size() int {
	return len(participants.content)
}
