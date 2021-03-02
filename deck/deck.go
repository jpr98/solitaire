package deck

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Deck represents the basic holding structure for cards
type Deck struct {
	Cards  []int
	jokerA int
	jokerB int
}

// New creates a new deck
func New() Deck {
	return Deck{
		Cards:  readCards(),
		jokerA: 27,
		jokerB: 28,
	}
}

func readCards() []int {
	data, err := ioutil.ReadFile("./deck.txt")
	if err != nil {
		log.Fatalf("Error reading file %s", err)
		return nil
	}
	var ints []int
	stringData := string(data)
	err = json.Unmarshal([]byte(stringData), &ints)
	if err != nil {
		log.Fatalf("Error parsing file %s", err)
	}
	return ints
}

func (d *Deck) swap(i, j int) {
	var l, h int
	if i < j {
		l, h = i, j
	} else {
		l, h = j, i
	}

	if l < 0 {
		l = len(d.Cards) - 1
	}

	if h >= len(d.Cards) {
		h = 0
	}

	d.Cards[l], d.Cards[h] = d.Cards[h], d.Cards[l]
}

func (d *Deck) swapRight(i, times int) {
	for times > 0 {
		for i < len(d.Cards) && times > 0 {
			d.swap(i, i+1)
			i++
			times--
		}
		i = 0
	}
}

func (d *Deck) swapLeft(i, times int) {
	for times > 0 {
		for i >= 0 && times > 0 {
			d.swap(i, i-1)
			i--
			times--
		}
		i = len(d.Cards) - 1
	}
}

/*
tripleCut takes 3 parts of the array and shifts them.
Everything before the first index goes to the back.
Everything after the second index goes to the front.
Everything in between stays the same.
*/
func (d *Deck) tripleCut(i, j int) {
	var first, second int
	if i < j {
		first, second = i, j
	} else {
		first, second = j, i
	}

	var newDeck []int
	newDeck = append(newDeck, d.Cards[second+1:len(d.Cards)]...)
	newDeck = append(newDeck, d.Cards[first:second+1]...)
	newDeck = append(newDeck, d.Cards[0:first]...)
	copy(d.Cards, newDeck)
}

func (d *Deck) moveBackCards(n int) {
	var newDeck []int
	newDeck = append(newDeck, d.Cards[n:len(d.Cards)-1]...)
	newDeck = append(newDeck, d.Cards[0:n]...)
	newDeck = append(newDeck, d.Cards[len(d.Cards)-1])
	copy(d.Cards, newDeck)
}

// IndexOf finds de index of a given value in a given array
func indexOf(value int, array []int) int {
	for index, val := range array {
		if val == value {
			return index
		}
	}
	return -1
}

// Step1 finds Joker A and swaps it with the value following it
func (d *Deck) Step1() {
	indexJokerA := indexOf(d.jokerA, d.Cards)
	d.swap(indexJokerA, indexJokerA+1)
}

// Step2 finds Joker B and moves it twice down the list
func (d *Deck) Step2() {
	indexJokerB := indexOf(d.jokerB, d.Cards)
	d.swapRight(indexJokerB, 2)
}

// Step3 does the triple cut
func (d *Deck) Step3() {
	indexJokerA, indexJokerB := indexOf(d.jokerA, d.Cards), indexOf(d.jokerB, d.Cards)
	d.tripleCut(indexJokerA, indexJokerB)
}

// Step4 gets the last card and brings the number of cards that its value indicates in front of it
func (d *Deck) Step4() {
	lastCardValue := d.Cards[len(d.Cards)-1] - 1
	d.moveBackCards(lastCardValue)
}

// Step5 returns the keystream for the current configuartion of Cards in Deck
func (d *Deck) Step5() int {
	topValue := d.Cards[0] - 1
	return d.Cards[topValue]
}

func (d *Deck) doSteps1To4() {
	d.Step1()
	d.Step2()
	d.Step3()
	d.Step4()
}

// GetKeystreamValue gets a single keystream value which is not a joker
func (d *Deck) GetKeystreamValue() int {
	d.doSteps1To4()
	keystream := d.Step5()
	for valueInArray(keystream, []int{d.jokerA, d.jokerB}) {
		d.doSteps1To4()
		keystream = d.Step5()
	}
	return keystream
}

func valueInArray(val int, array []int) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}
