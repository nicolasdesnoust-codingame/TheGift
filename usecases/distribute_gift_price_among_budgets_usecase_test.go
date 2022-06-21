package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeGiftPriceAmongBudgets_ShouldTellThatParticipantsHaveNotEnoughMoney(t *testing.T) {
	contributions := DistributeGiftPriceAmongBudgetsUsecase(120, []int{10, 20, 10})

	if len(contributions) != 0 {
		t.Fatalf(`Should not return any contribution when budgets are too low`)
	}
}

func TestDistributeGiftPriceAmongBudgets_ShouldFindContributionsWhenEveryoneCanPayTheSameAmount(t *testing.T) {
	contributions := DistributeGiftPriceAmongBudgetsUsecase(30, []int{10, 10, 10})

	assert.ElementsMatch(t, contributions, []int{10, 10, 10})
}

func TestDistributeGiftPriceAmongBudgets_ShouldFindContributionsForDifferentBudgets(t *testing.T) {
	contributions := DistributeGiftPriceAmongBudgetsUsecase(46, []int{5, 10, 14, 17})

	assert.ElementsMatch(t, contributions, []int{5, 10, 14, 17})
}
