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
