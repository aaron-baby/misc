package algorithms

import "testing"

func TestFNV0(t *testing.T) {
	fnv0 := NewFNV0()
	s := `chongo <Landon Curt Noll> /\../\`
	fnv0.Write([]byte(s))
	sum := int(fnv0.Sum32())
	expected := 2166136261

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
