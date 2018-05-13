# Chapter 2

## Exercise 2.3

There are other implementations for PopCount: _BitCount_, _PopCountByClearing_,
_PopCountByShifting_; see in `popcount_test`.

Explain why `popcount2` is slower:

    ❯ go test -bench=.
    goos: darwin
    goarch: amd64
    pkg: github.com/chrisduong/gopl.io/ch2/popcount2
    BenchmarkPopCount-4             	2000000000	         0.31 ns/op
    BenchmarkPopCount2-4            	100000000	        17.9 ns/op
    BenchmarkBitCount-4             	2000000000	         0.31 ns/op
    BenchmarkPopCountByClearing-4   	50000000	        30.4 ns/op
    BenchmarkPopCountByShifting-4   	20000000	        87.4 ns/op
    PASS
    ok  	github.com/chrisduong/gopl.io/ch2/popcount2	6.523s

    * Recap *
    The PopCount2 using loop will be slower because it need extra variable (initialization time) and every iteration also take time.
