package game

import (
	"fmt"
	"math/rand"
	"time"
)

type FoxPaw struct {
	Scores      [2]int
	PlayerCount int //hard coded at 2 currently
	//...
}

func NewGame(players int) *FoxPaw {
	rand.Seed(time.Now().Unix())

	podkin := &FoxPaw{
		PlayerCount: players,
	}

	return podkin
}

func (f *FoxPaw) Play() {
	fmt.Println("FoxPaw time!")

	for player := 0; player < f.PlayerCount; player++ {
		playerName := fmt.Sprint("Player", player+1)
		fmt.Println("Starting Player ", playerName)
		sess := NewUserSession(playerName)
		sess.executeUserSession()

		f.Scores[player] = sess.score
		fmt.Println()
	}

	f.DisplayResults()
}

func (f *FoxPaw) DisplayResults() {
	for i := 0; i < len(f.Scores); i++ {
		fmt.Printf("Player score is %v\n", f.Scores[i])
	}
}
