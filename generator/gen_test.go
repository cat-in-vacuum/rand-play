package generator_test

import (
	"testing"

	"github.com/cat-in-vacuum/rand-play/generator"
	"github.com/stretchr/testify/require"
)

func TestGen(main *testing.T) {
	main.Run("the numbers do not repeat", func(t *testing.T) {
		guardMap := make(map[int]struct{})

		count := 10

		gen := generator.Init(count)

		for i := 0; i < 10; i++ {
			next, err := gen.GetNextQuestion()
			require.NoError(t, err)
			if _, ok := guardMap[next]; ok {
				t.Fatalf("value %d already exist", next)
			}

			guardMap[next] = struct{}{}
		}

		return
	})

	main.Run("available values are exhausted", func(t *testing.T) {

		count := 10

		gen := generator.Init(count)

		for i := 0; i < 11; i++ {
			_, err := gen.GetNextQuestion()

			if i == 10 {
				require.EqualError(t, err, "end of avalible elems")
			}

		}

	})
}
