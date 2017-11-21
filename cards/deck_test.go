package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deckFile := "_decktesting"

	os.Remove(deckFile)

	d := newDeck()
	d.saveToFile(deckFile)

	loadedDeck := newDeckFromFile(deckFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, instead got %v", len(loadedDeck))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	shuffledDeck := newDeck()

	shuffledDeck.shuffle()

	if d[0] == shuffledDeck[0] {
		t.Errorf("Expected deck to be shuffled instead cards are the same, %v, %v.", d[0], shuffledDeck[0])
	}

	if d[len(d)-1] == shuffledDeck[len(d)-1] {
		t.Errorf("Expected deck to be shuffled instead cards are the same, %v, %v.", d[len(d)-1], shuffledDeck[len(d)-1])
	}

	if len(d) != len(shuffledDeck) {
		t.Errorf("Expected length to be the same, instead got %v and %v", len(d), len(shuffledDeck))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	auxDeck := newDeck()
	handSize := len(d) / 2

	firstHand, secondHand := deal(d, handSize)

	if len(firstHand) != len(secondHand) {
		t.Errorf("Expected length of half deck to be the same, instead got %v and %v", len(firstHand), len(secondHand))
	}

	if auxDeck[:handSize][0] != firstHand[0] {
		t.Errorf("Expected cards to be the same, instead got %v and %v", auxDeck[:handSize][0], firstHand[0])
	}

	if auxDeck[handSize:][0] != secondHand[0] {
		t.Errorf("Expected cards to be the same, instead got %v and %v", auxDeck[handSize:][0], secondHand[0])
	}
}

func TestToString(t *testing.T) {
	d := newDeck()

	if strings.Join([]string(d), ",") != d.toString() {
		t.Errorf("Expected strings to be equal, instead got %v and %v", strings.Join([]string(d), ","), d.toString())
	}
}

func TestSaveToFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	d := newDeck()

	d.saveToFile(fileName)

	bs, err := ioutil.ReadFile(fileName)

	loadedDeck := deck(strings.Split(string(bs), ","))

	if err != nil {
		t.Errorf("Error ocurred during TestSaveToFile: %v", err)
	}

	if loadedDeck == nil {
		t.Error("Expected to find a deck, but got nil")
	}

	if loadedDeck[0] != d[0] {
		t.Errorf("Expected cards to be equal, instead got %v and %v", loadedDeck[0], d[0])
	}
}
