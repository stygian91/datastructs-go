package queue_test

import (
	"testing"

	"github.com/stygian91/datastructs-go/queue"
)

func TestPriorityQueue(t *testing.T) {
	q := queue.NewPriorityQueue[string]()
	q.Add(3, "a")
	q.Add(1, "b")
	q.Add(2, "c")

	expected := []string{"b", "c", "a"}
	for _, e := range expected {
		a, err := q.Remove()
		if err != nil {
			t.Errorf("TestPriorityQueue(): %s", err)
			return
		}

		if a != e {
			t.Errorf("TestPriorityQueue(): expected %s, found %s", e, a)
			return
		}
	}

	if !q.Empty() {
		t.Error("TestPriorityQueue(): expected queue to be empty after removing all items")
	}
}
