package domain

import (
	"log"
	"sort"
)

type Participants struct {
	Content []Participant
}

func NewParticipants(participant []Participant) *Participants {
	return &Participants{participant}
}

func (allParticipants *Participants) DistributeAmongBudgets(giftPrice int) {
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
			participants.shareContributionLeftAmongParticipantsLeft(giftPrice, totalContribution, nextContribution)
			totalContribution = giftPrice
		}
	}
}

func (participants *Participants) shareContributionLeftAmongParticipantsLeft(giftPrice, totalContribution int, nextContribution int) {
	amountNeeded := giftPrice - totalContribution
	amountPerParticipant := amountNeeded / len(participants.Content)

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

func (participants *Participants) CanAfford(giftPrice int) bool {
	totalBudget := 0

	for _, participant := range participants.Content {
		totalBudget += participant.Budget
	}

	return totalBudget >= giftPrice
}

func (participants *Participants) CalculateAverageContribution(giftPrice int) int {
	return giftPrice / len(participants.Content)
}

func (participants *Participants) HasContributedEnoughFor(giftPrice int) bool {
	totalContribution := participants.CalculateTotalContribution()

	return totalContribution == giftPrice
}

func (participants *Participants) CalculateTotalContribution() int {
	totalContribution := 0

	for _, participant := range participants.Content {
		totalContribution += participant.Contribution
	}

	return totalContribution
}

func (participants *Participants) ExtractContributionsInAscendingOrder() []int {
	contributions := make([]int, len(participants.Content))

	for index := range participants.Content {
		contributions[index] = participants.Content[index].Contribution
	}

	orderedContributions := contributions[:]
	sort.Ints(orderedContributions)
	return orderedContributions
}

func (participants *Participants) FilterOutThoseWhoGaveEverything() *Participants {
	firstIndex := 0
	lastIndex := len(participants.Content)

	for _, participant := range participants.Content {
		if participant.CanContribute() {
			break
		}
		firstIndex++
	}

	return NewParticipants(participants.Content[firstIndex:lastIndex])
}
