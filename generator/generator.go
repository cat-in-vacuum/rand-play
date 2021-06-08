package generator

import (
	"fmt"
	"math/rand"
	"time"
)

type Gen struct {
	available []int
}

func Init(count int) Gen {
	rand.Seed(time.Now().Unix())
	return Gen{
		available: rand.Perm(count),
	}
}

func (g *Gen) GetNextQuestion() (int, error) {
	if g.available == nil {
		return 0, fmt.Errorf("must init gen first")
	}

	if len(g.available) == 0 {
		return 0, fmt.Errorf("end of avalible elems")
	}

	out := g.available[0]
	g.available = g.available[1:]

	return out, nil
}

func (g *Gen) Available() []int {
	return g.available
}
