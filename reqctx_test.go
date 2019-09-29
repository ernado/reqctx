package reqctx

import (
	"context"
	"net/http"
	"testing"
)

func newReq() *http.Request {
	req, err := http.NewRequest(http.MethodGet, "https://example.org", nil)
	if err != nil {
		panic(err)
	}
	return req
}

func TestSet(t *testing.T) {
	req := newReq()
	newCtx := context.WithValue(context.Background(), "foo", "bar")
	Set(req, newCtx)
	if v := req.Context().Value("foo").(string); v != "bar" {
		t.Errorf("unexpected value %s", v)
	}
}

func TestSetValue(t *testing.T) {
	req := newReq()
	SetValue(req, "foo", "bar")
	if v := req.Context().Value("foo").(string); v != "bar" {
		t.Errorf("unexpected value %s", v)
	}
}

func BenchmarkSet(b *testing.B) {
	b.ReportAllocs()
	req := newReq()
	newCtx := context.WithValue(context.Background(), "foo", "bar")

	for i := 0; i < b.N; i++ {
		Set(req, newCtx)
	}
}

func BenchmarkWithContext(b *testing.B) {
	b.ReportAllocs()
	req := newReq()
	newCtx := context.WithValue(context.Background(), "foo", "bar")

	for i := 0; i < b.N; i++ {
		req.WithContext(newCtx)
	}
}
