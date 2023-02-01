package bet

import (
	"testing"
)

type InitialCondition[T any] struct {
	d  string
	fn func(t *testing.T)
}

type State[T any] struct {
	d  string
	fn func(t *testing.T, given func(t *testing.T) T)
}

func Describe[T any](t *testing.T, givens ...InitialCondition[T]) {
	for _, g := range givens {
		t.Run(g.d, g.fn)
	}
}

func Given[T any](description string, given func(t *testing.T) T, when ...State[T]) InitialCondition[T] {
	return InitialCondition[T]{
		d: "Given " + description,
		fn: func(gt *testing.T) {
			for _, w := range when {
				gt.Run("When "+w.d, func(wt *testing.T) {
					w.fn(wt, given)
				})
			}
		},
	}
}

func When[T any](description string, when func(state T, t *testing.T) T, then []Assertion[T]) State[T] {
	return State[T]{
		d: description,
		fn: func(t *testing.T, given func(t *testing.T) T) {
			for _, th := range then {
				g := given(t)
				w := when(g, t)
				t.Run("Then it "+th.name, func(tt *testing.T) {
					th.fn(w, tt)
				})
			}
		},
	}
}

type Assertion[T any] struct {
	name string
	fn   func(v T, t *testing.T)
}

func Then[T any](assertions ...Assertion[T]) []Assertion[T] { return assertions }

func It[T any](name string, fn func(v T, t *testing.T)) Assertion[T] {
	return Assertion[T]{
		name: name,
		fn:   fn,
	}
}
