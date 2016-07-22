package data_structures

import (
	"testing"
)

func TestContainsElementsAddedToLinkedList(t *testing.T) {
	l := new(List);

	l.Add(1)
	l.Add(2)
	l.Add(3)

	for i := 1; i <= 3; i++ {
		if l.Search(i) == nil {
			t.Errorf("Should have contained a node with value %d", i)
		}
	}

	if l.size != 3 {
		t.Errorf("The incorrect size is %d", l.size)
	}

}

