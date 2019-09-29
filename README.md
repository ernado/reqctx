# reqctx

Package reqctx provides a way to set http request context without
creating a shallow copy.

**Please don't use in production!**

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

So, instead of doing that: 
```go
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "foo", "bar")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

You can now do this: 
```go
package reqctx_test

import (
	"context"
	"net/http"
	
	"github.com/ernado/reqctx"
)

func MiddlewareFast(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqctx.SetValue(r, "foo", "bar")
		// Same as:
		// ctx := context.WithValue(r.Context(), "foo", "bar")
		// reqctx.Set(r, ctx)
		next.ServeHTTP(w, r)
	})
}
```


## Benchmarks

```
goos: linux
goarch: amd64
BenchmarkSet-12          1000000000  0.532 ns/op  0 B/op    0 allocs/op
BenchmarkWithContext-12  26447012    45.8 ns/op   128 B/op  1 allocs/op
PASS
```