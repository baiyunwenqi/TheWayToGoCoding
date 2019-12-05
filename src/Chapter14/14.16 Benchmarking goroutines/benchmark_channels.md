# Benchmark channels

Here we apply it to a concrete example of a goroutine which is filled with ints, and then read. The
functions are called N times (e.g. N = 1000000) with testing.Benchmark, the BenchMarkResult
has a String() method for outputting its findings. The number N is decided upon by gotest,
judging this to be high enough to get a reasonable benchmark result.

```go
func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch { // wait for go func() ending

	}
}
```