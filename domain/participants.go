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

func (participants *Participants) CalculateAverageContribution(giftPrice int) int {
	return giftPrice / len(participants.Content)
}

func (participants *Participants) CanPay(giftPrice int) bool {
	amountRaised := participants.CalculateTotalContribution()

	return amountRaised == giftPrice
}

func (participants *Participants) CalculateTotalContribution() int {
	amountRaised := 0

	for _, participant := range participants.Content {
		amountRaised += participant.Contribution
	}

	return amountRaised
}

func (participants *Participants) ExtractContributions() []int {
	contributions := make([]int, len(participants.Content))

	for index := range participants.Content {
		contributions[index] = participants.Content[index].Contribution
		log.Println(participants.Content[index].Contribution)
	}

	orderedContributions := contributions[:]
	sort.Ints(orderedContributions)
	return orderedContributions
}

func (participants *Participants) FilterOutThoseWhoGaveEverything() *Participants {
	// participantsLeft := make([]Participant, 0)

	// for index := range participants.Content {
	// 	participant := &participants.Content[index]
	// 	if participant.Contribution < participant.Budget {
	// 		participantsLeft = append(participantsLeft, *participant)
	// 	}
	// }

	firstIndex := 0
	lastIndex := len(participants.Content)

	for i := 0; i < len(participants.Content); i++ {
		participant := &participants.Content[i]
		if participant.Contribution < participant.Budget {
			break
		}
		firstIndex++
	}

	return NewParticipants(participants.Content[firstIndex:lastIndex])
}
