package mutexvschannels

import "testing"

func BenchmarkSeriesSumWithMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SeriesSumWithMutex(2)
	}
}

func BenchmarkSeriesSumWithChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SeriesSumWithChan(2)
	}
}

func TestSeriesSumWithMutex(t *testing.T) {
	var (
		want = 3
		got  = SeriesSumWithMutex(2)
	)
	if want != got {
		t.Errorf("sum of the series: got %d, want %d", got, want)
	}
}

func TestSeriesSumWithChan(t *testing.T) {
	var (
		want = 3
		got  = SeriesSumWithChan(2)
	)
	if want != got {
		t.Errorf("sum of the series: got %d, want %d", got, want)
	}
}
