package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("__decktesting")

	deck := newDeck()
	deck.saveToFile("__decktesting")

	loadedDeck := newDeckFromFile("__decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(loadedDeck))
	}

	os.Remove("__decktesting")
}
