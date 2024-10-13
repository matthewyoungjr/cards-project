package main

import "testing"


func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(deck))
	}
}