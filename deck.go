package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func createCard() deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Diamond", "Club", "Heart"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit:= range cardSuit {
		for _, value := range values {
			cards=append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int)  (deck, deck){
	return d[:handSize], d[handSize:]
}

// First converting the deck into slice of string, 
// And joining the all string seperated by ","
func (d deck) toString() string {
	return  strings.Join([]string(d), ",")
}


// Converting the slice of deck --> string ---> slice of byte,
// if file doesn't exist, it creates a new file with that name and edits the file with byte data.
// SAVING DATA TO HARD DRIVE.
func(d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {


	// readfile returns two value of type []byte and error,
	content, err := os.ReadFile(filename)
	if err!=nil {
		fmt.Println("Error: ", err)

		os.Exit(1)
	}

	// We are converting []byte to string, 
	// spliting the string by "," and returns the slice of string
	s := strings.Split(string(content), ",")
	return deck(s)
}




func (d deck) shuffle(){
	for i := range d {
		newPosition := rand.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition] ,d[i]
	}
}