package game_social

import (
	"errors"
	"math/rand"
	"time"
)

const (
	Spade      uint8 = 0
	Heart      uint8 = 1
	Diamond    uint8 = 2
	Club       uint8 = 3
	SmallJoker uint8 = 4
	BigJoker   uint8 = 5
)

type Card struct {
	Value uint8
	Face  uint8
}

type Pot struct {
	Cards      []Card
	Reserve    []Card
	Active     []Card
	TrumpCards []Card
}

func NewPot() *Pot {
	return &Pot{}
}

func (p *Pot) BuildCards(numOfDecks int) {
	cards := make([]Card, 54*numOfDecks)
	start := 0
	for i := 0; i < numOfDecks; i++ {
		start = addDeck(cards, start)
	}
	p.Cards = cards
}

func (p *Pot) RemoveCard(card Card) {
	cards := make([]Card, len(p.Cards)-1)
	i := 0
	for _, c := range p.Cards {
		if c != card {
			cards[i] = c
			i++
		}
	}
	p.Cards = cards
}

func (p *Pot) ReservePot(reserve int) {
	num := len(p.Cards) - reserve
	p.Reserve = p.Cards[num:]
	p.Active = p.Cards[:num]
}

func (p *Pot) ShuffleCards() {
	shuffleCards(p.Cards)
}

func (p *Pot) FlipTrumpCards(cards []Card) {
	if p.isEligible(cards) {
		p.TrumpCards = cards
	}
}

func (p *Pot) isEligible(cards []Card) bool {
	return true
}

func (p *Pot) ReplaceReserve(cards []Card) {
	p.Reserve = cards
}

func (p *Pot) TakeCard() (Card, error){
	if len(p.Active) > 0 {
		card := p.Active[0]
		p.Active = p.Active[1:]
		return card, nil
	}
	return Card{}, errors.New("no more")
}

func (p *Pot) ShowReserve() {

}

func addDeck(cards []Card, start int) int{
	cards[start] = Card{Value: 0, Face: BigJoker}
	cards[start + 1] = Card{Value: 0, Face: SmallJoker}
	i := start + 2
	for value := 1; value <= 13; value++ {
		i = addSet(cards, uint8(value), i)
	}
	return i
}

func addSet(cards []Card, value uint8, i int) int {
	cards[i] = Card{Value: value, Face: Spade}
	cards[i+1] = Card{Value: value, Face: Heart}
	cards[i+2] = Card{Value: value, Face: Diamond}
	cards[i+3] = Card{Value: value, Face: Club}
	return i + 4
}

func shuffleCards(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}
