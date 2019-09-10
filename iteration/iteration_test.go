package iteration

import "testing"

func TestReapeat(t *testing.T) {
	got := Repeat("a", 8)
	want := "aaaaaaaa"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
