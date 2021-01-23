package game_social

type Call interface {
	Do()
}

type Command struct {
	Originator string
	Timestamp uint64
	Action string
}

type GameState struct {
	//active player
	//players
	//player team color
	//player seats
	//cards in the pot
	//cards player holds
	//trump cards
	//reward score
	//winner of a round
	//winner of game
}

func CallForGame() {
	// create Social
	// create Bets
	// create GameBundle
}


func SortHand() {

}

func ShowHand() {

}

func WinRound() {

}

func ShowPot() {

}

func WinGame() {

}