# Chapter 11

## 11. Testing

Testing, by which we implicitly mean *automated* testing, is the practice of writing small programs that check that the code under test (the *production* code) behaves as expected for certain inputs,...

### 11.1 The go test Tool

Within *_test.go files, three kinds of functions are treated specially: **tests**, **benchmarks**, and **examples**.

### 11. 4 Benchmarking

A benchmark function look like a test function, but with the *Benchmark* prefix and a **testing.B* parameter.

The fastest program is often the one that makes the fewest memory allocations.
