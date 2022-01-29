package bet_test

import (
	"testing"

	"github.com/crhntr/bet"
)

func TestFloat64(t *testing.T) {
	behavior := bet.New[float64](t)
	defer behavior.Run()

	behavior.Setup(func(t *testing.T) float64 {
		return 420
	})

	behavior.Spec("is high", func(t *testing.T, n float64) {
		if n != 420 {
			t.Fail()
		}
	})

	behavior.Spec("is not zero", func(t *testing.T, n float64) {
		if n == 0 {
			t.Fail()
		}
	})
}
