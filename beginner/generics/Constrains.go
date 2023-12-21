package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
CONSTRAINTS
Sometimes you need the logic in your generic function to know something about the types it operates on.
The example we used in the first exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

Constraints are just interfaces that allow us to write generics that only operate within the constraint of a given interface type.
In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.

CREATING A CUSTOM CONSTRAINT
Let's take a look at the example of a concat function.
It takes a slice of values and concatenates the values into a string.
This should work with any type that can represent itself as a string, even if it's not a string under the hood.
For example, a user struct can have a .String() that returns a string with the user's name and age.
*/
// https://www.digitalocean.com/community/tutorials/how-to-use-generics-in-go#restricting-generic-types

type PlayingCard struct {
	suit string
	rank string
}

type TradingCard struct {
	name string
}

func NewPlayingCard(suit string, rank string) *PlayingCard {
	return &PlayingCard{
		suit: suit,
		rank: rank,
	}
}

func NewTradingCard(name string) *TradingCard {
	return &TradingCard{name: name}
}

func (tc TradingCard) String() string {
	return tc.name
}

func (pc PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.rank, pc.suit)
}

type Deck[T any] struct {
	cards []T
}

func (d *Deck[T]) addCard(card T) {
	d.cards = append(d.cards, card)
}

func (d *Deck[T]) randomCard() T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	idx := r.Intn(len(d.cards))
	return d.cards[idx]
}

func NewPlayingCardDesk() *Deck[*PlayingCard] {
	suits := []string{"Diamond", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck[*PlayingCard]{}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck.addCard(NewPlayingCard(suit, rank))
		}
	}

	return deck
}

func NewTradingCardDesk() *Deck[*TradingCard] {
	cardNames := []string{"Sammy", "Droplets", "Space", "App Platform"}

	deck := &Deck[*TradingCard]{}
	for _, nm := range cardNames {
		deck.addCard(NewTradingCard(nm))
	}

	return deck
}

func main() {
	playingCardDeck := NewPlayingCardDesk()
	tradingCardDesk := NewTradingCardDesk()

	fmt.Println("---drawing playing card---")
	playingCard := playingCardDeck.randomCard()

	fmt.Println(playingCard.String())

	fmt.Println("---drawing trading card---")
	tradingCard := tradingCardDesk.randomCard()
	fmt.Println(tradingCard.String())
}
