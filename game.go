package game_social

type Hand struct {
	UserId string
	Cards  []Card
}

type Round struct {
	Id           uint16
	Hands        []Hand
	Winner       *Hand
	Bets         []Bet
	Rewards      uint64
	ActivePlayer *Player
}

type Bet struct {
	UserId string
	Amount uint64
}

type Game struct {
	Pot            Pot
	Rounds         []Round
	Bets           []Bet
	Rewards        uint64
	Teams          [2]string
	DominatedCards []*Card
	Table          *PlayerPool
}

func NewGame() *Game {
	table := NewPlayerPool()
	game := Game{Table: table}
	return &game
}

func (g *Game) JoinGame(playerId string) *Game {
	g.Table.AddPlayer(playerId)
	return g
}

func (g *Game) MakeHost(playerId string) *Game {
	g.Table.SetHost(playerId)
	return g
}

func (g *Game) SeatPlayer(playerId string, order uint8, team uint8) *Game {
	g.Table.SetPlayerOrder(playerId, order, team)
	return g
}

func (g *Game) UpdateActivePlayer() *Game {
	g.Table.UpdateActivePlayer()
	return g
}

func (g *Game) SetDecks(numOfDecks int) *Game {
	g.Pot.BuildCards(numOfDecks)
	return g
}

func (g *Game) RemoveCard(card Card) *Game {
	g.Pot.RemoveCard(card)
	return g
}

func (g *Game) ReservePot(reserve int) *Game {
	g.Pot.ReservePot(reserve)
	return g
}

func (g *Game) ShuffleCards() *Game {
	g.Pot.ShuffleCards()
	return g
}

func (g *Game) FlipTrumpCards(cards []Card) *Game {
	g.Pot.FlipTrumpCards(cards)
	return g
}

func (g *Game) PickReserve(player *Player) *Game {
	player.Cards = append(player.Cards, g.Pot.Reserve...)
	g.Pot.Reserve = g.Pot.Reserve[len(g.Pot.Reserve):]
	return g
}

func (g *Game) ReplacePot(cards []Card) *Game {
	g.Pot.ReplaceReserve(cards)
	return g
}

func (g *Game) PickCard(player *Player) *Game{
	card, err := g.Pot.TakeCard()
	if err == nil {
		player.Cards = append(player.Cards, card)
		g.UpdateActivePlayer()
	}
	return g
}

func (g *Game) DropReserve(host *Player, card Card) *Game {
	return g
}