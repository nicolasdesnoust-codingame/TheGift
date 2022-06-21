package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContribute_ShouldTellThatOneContributionOverBudgetAreNotFeasible(t *testing.T) {
	participant := NewParticipant(50)

	logicUnderTest := func() {
		participant.Contribute(60)
	}

	assert.Panics(t, logicUnderTest, "The code did not panic")
}

func TestContribute_ShouldTellThatMultipleContributionsOverBudgetAreNotFeasible(t *testing.T) {
	participant := NewParticipant(50)

	logicUnderTest := func() {
		participant.Contribute(10)
		participant.Contribute(41)
	}

	assert.Panics(t, logicUnderTest, "The code did not panic")
}

func TestContributeAsMuchAsPossible_ShouldNotContributeAboveBudget(t *testing.T) {
	participant := NewParticipant(50)

	participant.ContributeAsMuchAsPossible(10)
	participant.ContributeAsMuchAsPossible(50)
	participant.ContributeAsMuchAsPossible(10)

	assert.Equal(t, 50, participant.Contribution)
}

func TestContributeAsMuchAsPossible_ShouldAcceptBellowBudget(t *testing.T) {
	participant := NewParticipant(50)

	participant.ContributeAsMuchAsPossible(10)
	participant.ContributeAsMuchAsPossible(10)

	assert.Equal(t, 20, participant.Contribution)
}

func TestCanContribute_ShouldTellParticipantCanContribute(t *testing.T) {
	participant := NewParticipant(10)
	participant.Contribute(5)

	canContribute := participant.CanContribute()

	assert.True(t, canContribute)
}

func TestCanContribute_ShouldTellParticipantCanNotContribute(t *testing.T) {
	participant := NewParticipant(10)
	participant.Contribute(10)

	canContribute := participant.CanContribute()

	assert.False(t, canContribute)
}
