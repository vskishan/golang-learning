package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//Create a deck type
type deck []string

//Function to print the cards
func (d deck) print(){
	
	fmt.Println("Total cards : ", len(d))

	for i,cards := range d {
		fmt.Println(i, cards)
	}
}

//Function to create the deck of cards
func createDeck() deck {
	cards := deck{}

	cardSuites := []string { "Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string { "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _,suite := range cardSuites {
		for _,value := range cardValues {
			cards = append(cards , value + " of " + suite)
		}
	}

	return cards
}

//Function to create a hand of cards
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Function that returns a string from deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Function to write cards to a file
func (d deck) writeToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//Function to read cards from a file
func readFromFile(fileName string) deck {
	deckByte, err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Print("Error in reading the file : ", err)
		os.Exit(1)
	}
	return strings.Split(string(deckByte), ",")
}

//Function to shuffle the cards
func (d deck) shuffle() {
	for i := range d {
		randomNumber := rand.Intn(51)
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}