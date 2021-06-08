package main

import (
	"log"

	"github.com/cat-in-vacuum/rand-play/generator"
)

const count = 10

func main() {
	gen := generator.Init(count)

	for i := 0; i < count+1; i++ {
		next, err := gen.GetNextQuestion()
		if err != nil {
			log.Printf("[error] get next question errored: %s\n", err)
			return
		}

		log.Printf("[info] next: %d; available: %+v\n", next, gen.Available())
	}
}
