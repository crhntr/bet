# Behavior Tests

This is an exploration in writing structured Go tests using type parameters.

## Specs

```golang
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
```

## Given / When / Then

```golang
package bet_test

import (
	"testing"

	"github.com/crhntr/bet"
)

type Set map[string]struct{}

func (s Set) Add(v string) {
	s[v] = struct{}{}
}

func (s Set) Contains(v string) bool {
	_, ok := s[v]
	return ok
}

func (s Set) Length() int {
	return len(s)
}

func TestSet(t *testing.T) {
	bet.Describe[Set](t, "",
		bet.Given("an empty set", func() Set {
			return make(Set)
		},
			bet.When("a value is added", func(t *testing.T, set Set) Set {
				t.Cleanup(func() {
					// do some clean up here.
				})
				set.Add("hello")
				return set
			},
				bet.Then(
					bet.It("has the field", func(t *testing.T, set Set) {
						if !set.Contains("hello") {
							t.Fail()
						}
					}),
					bet.It("does not have some other key", hasOtherKey),
					bet.It("has a length of 1", lengthIsNot1),
				),
			),
		),
	)
}

func hasOtherKey(t *testing.T, set Set) {
	if set.Contains("greetings") {
		t.Fail()
	}
}
func lengthIsNot1(t *testing.T, set Set) {
	if set.Length() != 1 {
		t.Fail()
	}
}

```

