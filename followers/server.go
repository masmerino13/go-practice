package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	PostByUser(1)
}

type User struct {
	ID    int
	Name  string
	Likes []string
}

type Post struct {
	Author    int
	Timestamp time.Time
	Text      string
}

type Follow struct {
	SourceID int
	TargetID int
}

var Users = []User{
	{1, "Chepe", []string{"cheese", "bacon"}},
	{2, "Jose", []string{"cheese", "bacon"}},
	{3, "Mario", []string{"cheese", "bacon"}},
}

var Follows = []Follow{
	{1, 2},
	{1, 3},
	{3, 1},
	{2, 3},
}

var Posts = []Post{
	Post{1, time.Date(2013, 5, 9, 8, 22, 0, 0, time.UTC), "Prof Jamyn Edis, founder @dash_labs featured on PBSs innovation panel"},
	Post{2, time.Date(2018, 2, 28, 19, 43, 0, 0, time.UTC), "Ever wonder how I spend my days? Heres an in-depth look at how I spent a few of them!"},
	Post{3, time.Date(2018, 3, 18, 10, 19, 0, 0, time.UTC), "Humans are too good at shooting basketballs. They are ruining the game."},
	Post{1, time.Date(2013, 5, 7, 15, 6, 0, 0, time.UTC), ".@wee_Spring CEO Ally Downey from @TechStarsNYC talks"},
}

func PostByUser(uid int) ([]Post, error) {
	followed := make(map[int]bool, 0)

	for _, u := range Follows {
		if u.SourceID == uid {
			followed[u.TargetID] = true
		}
	}

	feed := make([]Post, 0)

	for _, p := range Posts {
		if followed[p.Author] {
			feed = append(feed, p)
		}
	}

	fmt.Println(feed)

	return nil, errors.New("failed")
}
