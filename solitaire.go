package main

import (
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/jpr98/solitaire/deck"
)

// Solitaire contains the logic for encoding and decoding messages with the Solitaire algorithm
type Solitaire struct {
	deck       *deck.Deck
	message    string
	numMessage []int
	keystreams []int
}

// NewSolitaire creates a new solitarie object
func NewSolitaire(deck *deck.Deck) Solitaire {
	return Solitaire{deck: deck}
}

// SetMessage sanitizes the input and sets it to message
func (s *Solitaire) SetMessage(msg string) {
	cleanMsg := sanitize(msg)
	upperMsg := strings.ToUpper(cleanMsg)
	paddedMsg := padUntilMultipleOf5(upperMsg)
	s.message = paddedMsg
}

// Encode encodes the previously set message
func (s *Solitaire) Encode() string {
	s.msgToNum()
	s.generateKeystreams()

	var encodedMessage string
	for i, v := range s.numMessage {
		sum := v + s.keystreams[i]
		num := numToBase(sum)
		letter := numToLetter(num)
		encodedMessage += string(letter)
	}
	return encodedMessage
}

// Decode decodes the previously set message
func (s *Solitaire) Decode() string {
	s.msgToNum()
	s.generateKeystreams()

	var decodedMessage string
	for i, v := range s.numMessage {
		top := v
		if top <= s.keystreams[i] {
			top += 26
		}
		sub := top - s.keystreams[i]
		sub = numToBase(sub)
		letter := numToLetter(sub)
		decodedMessage += string(letter)
	}
	return decodedMessage
}

func (s *Solitaire) generateKeystreams() {
	s.keystreams = nil
	for range s.numMessage {
		s.keystreams = append(s.keystreams, s.deck.GetKeystreamValue())
	}
}

func (s *Solitaire) msgToNum() {
	s.numMessage = nil
	for _, v := range s.message {
		num := letterToNum(byte(v))
		s.numMessage = append(s.numMessage, num)
	}
}

func numToBase(num int) int {
	modulus := num % 26
	if modulus == 0 {
		return 26
	}
	return modulus
}

func numToLetter(num int) byte {
	return byte('A' + num - 1)
}

func letterToNum(c byte) int {
	return int(c) - int('A') + 1
}

func sanitize(msg string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(msg, "")
}

func padUntilMultipleOf5(msg string) string {
	for len(msg)%5 != 0 {
		msg = strPad(msg, len(msg)+1, "X", "RIGHT")
	}
	return msg
}

func strPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}
