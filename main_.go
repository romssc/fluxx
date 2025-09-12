package fluxx

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Config struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	TLS *tls.Config

	Mux http.Handler
}

func New(c Config) *App {
	return &App{
		config: config{
			address:      c.Address,
			readTimeout:  c.ReadTimeout,
			writeTimeout: c.WriteTimeout,
			tls:          c.TLS,
			mux:          c.Mux,
			version:      "v0.1.0",
			author:       "https://github.com/romssc",
		},
		server: &http.Server{
			Addr:         c.Address,
			Handler:      c.Mux,
			TLSConfig:    c.TLS,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
		},
	}
}

type FluxxHandlerFunc func(c *Ctx)

func HandlerFuncAdapter(h FluxxHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &Ctx{
			r: Reader{
				Request: r,
			},
			s: Sender{
				Writer: w,
				r:      r,
			},
		}
		h(c)
	}
}
