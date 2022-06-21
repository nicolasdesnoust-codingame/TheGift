package domain

import (
	"log"
	"sort"
)

type ParticipantFactory struct {
}

func NewParticipantFactory() *ParticipantFactory {
	return &ParticipantFactory{}
}

func (factory *ParticipantFactory) CreateParticipants(budgets []int) *Participants {
	orderedBudgets := budgets[:]
	sort.Ints(orderedBudgets)

	participants := make([]Participant, len(orderedBudgets))
	log.Printf("fef %d", len(budgets))

	for budgetIndex, budget := range orderedBudgets {
		participants[budgetIndex] = *NewParticipant(budget)
	}
	return NewParticipants(participants)
}
