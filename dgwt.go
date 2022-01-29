package bet

import (
	"testing"
)

type Target[T any] struct {
	d  string
	fn func(t *testing.T)
}

type State[T any] struct {
	d  string
	fn func(t *testing.T, given func() T)
}

func Describe[T any](t *testing.T, name string, givens ...Target[T]) {
	if name == "" {
		for _, g := range givens {
			t.Run(g.d, g.fn)
		}
		return
	}

	t.Run("Describe "+name, func(t *testing.T) {
		for _, g := range givens {
			t.Run(g.d, g.fn)
		}
	})
}

func Given[T any](description string, given func() T, when ...State[T]) Target[T] {
	return Target[T]{
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

func When[T any](description string, when func(t *testing.T, state T) T, then []Assertion[T]) State[T] {
	return State[T]{
		d: description,
		fn: func(t *testing.T, given func() T) {
			for _, th := range then {
				g := given()
				w := when(t, g)
				t.Run("Then it "+th.name, func(tt *testing.T) {
					th.fn(tt, w)
				})
			}
		},
	}
}

type Assertion[T any] struct {
	name string
	fn   func(t *testing.T, v T)
}

func Then[T any](assertions ...Assertion[T]) []Assertion[T] { return assertions }

func It[T any](name string, fn func(t *testing.T, v T)) Assertion[T] {
	return Assertion[T]{
		name: name,
		fn:   fn,
	}
}
