package domain

import (
	"sort"
)

type Participants struct {
	Content []Participant
}

func NewParticipants(participant []Participant) *Participants {
	return &Participants{participant}
}

func (participants *Participants) CalculateAverageContribution(giftPrice int) int {
	return giftPrice / len(participants.Content)
}

func (participants *Participants) CanPay(giftPrice int) bool {
	amountRaised := participants.CalculateTotalContribution()

	return amountRaised == giftPrice
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
