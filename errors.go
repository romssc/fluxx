package fluxx

import "errors"

var (
	ErrListening    = errors.New("[ERROR] fluxx: FAILED WHILE LISTENING")
	ErrShuttingDown = errors.New("[ERROR] fluxx: FAILED WHILE SHUTTING DOWN")
)
