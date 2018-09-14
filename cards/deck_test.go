package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	deckLength := len(deck)

	if deckLength != 16 {
		t.Errorf("Expected deck length of 16, but got %v", deckLength)
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}

	if deck[deckLength-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", deck[deckLength-1])
	}

}

func TestDeal(t *testing.T) {
	cards := newDeck()
	handLength := 5
	hand, leftCardsInDeck := deal(cards, handLength)

	if len(hand) != handLength {
		t.Errorf("Expected hand length of %v but received %v", handLength, len(hand))
	}

	if len(leftCardsInDeck) != (len(cards) - handLength) {
		t.Errorf("Expected deck length of %v but received %v", (len(cards) - handLength), len(leftCardsInDeck))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFilePath := "_decktesting"
	os.Remove(testFilePath)

	deck := newDeck()
	deck.saveToFile(testFilePath)

	loadedDeck := newDeckFromFile(testFilePath)
	deckLength := len(loadedDeck)

	if deckLength != 16 {
		t.Errorf("Expected deck length of 16, but got %v", deckLength)
	}

	os.Remove(testFilePath)
}
