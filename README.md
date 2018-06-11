# Go mutex vs channels example

[![build status](https://travis-ci.org/danil/go_mutex_vs_channels_example.svg?branch=master)](https://travis-ci.org/danil/go_mutex_vs_channels_example)

## Benchmarks

    go test -bench=. github.com/danil/go_mutex_vs_channels_example/...
    Benchmark_seriesSumWithMutex-4       1000000          1369 ns/op
    Benchmark_seriesSumWithChan-4        1000000          1609 ns/op

## License

Licensed under [MIT License](./LICENSE)
