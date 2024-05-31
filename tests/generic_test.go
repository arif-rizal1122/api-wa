package tests

import "testing"

func TestGeneric(t *testing.T) {
	// Test implementation here
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	t.Log(SumIntsOrFloats(ints))
}

func SumIntsOrFloats[K comparable, V int | float32 | int32 | int64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
