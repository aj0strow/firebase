package firebase

import (
	"testing"
)

func TestNewPushID(t *testing.T) {
	outs := map[string]bool{}
	var order []string
	for i := 0; i < 1000; i++ {
		id := NewPushID()
		if outs[id] {
			t.Errorf("duplicate ID: %s", id)
		}
		outs[id] = true
		order = append(order, id)
	}
	for i := 1; i < len(order); i++ {
		if !(order[i] > order[i-1]) {
			t.Errorf("not k-sortable")
		}
	}
}
