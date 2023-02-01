package examples_test

import (
	"github.com/crhntr/bet/examples"
	"testing"

	"github.com/crhntr/bet"
)

func TestSet(t *testing.T) {
	bet.Describe[SetBehavior](t,
		bet.Given("an empty set", NewEmptySet,
			bet.When("a value is added", SetBehavior.aValueIsAdded,
				bet.Then(
					bet.It("has the field", func(scope SetBehavior, t *testing.T) {
						if !scope.testTarget.Contains("hello") {
							t.Fail()
						}
					}),
					bet.It("does not have some other key", SetBehavior.hasOtherKey),
					bet.It("has a length of 1", SetBehavior.lengthIsNot1),
				),
			),
			bet.When("a value is added twice", func(state SetBehavior, t *testing.T) SetBehavior {
				state.testTarget.Add("banana")
				state.testTarget.Add("banana")
				return state
			},
				bet.Then(
					bet.It("has the field", func(scope SetBehavior, t *testing.T) {
						if !scope.testTarget.Contains("banana") {
							t.Fail()
						}
					}),
					bet.It("has a length of 1", SetBehavior.lengthIsNot1),
				),
			),
		),
	)
}

type SetBehavior struct {
	testTarget examples.Set
}

func NewEmptySet(*testing.T) SetBehavior {
	return SetBehavior{
		testTarget: make(examples.Set),
	}
}

func (scope SetBehavior) aValueIsAdded(t *testing.T) SetBehavior {
	t.Cleanup(func() {
		// do some clean up here.
	})
	scope.testTarget.Add("hello")
	return scope
}

func (scope SetBehavior) hasOtherKey(t *testing.T) {
	if scope.testTarget.Contains("greetings") {
		t.Fail()
	}
}

func (scope SetBehavior) lengthIsNot1(t *testing.T) {
	if scope.testTarget.Length() != 1 {
		t.Fail()
	}
}
