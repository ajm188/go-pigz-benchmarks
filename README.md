# go-pigz-benchmarks

Benchmarks of Go gzip compression implementations.

## Usage

```
❯ make benchmark
```

## Results

On my laptop, as of 5/21:

```
❯ make benchmark
go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: ajm188.scratchpad/go-pigz-benchmarks
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkCompression/alice29.txt/pargzip-8           148           8153158 ns/op         2336456 B/op         40 allocs/op
BenchmarkCompression/alice29.txt/pgzip-8             363           2974656 ns/op         5860418 B/op         41 allocs/op
BenchmarkCompression/asyoulik.txt/pargzip-8                  165           7146528 ns/op         2287301 B/op         40 allocs/op
BenchmarkCompression/asyoulik.txt/pgzip-8                    434           2628958 ns/op         5860422 B/op         41 allocs/op
BenchmarkCompression/cp.html/pargzip-8                      1110            954799 ns/op         1936318 B/op         37 allocs/op
BenchmarkCompression/cp.html/pgzip-8                        1414            727058 ns/op         5860365 B/op         41 allocs/op
BenchmarkCompression/fields.c/pargzip-8                     1801            561185 ns/op         1896664 B/op         36 allocs/op
BenchmarkCompression/fields.c/pgzip-8                       2121            482658 ns/op         4746220 B/op         39 allocs/op
BenchmarkCompression/grammar.lsp/pargzip-8                  4020            302613 ns/op         1875408 B/op         35 allocs/op
BenchmarkCompression/grammar.lsp/pgzip-8                    2953            382537 ns/op         4746187 B/op         39 allocs/op
BenchmarkCompression/kennedy.xls/pargzip-8                    30          38550774 ns/op         4556435 B/op         42 allocs/op
BenchmarkCompression/kennedy.xls/pgzip-8                     133           8964113 ns/op         5860425 B/op         41 allocs/op
BenchmarkCompression/lcet10.txt/pargzip-8                     57          20085820 ns/op         3049246 B/op         42 allocs/op
BenchmarkCompression/lcet10.txt/pgzip-8                      182           6468213 ns/op         5860395 B/op         41 allocs/op
BenchmarkCompression/plrabn12.txt/pargzip-8                   37          28068500 ns/op         3458767 B/op         42 allocs/op
BenchmarkCompression/plrabn12.txt/pgzip-8                    146           8369082 ns/op         5860405 B/op         41 allocs/op
BenchmarkCompression/ptt5/pargzip-8                          148           7825150 ns/op         3057502 B/op         42 allocs/op
BenchmarkCompression/ptt5/pgzip-8                            331           3534940 ns/op         5860402 B/op         41 allocs/op
BenchmarkCompression/sum/pargzip-8                           666           1598065 ns/op         1984196 B/op         38 allocs/op
BenchmarkCompression/sum/pgzip-8                            1261            952427 ns/op         5860402 B/op         41 allocs/op
BenchmarkCompression/xargs.1/pargzip-8                      3602            423490 ns/op         1876948 B/op         35 allocs/op
BenchmarkCompression/xargs.1/pgzip-8                        2449            420577 ns/op         4746204 B/op         39 allocs/op
PASS
ok      ajm188.scratchpad/go-pigz-benchmarks       36.223s
```