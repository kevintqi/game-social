package game_social

import "net/url"

type User struct {
	UserId string
	Alias string
	Icon url.URL
	Scores uint64
}

