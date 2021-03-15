package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght 16,but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[0])
	}
}

func TestSaceToDeckAndNewDeckFromFile(t *testing.T) {

	tempFileName := "_decktesting"
	os.Remove(tempFileName)

	deck := newDeck()
	deck.saveToFile(tempFileName)

	loadedDeck := newDeckFromFile(tempFileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght 16,but got %d", len(loadedDeck))
	}

	os.Remove(tempFileName)
}
