// Package reqctx provides a totally unsafe, but fast way to set request
// context.
package reqctx

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"unsafe"
)

// httpRequest is copied version of http.Request structure.
type httpRequest struct {
	Method           string
	Proto            string // "HTTP/1.0"
	URL              *url.URL
	ProtoMajor       int // 1
	ProtoMinor       int // 0
	Header           http.Header
	Body             io.ReadCloser
	GetBody          func() (io.ReadCloser, error)
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	Trailer          http.Header
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
	Cancel           <-chan struct{}
	Response         *http.Response
	ctx              context.Context
}

// Set sets request context without shallow copy of request.
func Set(req *http.Request, ctx context.Context) {
	p := (*httpRequest)(unsafe.Pointer(req))
	p.ctx = ctx
}

// SetValue wraps context.WithValue call on request context.
func SetValue(req *http.Request, k, v interface{}) {
	ctx := context.WithValue(req.Context(), k, v)
	Set(req, ctx)
}

func init() {
	// Explicitly check that structures have at least equal size.
	stdSize := unsafe.Sizeof(http.Request{})
	gotSize := unsafe.Sizeof(httpRequest{})
	if stdSize != gotSize {
		panic(fmt.Errorf("%d (net/http) != %d", stdSize, gotSize))
	}
}
