package timeout

import (
	"context"
	"net/http"
	"time"
)

type timeoutWriter struct {
	http.ResponseWriter
	timedOut chan struct{}
}

func (tw *timeoutWriter) Write(b []byte) (int, error) {
	select {
	case <-tw.timedOut:
		return 0, nil
	default:
		return tw.ResponseWriter.Write(b)
	}
}

func (tw *timeoutWriter) WriteHeader(statusCode int) {
	select {
	case <-tw.timedOut:
		return
	default:
		tw.ResponseWriter.WriteHeader(statusCode)
	}
}

func New(timeout time.Duration, errs string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()
			tw := &timeoutWriter{ResponseWriter: w, timedOut: make(chan struct{})}
			r = r.WithContext(ctx)
			done := make(chan struct{})
			go func() {
				next.ServeHTTP(tw, r)
				close(done)
			}()
			select {
			case <-ctx.Done():
				close(tw.timedOut)
				http.Error(w, errs, http.StatusGatewayTimeout)
			case <-done:
			}
		})
	}
}
