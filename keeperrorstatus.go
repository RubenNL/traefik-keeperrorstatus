// Package keeperrorstatus is my package.
package keeperrorstatus

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Config is my config.
type Config struct {
	Header string `json:"header,omitempty"`
}

// CreateConfig function.
func CreateConfig() *Config {
	return &Config{
		Header: "TEMPLATEHEADER",
	}
}

// Keeperrorstatus type.
type Keeperrorstatus struct {
	next     http.Handler
	header   string
	name     string
	template *template.Template
}

// New function.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Header == "TEMPLATEHEADER" {
		return nil, fmt.Errorf("Header needs to be set")
	}

	return &Keeperrorstatus{
		header:   config.Header,
		next:     next,
		name:     name,
		template: template.New("Keeperrorstatus").Delims("[[", "]]"),
	}, nil
}

// ServeHTTP function.
func (a *Keeperrorstatus) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("test")
	if header != "" {
		statusCode, err := strconv.Atoi(header)
		if err == nil {
			rw.WriteHeader(statusCode)
		}
	}
	a.next.ServeHTTP(rw, req)
}
