package domain

import (
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

	for budgetIndex, budget := range orderedBudgets {
		participants[budgetIndex] = *NewParticipant(budget)
	}
	return NewParticipants(participants)
}
