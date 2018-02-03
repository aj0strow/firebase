package firebase

import (
	"errors"
)

var ErrInvalidBatch = errors.New("firebase: batch update has empty or conflicting references")

func NewBatch() *Batch {
	return &Batch{}
}

type Batch struct {
	keys   []Reference
	values []interface{}
}

func (b *Batch) Set(key Reference, value interface{}) {
	b.keys = append(b.keys, key)
	b.values = append(b.values, value)
}

func (b *Batch) Merge() (interface{}, error) {
	out := map[string]interface{}{}
	for i := range b.keys {
		err := setDeep(out, b.keys[i], b.values[i])
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func setDeep(m map[string]interface{}, ref Reference, value interface{}) error {
	if len(ref) == 0 {
		return ErrInvalidBatch
	}
	n, ok := m[ref[0]]
	if !ok {
		if len(ref) == 1 {
			m[ref[0]] = value
			return nil
		}
		q := map[string]interface{}{}
		m[ref[0]] = q
		return setDeep(q, ref[1:], value)
	}
	q, ok := n.(map[string]interface{})
	if !ok {
		return ErrInvalidBatch
	}
	return setDeep(q, ref[1:], value)

}
