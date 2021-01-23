package game_social

import "testing"

func TestPlayerPool_AddPlayer(t *testing.T) {
	table := NewPlayerPool()
	table.AddPlayer("playerId")
	assertStrEq(t, "playerId", table.Players["playerId"].UserId)
}

func TestPlayerPool_SetHost(t *testing.T) {
	table := NewPlayerPool()
	table.AddPlayer("playerOne")
	table.AddPlayer("playerTwo")
	table.AddPlayer("playerThree")
	table.SetHost("playerTwo")
	assertStrEq(t, "playerTwo", table.Host.UserId)
}

func TestPlayerPool_SetPlayerOrder(t *testing.T) {
	table := NewPlayerPool()
	table.AddPlayer("playerOne")
	table.AddPlayer("playerTwo")
	table.AddPlayer("playerThree")
	table.AddPlayer("playerFour")
	table.SetPlayerOrder("playerOne", 0, Red)
	table.SetPlayerOrder("playerTwo", 1, Blue)
	table.SetPlayerOrder("playerThree", 2, Red)
	table.SetPlayerOrder("playerFour", 3, Blue)
	assertStrEq(t, "playerOne", table.PlayerOrder[0].UserId)
	assertUint8Eq(t, Red, table.PlayerOrder[0].TeamColor)
	assertStrEq(t, "playerTwo", table.PlayerOrder[1].UserId)
	assertUint8Eq(t, Blue, table.PlayerOrder[1].TeamColor)
	assertStrEq(t, "playerThree", table.PlayerOrder[2].UserId)
	assertUint8Eq(t, Red, table.PlayerOrder[2].TeamColor)
	assertStrEq(t, "playerFour", table.PlayerOrder[3].UserId)
	assertUint8Eq(t, Blue, table.PlayerOrder[3].TeamColor)
}

func TestPlayerPool_UpdateActivePlayer(t *testing.T) {
	table := NewPlayerPool()
	table.AddPlayer("playerOne")
	table.AddPlayer("playerTwo")
	table.AddPlayer("playerThree")
	table.AddPlayer("playerFour")
	table.SetPlayerOrder("playerOne", 0, Red)
	table.SetPlayerOrder("playerTwo", 1, Blue)
	table.SetPlayerOrder("playerThree", 2, Red)
	table.SetPlayerOrder("playerFour", 3, Blue)
	table.UpdateActivePlayer()
	assertStrEq(t, "playerOne", table.ActivePlayer.Player.UserId)
	table.UpdateActivePlayer()
	assertStrEq(t, "playerTwo", table.ActivePlayer.Player.UserId)
	table.UpdateActivePlayer()
	assertStrEq(t, "playerThree", table.ActivePlayer.Player.UserId)
	table.UpdateActivePlayer()
	assertStrEq(t, "playerFour", table.ActivePlayer.Player.UserId)
	table.UpdateActivePlayer()
	assertStrEq(t, "playerOne", table.ActivePlayer.Player.UserId)
}

func assertStrEq(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Expected: %s but %s", expected, actual)
	}
}

func assertUint8Eq(t *testing.T, expected uint8, actual uint8) {
	if actual != expected {
		t.Errorf("Expected: %d but %d", expected, actual)
	}
}