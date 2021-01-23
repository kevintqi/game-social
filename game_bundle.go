package game_social

type GameBundle struct {
	Id string
	GameName string
	Caller string
	StartTime uint64
	Duration uint64
	Game Game
	Posts[] Post
	Players[] Player
}

