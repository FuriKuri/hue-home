package main

import (
	"gbbr.io/hue"
	"log"
)

func main() {
	b, err := hue.Discover()
	if err != nil {
		log.Fatal(err)
	}
	if !b.IsPaired() {
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			log.Fatal(err)
		}
	}
	light, err := b.Lights().Get("Hue white lamp 1")
	if err != nil {
		log.Fatal(err)
	}
	if err := light.Off(); err != nil {
		log.Fatal(err)
	}
}
