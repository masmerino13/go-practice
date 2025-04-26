package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 20 cards %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, but got %v", len(deck))
	}

	os.Remove(filename)
}
