package main

import "github.com/BrandonLaseter/foxpaw_go/game"

func main() {
	game := game.NewGame(2)
	game.Play()
}

/*func main() {
	players := 1
	foxpaw := game.New(players)
	const rollsPerPlayer = 3
	foxpawRolled := false

	results := [rollsPerPlayer]int{}
	total := 0

	for roll := 0; roll < rollsPerPlayer; roll++ {
		results[roll] = foxpaw.Roll()

		if results[roll] != 6 {
			foxpawRolled = true
			fmt.Println("you rolled a foxpaw :(")
		} else {
			fmt.Println("You rolled ", results[roll])
		}

		total += results[roll]
	}

	if !foxpawRolled {
		fmt.Println("You rolled a FoxPaw total score is ", total)
	} else {
		total = 0
		fmt.Println("Total score is ", total)
	}
}*/
