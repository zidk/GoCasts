package main

import (
	"os"
	"testing"
)

const MaxDeckLength = 16
const TestFileName = "_decktesting"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	checkDeckLenght(t, d)
	checkDeckValue(t, 0, d, "Ace of Spades")
	checkDeckValue(t, len(d)-1, d, "Four of Clubs")
}

func TestSaveToFile(t *testing.T) {
	d := newDeck()
	deleteSavedDeck()
	err := d.saveToFile(TestFileName)
	if err != nil {
		t.Errorf("Got error loading deck from file, err: %v", err)
	}
}

func TestLoadFromFile(t *testing.T) {
	d := loadFromFile(TestFileName)

	checkDeckLenght(t, d)
	deleteSavedDeck()
}

func deleteSavedDeck() {
	os.Remove(TestFileName)
}

func checkDeckLenght(t *testing.T, d deck) {
	l := len(d)

	if l != MaxDeckLength {
		t.Errorf("Expected deck lenght of %v but got %v", MaxDeckLength, l)
	}
}

func checkDeckValue(t *testing.T, p int, d deck, cn string) {
	if d[p] != cn {
		t.Errorf("Expected name is %v but got %v", cn, d[p])
	}
}
