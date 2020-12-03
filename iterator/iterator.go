package iterator

// Person person
type Person struct {
	Name string
}

// PersonList PersonList
type PersonList struct {
	Persons []Person
}

// Iterator return iterator
func (pl *PersonList) Iterator() PersonListIterator {
	return PersonListIterator{PersonList: pl}
}

// PersonListIterator PersonListIterator
type PersonListIterator struct {
	PersonList *PersonList
	index      int
}

// HasNext 次があるか
func (pi *PersonListIterator) HasNext() bool {
	if pi.index >= len(pi.PersonList.Persons) {
		return false
	}
	return true
}

// Next 次の要素
func (pi *PersonListIterator) Next() Person {
	person := pi.PersonList.Persons[pi.index]
	pi.index++
	return person
}
