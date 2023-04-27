package keeperrorstatus

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Config struct {
	Header string `test`
}

func CreateConfig() *Config {
	return &Config{
		Header: "TEMPLATEHEADER",
	}
}

type Keeperrorstatus struct {
	next     http.Handler
	header   string
	name     string
	template *template.Template
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Header == "TEMPLATEHEADER" {
		return nil, fmt.Errorf("header needs to be set!")
	}

	return &Keeperrorstatus{
		header:   config.Header,
		next:     next,
		name:     name,
		template: template.New("Keeperrorstatus").Delims("[[", "]]"),
	}, nil
}

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
