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
		wont = 3
		get  = SeriesSumWithMutex(2)
	)
	if wont != get {
		t.Errorf("sum of the series: get %d, wont %d", get, wont)
	}
}

func TestSeriesSumWithChan(t *testing.T) {
	var (
		wont = 3
		get  = SeriesSumWithChan(2)
	)
	if wont != get {
		t.Errorf("sum of the series: get %d, wont %d", get, wont)
	}
}
