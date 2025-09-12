package fluxx

import (
	"encoding/json"
	"net/http"
)

type Ctx struct {
	r Reader
	s Sender
}

func (c *Ctx) Read() *Reader {
	return &c.r
}

func (c *Ctx) Send() *Sender {
	return &c.s
}

type Reader struct {
	Request *http.Request
}

func (r *Reader) QueryParam(key string, defaultValue ...string) (string, bool) {
	if ok := r.Request.URL.Query().Has(key); !ok {
		if len(defaultValue) != 0 {
			return defaultValue[0], true
		}
		return "", false
	}
	return r.Request.URL.Query().Get(key), true
}

type Sender struct {
	Writer http.ResponseWriter

	r *http.Request
}

func (s *Sender) Error(status int, message string) {
	http.Error(s.Writer, message, status)
}

func (s *Sender) JSON(status int, data any, customHeaders ...map[string]string) error {
	if len(customHeaders) != 0 {
		for k, v := range customHeaders[0] {
			s.Writer.Header().Set(k, v)
		}
	}
	s.Writer.Header().Set("Content-Type", "application/json")
	s.Writer.WriteHeader(status)
	if err := json.NewEncoder(s.Writer).Encode(data); err != nil {
		return err
	}
	return nil
}

func (s *Sender) File(content string, filename string, path string, customHeaders ...map[string]string) {
	if len(customHeaders) != 0 {
		for k, v := range customHeaders[0] {
			s.Writer.Header().Set(k, v)
		}
	}
	s.Writer.Header().Set("Content-Type", content)
	s.Writer.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	http.ServeFile(s.Writer, s.r, path)
}
