package examples

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
