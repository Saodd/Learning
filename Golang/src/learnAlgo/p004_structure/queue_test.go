package p004_structure

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMyPriorityQueueInt_Push(t *testing.T) {
	q := MyPriorityQueueInt{}
	for i := 0; i < 10; i++ {
		q.Push(rand.Int())
	}
	// checkOrder
	intold, e := q.Pop()
	if e != nil {
		t.Error("Pop failed: ", fmt.Sprint(e))
	}
	for i := 0; i < 9; i++ {
		intnew, e := q.Pop()
		if e != nil {
			t.Error("Pop failed: ", fmt.Sprint(e))
		}
		if intnew > intold {
			t.Error("Order wrong: ", intold, intnew)
		}
		intold = intnew
	}
}
