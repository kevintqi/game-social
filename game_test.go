package game_social

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	g.JoinGame("player1")
	g.JoinGame("player2")
	g.JoinGame("player3")
	g.JoinGame("player4")
	g.SeatPlayer("player1", 0, Red)
	g.SeatPlayer("player2", 1, Blue)
	g.SeatPlayer("player3", 2, Red)
	g.SeatPlayer("player4", 3, Blue)
	g.MakeHost("player3")
	t.Logf("%+v", g.Table.ActivePlayer.Player)
	g.SetDecks(2)
	g.RemoveCard(Card{Face: BigJoker, Value: 0})
	g.RemoveCard(Card{Face: SmallJoker, Value: 0})
	g.ShuffleCards()
	g.ReservePot(6)
	t.Logf("%+v", g.Table)
	t.Logf("%+v", g.Pot)
	for len(g.Pot.Active) > 0 {
		t.Logf("%+v", g.Table.ActivePlayer.Player)
		g.PickCard(g.Table.ActivePlayer.Player)
	}
	for _,p := range g.Table.PlayerOrder {
		if len(p.Cards) != 25 {
			t.Fail()
		}
	}
	g.PickReserve(g.Table.Host)
	t.Logf("%+v", g.Table.Host)
	if len(g.Table.Host.Cards) != 31 {
		t.Fail()
	}
	if len(g.Pot.Reserve) != 0 {
		t.Fail()
	}
	t.Logf("%+v", g.Pot.Reserve)

	g.DropReserve(g.Table.Host, Card{Value: 2, Face: Club})
}
