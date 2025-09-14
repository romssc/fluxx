⚠️ ***Status:** package is under development — APIs may change.*

<img src="https://github.com/user-attachments/assets/d5447791-c42c-421b-9bde-fdbbdad407f1" alt="logo" width="200"/>

**Fluxx** is a lightweight, opinionated HTTP library for **Go**. It provides a thin abstraction over the **Go** **[net/http](https://pkg.go.dev/net/http@go1.25.1)** package, with utilities for request reading, response sending, and graceful server lifecycle management.

Designed to stay close to the **Go** **[standard library](https://pkg.go.dev/std)**, but with just enough ergonomics to make your web server development simpler!

## 🕷️ Quick Start

This guide will follow you through your first steps to use this package.

### Get it

```bash
go get -u github.com/romssc/fluxx
```

### Use it

Getting started with **Fluxx** is easy. Here's a basic example to create a simple web server that responds with "hello!" on the root path. This example demonstrates initializing a new **Fluxx** **[App](#typeApp)**, setting up a route, and starting the server.

```go
func main() {
    mux := http.NewServeMux()
	mux.HandleFunc("/", fluxx.HandlerFuncAdapter(func(c *fluxx.Ctx) {
        c.Send().JSON(http.StatusOK, "hello!")
	}))

	app := fluxx.New(fluxx.Config{
		Address:      ":8080",            // or "localhost:8080" just like in the standart net/http
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Mux: http.NewServeMux(),          // in this example it's the standart mux, might as well be chi or any other.
	})

	if err := app.Listen(); err != nil {
		panic(err)
	}
}
```

### Run it

```bash
go run main.go
```

### Check out the fancy startup message

This message will also contain some of your **[Config](#typeConfig)** settings, just for convenience.

<img src="https://github.com/user-attachments/assets/7594c989-fc15-4b5a-b1f2-e3d7f9b68252" alt="screenshot" width="300"/>

## 🧬 Middleware

Here is a list of middleware that are included within the **Fluxx**.

| Title | Description |
| :------- | :------- |
| [rid](https://github.com/romssc/fluxx/tree/9f657007e7c081c3823537b2513ae345f9f593b1/middleware/rid) | Allow to track Request ID, creates one if there's none. |
| [timeout](https://github.com/romssc/fluxx/tree/9f657007e7c081c3823537b2513ae345f9f593b1/middleware/timeout) | Sets a time to handle the request. |
| [key](https://github.com/romssc/fluxx/tree/9f657007e7c081c3823537b2513ae345f9f593b1/middleware/key) | Simple API Key validator. |

## 📄 Documentation <a name="index"></a>

[Index](#index)

[Variables](#variables)

[Functions](#functions)
- [func New(c Config) *App](#funcNew)
- [func HandlerFuncAdapter(h FluxxHandlerFunc) http.HandlerFunc](#funcHandlerFuncAdapter)

[Types](#types)
- [type App](#typeApp)
    - [func (a *App) Listen() error](#appFuncListen)
    - [func (a *App) ListenTLS(certificate, key string) error](#appFuncListenTLS)
    - [func (a *App) GracefulShutdown(timeout time.Duration) error](#appFuncGracefulShutdown)
- [type Config](#typeConfig)
- [type Ctx](#typeCtx)
    - [func (c *Ctx) Read() *Reader](#ctxFuncRead)
    - [func (c *Ctx) Send() *Sender](#ctxFuncSend)
- [type Reader](#typeReader)
    - [func (r *Reader) QueryParam(key string, defaultValue …string) (string, bool)](#readerFuncQueryParam)
- [type Sender](#typeSender)
    - [func (s *Sender) Error(status int, message string)](#senderFuncError)
    - [func (s *Sender) JSON(status int, data any, customHeaders …map[string]string) error](#senderFuncJSON)
    - [func (s *Sender) File(content, filename, path string, customHeaders …map[string]string)](#senderFuncFile)
- [type FluxxHandlerFunc](#typeFluxxHandlerFunc)

### Variables <a name="variables"></a>

```go
ErrListening    = errors.New("[ERROR] fluxx: FAILED WHILE LISTENING")
ErrShuttingDown = errors.New("[ERROR] fluxx: FAILED WHILE SHUTTING DOWN")
```

*Error values returned by server lifecycle operations.*

### Functions <a name="functions"></a>

#### func New <a name="funcNew"></a>

```go
func New(c Config) *App
```

#### func HandlerFuncAdapter <a name="funcHandlerFuncAdapter"></a>

```go
func HandlerFuncAdapter(h FluxxHandlerFunc) http.HandlerFunc
```

***HandlerFuncAdapter** adapts a **FluxxHandlerFunc** into a standard **http.HandlerFunc**.*

### Types <a name="types"></a>

#### type App <a name="typeApp"></a>

```go
type App struct {
    // contains filtered or unexported fields
}
```

***App** represents the **Fluxx** HTTP application. It manages server configuration and lifecycle.*

#### func Listen <a name="appFuncListen">

```go
func (a *App) Listen() error
```

***Listen** starts the HTTP server. It blocks until the server shuts down or fails.*

#### func ListenTLS <a name="appFuncListenTLS"></a>

```go
func (a *App) ListenTLS(certificate, key string) error
```

***ListenTLS** starts the server with TLS enabled.*

#### func GracefulShutdown <a name="appFuncGracefulShutdown"></a>

```go
func (a *App) GracefulShutdown(timeout time.Duration) error
```

***GracefulShutdown** attempts to shut down the server gracefully within the given timeout.*

#### type Config <a name="typeConfig"></a>

```go
type Config struct {
    Address      string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
    TLS          *tls.Config
    Mux          http.Handler
}
```

***Config** holds the configuration options for creating a new **App**.*

#### type Ctx <a name="typeCtx"></a>

```go
type Ctx struct {
    // contains filtered or unexported fields
}
```

***Ctx** wraps the HTTP request and response, providing ergonomic helpers for reading inputs and sending outputs.*

#### func Read <a name="ctxFuncRead"></a>

```go
func (c *Ctx) Read() *Reader
```

***Read** returns the request reader.*

#### func Send <a name="ctxFuncSend"></a>

```go
func (c *Ctx) Send() *Sender
```

***Send** returns the response sender.*

#### type Reader <a name="typeReader"></a>

```go
type Reader struct {
    Request *http.Request
}
```

***Reader** provides methods for accessing request data.*

#### func QueryParam <a name="readerFuncQueryParam"></a>

```go
func (r *Reader) QueryParam(key string, defaultValue ...string) (string, bool)
```

***QueryParam** returns the query parameter by key. If not present, an optional default value may be used. The second return value indicates whether a value was found or defaulted.*

#### type Sender <a name="typeSender"></a>

```go
type Sender struct {
    Writer http.ResponseWriter
    // contains filtered or unexported fields
}
```

***Sender** provides methods for writing responses.*

#### func Error <a name="senderFuncError"></a>

```go
func (s *Sender) Error(status int, message string)
```

***Error** sends an HTTP error response with the given status code and message.*

#### func JSON <a name="senderFuncJSON"></a>

```go
func (s *Sender) JSON(status int, data any, customHeaders ...map[string]string) error
```

***JSON** sends a JSON response with the given status code, data, and optional headers.*

#### func File <a name="senderFuncFile"></a>

```go
func (s *Sender) File(content, filename, path string, customHeaders ...map[string]string)
```

***File** serves a file response with content type and disposition headers.*

#### type FluxxHandlerFunc <a name="typeFluxxHandlerFunc"></a>

```go
type FluxxHandlerFunc func(c *Ctx)
```

***FluxxHandlerFunc** defines a handler that operates on a **Fluxx** context.*

## 🏴󠁩󠁤󠁳󠁭󠁿 License

[MIT](https://opensource.org/license/mit)





