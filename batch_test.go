package firebase

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBatch(t *testing.T) {
	tests := []struct {
		getBatch func() *Batch
		update   interface{}
	}{
		{
			getBatch: func() *Batch {
				b := NewBatch()
				b.Set(Reference{"foo", "bar"}, 1)
				b.Set(Reference{"foo", "baz"}, 1)
				return b
			},
			update: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": 1,
					"baz": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		batch := tt.getBatch()
		out, err := batch.Merge()
		if err != nil {
			t.Fatal(err)
		}
		j1, err := json.Marshal(out)
		if err != nil {
			t.Fatal(err)
		}
		j2, err := json.Marshal(tt.update)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(j1, j2) {
			t.Errorf("broken")
		}
	}
}
