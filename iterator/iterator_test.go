package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	personlist := PersonList{}
	names := []string{"ray", "rey", "ruy", "roy"}
	for _, name := range names {
		personlist.Persons = append(personlist.Persons, Person{Name: name})
	}

	index := 0
	for iterator := personlist.Iterator(); iterator.HasNext(); {
		p := iterator.Next()
		if names[index] != p.Name {
			t.Errorf("Expect %s, found %s", names[index], p.Name)
		}
		index++
	}
}
