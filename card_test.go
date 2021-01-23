package game_social

import "testing"

func TestPot_BuildCards(t *testing.T) {
	pot := NewPot()
	pot.BuildCards(1)
	if len(pot.Cards) != 54 {
		t.Errorf("Actual number of cards %d", len(pot.Cards))
	}
	if pot.Cards[0] != (Card{Face:BigJoker, Value: 0}) {
		t.Errorf("Actual Cards[0] %+v", pot.Cards[0])
	}
	if pot.Cards[1] != (Card{Face:SmallJoker, Value: 0}) {
		t.Errorf("Actual Cards[0] %+v", pot.Cards[1])
	}
	if pot.Cards[2] != (Card{Face:Spade, Value: 1}) {
		t.Errorf("Actual Cards[0] %+v", pot.Cards[2])
	}
	if pot.Cards[53] != (Card{Face:Club, Value: 13}) {
		t.Errorf("Actual Cards[0] %+v", pot.Cards[53])
	}
}

func TestPot_RemoveCard(t *testing.T) {
	pot := NewPot()
	pot.BuildCards(1)
	pot.RemoveCard(Card{Face: BigJoker, Value: 0})
	if len(pot.Cards) != 53 {
		t.Errorf("Actual number of cards %d", len(pot.Cards))
	}
	for _,c := range pot.Cards {
		t.Logf("%+v", c)
	}
}

func TestPot_ShuffleCards(t *testing.T) {
	pot := NewPot()
	pot.BuildCards(1)
	pot.ShuffleCards()
	t.Logf("Face:%+v", pot.Cards[12])
}

func TestGame_ReservePot(t *testing.T) {
	pot := NewPot()
	pot.BuildCards(1)
	pot.ReservePot(6)
	if len(pot.Reserve) != 6 {
		t.Errorf("Actual %d", len(pot.Reserve))
	}
	if pot.Reserve[0] != (Card{Value: 12, Face: Diamond}) {
		t.Errorf("Actual %+v", pot.Reserve[0])
	}
	if len(pot.Active) != 48 {
		t.Errorf("Actual %d", len(pot.Active))
	}
	if pot.Active[47] != (Card{Value: 12, Face: Heart}) {
		t.Errorf("Actual %+v", pot.Active[47])
	}
}

func TestPot_TakeCard(t *testing.T) {
	pot := NewPot()
	pot.BuildCards(1)
	pot.ReservePot(6)
	card,_ := pot.TakeCard()
	if card != (Card{Face: BigJoker, Value: 0}) {
		t.Errorf("Actual %+v", card)
	}
	if len(pot.Active) != 47 {
		t.Errorf("Actual %d", len(pot.Active))
	}
	card,_ = pot.TakeCard()
	card,_ = pot.TakeCard()
	if card != (Card{Face: Spade, Value: 1}) {
		t.Errorf("Actual %+v", card)
	}
	if len(pot.Active) != 45 {
		t.Errorf("Actual %d", len(pot.Active))
	}
}