package mutexvschannels

import "testing"

func Benchmark_seriesSumWithMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		seriesSumWithMutex(2)
	}
}

func Benchmark_seriesSumWithChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		seriesSumWithChan(2)
	}
}

func Test_seriesSumWithMutex(t *testing.T) {
	var (
		wont = 3
		get  = seriesSumWithMutex(2)
	)
	if wont != get {
		t.Errorf("sum of the series: get %d, wont %d", get, wont)
	}
}

func Test_seriesSumWithChan(t *testing.T) {
	var (
		wont = 3
		get  = seriesSumWithChan(2)
	)
	if wont != get {
		t.Errorf("sum of the series: get %d, wont %d", get, wont)
	}
}
