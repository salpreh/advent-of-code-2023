package cardgame

import (
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const CardHandSize = 5

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

const (
	None CardType = 0
	Two  CardType = iota + 1
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	J
	Q
	K
	A
)

func CalculateGameTotalWinnings(cards []CardHand) int {
	playRound := CardPlayRound(cards)
	sort.Sort(playRound)

	totalWinnings := 0
	totalPlayers := len(playRound)
	for i, hand := range playRound {
		totalWinnings += hand.GetBid() * (totalPlayers - i)
	}

	return totalWinnings
}

func ParseRegularCardGame(input []string) []CardHand {
	hands := make([]CardHand, 0)
	for _, handData := range input {
		data := strings.Split(handData, " ")
		bid, _ := strconv.Atoi(data[1])
		hands = append(hands, NewRegularHand(data[0], bid))
	}

	return hands
}

type HandType int
type CardType int

type CardPlayRound []CardHand

func (r CardPlayRound) Len() int {
	return len(r)
}

func (r CardPlayRound) Less(i, j int) bool {
	lCard := r[i]
	rCard := r[j]

	return lCard.Compare(rCard) < 0
}

func (r CardPlayRound) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type CardHand interface {
	GetHandType() HandType
	Compare(cardHand CardHand) int
	GetCardTypes() []CardType
	GetCardType(i int) CardType
	GetBid() int
}

type RegularCardHand struct {
	CardTypes []CardType
	Bid       int
}

func NewRegularHand(cards string, bid int) *RegularCardHand {
	cardTypes := make([]CardType, 0)
	for _, cardChar := range cards {
		card, _ := parseCard(cardChar)
		cardTypes = append(cardTypes, card)
	}

	cardHand := &RegularCardHand{cardTypes, bid}

	return cardHand
}

func (h *RegularCardHand) GetCardTypes() []CardType {
	return h.CardTypes
}

func (h *RegularCardHand) GetCardType(i int) CardType {
	return h.CardTypes[i]
}

func (h *RegularCardHand) GetBid() int {
	return h.Bid
}

func (h *RegularCardHand) GetHandType() HandType {
	cardByType := make(map[CardType]int)
	for _, card := range h.CardTypes {
		count := utils.GetOrDefault(cardByType, card, 0)
		count += 1
		cardByType[card] = count
	}

	maxCard := 0
	secondMaxCard := 0
	for _, count := range cardByType {
		if count > maxCard {
			secondMaxCard = maxCard
			maxCard = count
		} else if count > secondMaxCard {
			secondMaxCard = count
		}
	}

	return getHandType(maxCard, secondMaxCard)
}

func (h *RegularCardHand) Compare(other CardHand) int {
	thisHand := h.GetHandType()
	otherHand := other.GetHandType()
	if thisHand > otherHand {
		return -1
	} else if otherHand > thisHand {
		return 1
	}

	return h.compareByStrongestCard(other)
}

func (h *RegularCardHand) compareByStrongestCard(other CardHand) int {
	for i, card := range h.CardTypes {
		otherCard := other.GetCardType(i)
		if card > otherCard {
			return -1
		} else if otherCard > card {
			return 1
		}
	}

	return 0
}

// sortCards Deprecated
func (h *CardHand) sortCards() {
	sortedCards := make([]CardType, CardHandSize)
	for _, card := range h.CardTypes {
		for i := 0; i < len(sortedCards); i++ {
			if sortedCards[i] > card {
				utils.AppendAndShift(sortedCards, i, card)
				break
			} else if sortedCards[i] == None {
				sortedCards[i] = card
				break
			}
		}
	}

	h.sortedCards = sortedCards
}

func getHandType(maxCardCount int, secondMaxCardCount int) HandType {
	if maxCardCount == 5 {
		return FiveOfKind
	} else if maxCardCount == 4 {
		return FourOfKind
	} else if maxCardCount == 3 && secondMaxCardCount == 2 {
		return FullHouse
	} else if maxCardCount == 3 {
		return ThreeOfKind
	} else if maxCardCount == 2 && secondMaxCardCount == 2 {
		return TwoPair
	} else if maxCardCount == 2 {
		return OnePair
	} else {
		return HighCard
	}
}

func parseCard(card rune) (CardType, error) {
	switch card {
	case '2':
		return Two, nil
	case '3':
		return Three, nil
	case '4':
		return Four, nil
	case '5':
		return Five, nil
	case '6':
		return Six, nil
	case '7':
		return Seven, nil
	case '8':
		return Eight, nil
	case '9':
		return Nine, nil
	case 'T':
		return Ten, nil
	case 'J':
		return J, nil
	case 'Q':
		return Q, nil
	case 'K':
		return K, nil
	case 'A':
		return A, nil
	default:
		return -1, fmt.Errorf("unknown card type: %s", card)
	}
}
