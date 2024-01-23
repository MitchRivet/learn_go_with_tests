package iteration

import "testing"

func TestIteration(t *testing.T) {

	got := Repeat("a")
	expect := "aaaaa"

	if got != expect {
		t.Errorf("expected: %q, got: %q", got, expect)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b")
	}
}
