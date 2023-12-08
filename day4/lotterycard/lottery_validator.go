package lotterycard

import (
	"com.github/salpreh/advent-of-code-2023/utils"
	"math"
	"strconv"
	"strings"
)

const (
	cardNumSep    = ":"
	lotteryNumSep = "|"
	numSep        = " "
)

func SumLotteryPoints(input []Card) int {
	points := 0
	for _, card := range input {
		points += card.GetCardPoints()
	}

	return points
}

func ParseLotteryCards(input []string) []Card {
	lotteryCards := make([]Card, 0)
	for num, line := range input {
		numSepIdx := strings.Index(line, cardNumSep)
		lotteryNumSepIdx := strings.Index(line, lotteryNumSep)

		winningNums := parseNumsLine(line[numSepIdx+2 : lotteryNumSepIdx])
		playerNums := parseNumsLine(line[lotteryNumSepIdx+1:])

		lotteryCards = append(lotteryCards, *NewLotteryCard(num+1, winningNums, playerNums))
	}

	return lotteryCards
}

func parseNumsLine(numLine string) []int {
	nums := make([]int, 0)
	for _, numItem := range strings.Split(numLine, numSep) {
		if num, err := strconv.Atoi(numItem); err == nil {
			nums = append(nums, num)
		}
	}

	return nums
}

type Card struct {
	Num         int
	WinningNums utils.Set[int]
	PlayerNums  []int
}

func NewLotteryCard(num int, winningNums []int, playerNums []int) *Card {
	winningSet := utils.NewSet(winningNums)

	return &Card{num, *winningSet, playerNums}
}

func (c *Card) GetPlayerWinningNums() []int {
	winningNums := make([]int, 0)
	for _, num := range c.PlayerNums {
		if c.WinningNums.Contains(num) {
			winningNums = append(winningNums, num)
		}
	}

	return winningNums
}

func (c *Card) GetCardPoints() int {
	winningNums := len(c.GetPlayerWinningNums())
	if winningNums == 0 {
		return 0
	}

	return int(math.Pow(2, float64(winningNums-1)))
}
