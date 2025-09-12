package fluxx

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/romssc/fluxx/internal/utils"
)

type App struct {
	config config
	server *http.Server
}

func (a *App) Listen() error {
	utils.StartupMessage(a.config.version, a.config.author, false, a.config.address, a.config.readTimeout, a.config.writeTimeout)
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%v. %w", ErrListening, err)
	}
	return nil
}

func (a *App) ListenTLS(certificate string, key string) error {
	utils.StartupMessage(a.config.version, a.config.author, true, a.config.address, a.config.readTimeout, a.config.writeTimeout)
	if err := a.server.ListenAndServeTLS(certificate, key); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%v. %w", ErrListening, err)
	}
	return nil
}

func (a *App) GracefulShutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("%v. %w", ErrShuttingDown, err)
	}
	return nil
}

type config struct {
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	tls          *tls.Config
	mux          http.Handler

	version string
	author  string
}
