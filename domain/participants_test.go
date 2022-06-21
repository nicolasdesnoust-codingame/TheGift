package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverageContribution(t *testing.T) {
	participants := NewParticipants([]Participant{
		*NewParticipant(10),
		*NewParticipant(20),
		*NewParticipant(30)})

	actualAverageContribution := participants.CalculateAverageContribution(100)

	expectedAverageContribution := 33
	assert.Equal(t, expectedAverageContribution, actualAverageContribution)
}

func TestExtractContributionsInAscendingOrder(t *testing.T) {
	participant1 := *NewParticipant(10)
	participant1.Contribute(10)

	participant2 := *NewParticipant(20)
	participant2.Contribute(20)

	participant3 := *NewParticipant(30)
	participant3.Contribute(30)

	participants := NewParticipants([]Participant{
		participant1,
		participant2,
		participant3})

	actualContributions := participants.ExtractContributionsInAscendingOrder()

	assert.ElementsMatch(t, actualContributions, []int{10, 20, 30})
}

func TestExtractContributionsInAscendingOrder_ShouldSortContributionsInAscendingOrder(t *testing.T) {
	participant1 := *NewParticipant(30)
	participant1.Contribute(20)

	participant2 := *NewParticipant(30)
	participant2.Contribute(30)

	participant3 := *NewParticipant(30)
	participant3.Contribute(10)

	participants := NewParticipants([]Participant{
		participant1,
		participant2,
		participant3})

	actualContributions := participants.ExtractContributionsInAscendingOrder()

	assert.Exactly(t, actualContributions, []int{10, 20, 30})
}

func TestHasContributedEnoughForGift_ShouldReturnFalse(t *testing.T) {
	participants := NewParticipants([]Participant{
		*NewParticipant(10),
		*NewParticipant(10),
		*NewParticipant(10)})

	hasContributedEnough := participants.HasContributedEnoughFor(100)

	assert.False(t, hasContributedEnough)
}

func TestHasContributedEnoughForGift_ShouldReturnTrue(t *testing.T) {
	participant1 := *NewParticipant(30)
	participant1.Contribute(20)

	participant2 := *NewParticipant(30)
	participant2.Contribute(30)

	participant3 := *NewParticipant(30)
	participant3.Contribute(10)

	participants := NewParticipants([]Participant{
		participant1,
		participant2,
		participant3})

	hasContributedEnough := participants.HasContributedEnoughFor(60)

	assert.True(t, hasContributedEnough)
}

func TestCanAfford_ShouldReturnFalse(t *testing.T) {
	participants := NewParticipants([]Participant{
		*NewParticipant(10),
		*NewParticipant(10),
		*NewParticipant(10)})

	canAfford := participants.CanAfford(100)

	assert.False(t, canAfford)
}

func TestCanAfford_ShouldReturnTrue(t *testing.T) {
	participants := NewParticipants([]Participant{
		*NewParticipant(10),
		*NewParticipant(10),
		*NewParticipant(10)})

	canAfford := participants.CanAfford(30)

	assert.True(t, canAfford)
}

func TestCalculateTotalContribution_ShouldReturnValidValue(t *testing.T) {
	participant1 := *NewParticipant(30)
	participant1.Contribute(20)

	participant2 := *NewParticipant(30)
	participant2.Contribute(30)

	participant3 := *NewParticipant(30)
	participant3.Contribute(10)

	participants := NewParticipants([]Participant{
		participant1,
		participant2,
		participant3})

	totalContribution := participants.CalculateTotalContribution()

	assert.Equal(t, 60, totalContribution)
}
