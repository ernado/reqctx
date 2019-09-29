# reqctx

Package reqctx provides a way to set http request context without
creating a shallow copy.

Current implementation of WithContext is following:
```go
func (r *Request) WithContext(ctx context.Context) *Request {
	if ctx == nil {
		panic("nil context")
	}
	r2 := new(Request)
	*r2 = *r
	r2.ctx = ctx
	r2.URL = cloneURL(r.URL) // legacy behavior; TODO: try to remove. Issue 23544
	return r2
}
```

The `Set` function is effectively doing the following:
```go
func Set(r *Request, ctx context.Context) {
	r.ctx = ctx
}
```

## Benchmarks

```
goos: linux
goarch: amd64
BenchmarkSet-12                 1000000000               0.532 ns/op           0 B/op          0 allocs/op
BenchmarkWithContext-12         26447012                45.8 ns/op           128 B/op          1 allocs/op
PASS
```