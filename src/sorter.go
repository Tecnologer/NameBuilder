package randname

import "sort"

type letterValue struct {
	letter      string
	probability float32
}

type letterSorter struct {
	letters []*letterValue
	by      By
}

func byProbilityDesc(l, l2 *letterValue) bool {
	return l.probability > l2.probability
}

// By is the type of a "less" function that defines the ordering of its LetterProbability arguments.
type By func(p1, p2 *letterValue) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(letters []*letterValue) {
	ps := &letterSorter{
		letters: letters,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *letterSorter) Len() int {
	return len(s.letters)
}

// Swap is part of sort.Interface.
func (s *letterSorter) Swap(i, j int) {
	s.letters[i], s.letters[j] = s.letters[j], s.letters[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *letterSorter) Less(i, j int) bool {
	return s.by(s.letters[i], s.letters[j])
}
