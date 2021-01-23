package game_social

const (
	Blue uint8 = 1
	Red  uint8 = 0
)

type Seat struct {
	Player *Player
	Next *Seat
}

type Player struct {
	UserId    string
	Seat      uint8
	Cards     []Card
	TeamColor uint8
}

type PlayerPool struct {
	Host           *Player
	ActivePlayer   *Seat
	PlayerOrder    []*Player
	Players        map[string]*Player
}

func NewPlayerPool() *PlayerPool {
	players := make(map[string]*Player)
	return &PlayerPool{Players: players}
}

func (p *PlayerPool) AddPlayer(playerId string) {
	p.Players[playerId] = &Player{UserId: playerId}
}

func (p *PlayerPool) SetHost(playerId string) {
	player := p.Players[playerId]
	p.Host = player
	if p.ActivePlayer == nil {
		p.ActivePlayer = buildSeatLink(p.PlayerOrder)
	}
	ap := p.ActivePlayer
	for _,_ = range p.PlayerOrder {
		if ap.Player == p.Host {
			break
		} else {
			ap = ap.Next
		}
	}
	p.ActivePlayer = ap
}

func (p *PlayerPool) SetPlayerOrder(playerId string, order uint8, team uint8) {
	if p.PlayerOrder == nil {
		p.PlayerOrder = make([]*Player, len(p.Players))
	}
	player := p.Players[playerId]
	player.Seat = order
	player.TeamColor = team
	p.PlayerOrder[order] = player
}

func (p *PlayerPool) UpdateActivePlayer() {
	p.ActivePlayer = p.ActivePlayer.Next
}

func buildSeatLink(players []*Player) *Seat {
	seat := &Seat{}
	first := seat
	last := seat
	for _, p := range players {
		seat.Player = p
		seat.Next = &Seat{}
		last = seat
		seat = seat.Next
	}
	last.Next = first
	return first
}
