package domain

import (
	"fmt"
)

type Participant struct {
	Budget       int
	Contribution int
}

func NewParticipant(budget int) *Participant {
	if budget < 0 {
		panic("Budget must be a positive integer")
	}

	return &Participant{budget, 0}
}

func (participant *Participant) Contribute(contribution int) {
	totalContribution := participant.Contribution + contribution

	if totalContribution > participant.Budget {
		panic(fmt.Sprintf(`Total contribution of %d is over budget of %d`, totalContribution, participant.Budget))
	}

	participant.Contribution = totalContribution
}

func (participant *Participant) ContributeAsMuchAsPossible(targetContribution int) {
	totalContribution := Min(participant.Contribution+targetContribution, participant.Budget)
	participant.Contribution = totalContribution
}

func (participant *Participant) CanContribute() bool {
	return participant.Contribution < participant.Budget
}

func (participant *Participant) CalculateBudgetLeft() int {
	return participant.Budget - participant.Contribution
}
