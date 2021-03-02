package main

import (
	"flag"
	"fmt"

	"github.com/jpr98/solitaire/deck"
)

func main() {
	encode := flag.Bool("e", false, "set to encode a message")
	decode := flag.Bool("d", false, "set to decode a message")
	msg := flag.String("msg", "", "a message to encode/decode")

	d := deck.New()
	s := NewSolitaire(&d)

	flag.Parse()

	if *msg == "" {
		fmt.Println("You need to input a message to encode/decode")
		return
	}

	s.SetMessage(*msg)

	if *encode {
		fmt.Println(s.Encode())
	} else if *decode {
		fmt.Println(s.Decode())
	} else {
		fmt.Println("You need to set either encode or decode to true")
	}
}
