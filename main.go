package main

import (
	"fmt"
	"strings"
	"wordle-cit/word_generator"

	"github.com/go-playground/validator/v10"
)

func main() {

	g := new(Game)
	g.Init()
	fmt.Println("To play Wordle, enter a 5 letter word and then press Enter!")
	for g.Continue() {
		guess := TakeInput()
		g.PlayTurn(guess)
		g.Print()
	}
}

var validate = validator.New()

func TakeInput() string {
	for {
		var input string
		fmt.Scanln(&input)
		errs := validate.Var(input, "required,len=5,alpha")
		if errs != nil {
			fmt.Println(errs.Error())
		} else {
			return input
		}
	}
}

type Game struct {
	won     bool
	live    bool
	word    string
	guesses []string
}

func (b *Game) Init() {
	b.word = word_generator.Get()
	b.live = true
	b.won = false
}

func (b *Game) Print() {
	fmt.Println("-----")
	for i, guess := range b.guesses {
		fmt.Println(i+1, guess)
	}
	fmt.Println("-----")
}

func (b *Game) Continue() bool {
	return b.live
}

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func (g *Game) PlayTurn(guess string) {
	// populate output text
	outputTxt := ""
	for i, char := range guess {
		correctChar := g.word[i]
		if correctChar == byte(char) {
			outputTxt += Green + string(char) + Green + Reset
		} else if strings.Contains(g.word, string(char)) {
			outputTxt += Yellow + string(char) + Yellow + Reset
		} else {
			outputTxt += Reset + string(char) + Reset + Reset
		}
	}
	g.guesses = append(g.guesses, outputTxt)

	// lets see if we won
	if guess == g.word {
		g.won = true
		g.live = false
	}

	// Check status of game
	if len(g.guesses) == 5 {
		g.live = false
	}

}
