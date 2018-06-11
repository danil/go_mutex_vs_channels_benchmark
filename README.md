# Go mutex vs channels benchmark

[![build status](https://travis-ci.org/danil/go_mutex_vs_channels_benchmark.svg?branch=master)](https://travis-ci.org/danil/go_mutex_vs_channels_benchmark)

## Benchmark

    go test -bench=. github.com/danil/go_mutex_vs_channels_benchmark/...
    Benchmark_seriesSumWithMutex-4       1000000          1369 ns/op
    Benchmark_seriesSumWithChan-4        1000000          1609 ns/op

## License

Licensed under [MIT License](./LICENSE)
