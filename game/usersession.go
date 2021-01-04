package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

///Stores the sesion for a single user
type UserSession struct {
	Username string
	score    int
}

func NewUserSession(userName string) *UserSession {
	session := &UserSession{
		Username: userName,
	}

	return session
}

/*type FoxPawResult struct {
	Roll1 int
	Roll2 int
	Roll3 int
}*/

func (f *UserSession) Roll() int {
	value := rand.Intn(6) + 1

	return value //need random 1-6
}

func (f *UserSession) executeUserSession() {
	reader := bufio.NewReader(os.Stdin)
	rolledFoxPaw := false

	fmt.Printf("Welcome player %v, Press enter to roll...\n", f.Username)
	reader.ReadString('\n')

	for {
		roll1 := f.Roll()
		roll2 := f.Roll()
		roll3 := f.Roll()

		if roll1 == 6 || roll2 == 6 || roll3 == 6 {
			rolledFoxPaw = true
		}

		if rolledFoxPaw {
			fmt.Println("You rolled a FoxPaw!")
			f.score = 0
			return
		}

		fmt.Printf("You rolled %v, %v, %v, for a total of %v\n Do you want to roll again?\n", roll1, roll2, roll3, f.score+roll1+roll2+roll3)

		f.score += roll1 + roll2 + roll3

		ans, err := reader.ReadString('\n')
		if err != nil {
			panic("Keyboard Error")
		}

		ans = strings.TrimSpace(ans)

		if ans != "y" && ans != "Y" {
			fmt.Println("you have chosen to stop at ", f.score)
			break
		}
	}
}
