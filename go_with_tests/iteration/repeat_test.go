package iteration

import "testing"

func TestIteration(t *testing.T) {

	t.Run("Will repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		expect := "aaaaa"

		if got != expect {
			t.Errorf("expected: %q, got: %q", got, expect)
		}
	})

	t.Run("Will repeat 8 times", func(t *testing.T) {
		got := Repeat("b", 8)
		expect := "bbbbbbbb"
		if got != expect {
			t.Errorf("expected: %q, got: %q", got, expect)
		}
	})

	// how did we do multiple tests?

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 5)
	}
}
