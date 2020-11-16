package swagger

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

const html = `<html><head><meta charset="UTF-8"/>{{if .Title}}<title>{{.Title}}</title>{{else}}<title>API Documentation</title>{{end}}{{if .Favicon}}<link rel="icon" href="{{.Favicon}}">{{end}}<style>body{margin:0;padding:0;}redoc{display:block;}</style><link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet"></head><body><redoc spec-url='{{.SpecURL}}'></redoc><script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script></body></html>`

// HTTPHandleFunc returns http.HandlerFunc
func HTTPHandleFunc(specURL string, opts ...ParameterOption) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		p := &Parameter{SpecURL: specURL}
		for _, opt := range opts {
			opt(p)
		}
		buf, err := createHTML(p)
		if err != nil {
			res.Write([]byte("unexpected error"))
			return
		}
		res.Header().Set(echo.HeaderContentType, "text/html")
		res.Write(buf)
	}
}

// EchoHandleFunc returns echo.HandleFunc
func EchoHandleFunc(specURL string, opts ...ParameterOption) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := &Parameter{SpecURL: specURL}
		for _, opt := range opts {
			opt(p)
		}
		buf, err := createHTML(p)
		if err != nil {
			return fmt.Errorf("failed to create html: %v", err)
		}
		return c.HTMLBlob(
			http.StatusOK,
			buf,
		)
	}
}

func createHTML(param *Parameter) ([]byte, error) {
	tmpl, err := template.New("swagger").Parse(html)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, param); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ParameterOption sets Parameter fields
type ParameterOption func(*Parameter)

// WithFavicon sets favicon url to Parameter.Favicon
func WithFavicon(favicon string) ParameterOption {
	return func(p *Parameter) {
		p.Favicon = favicon
	}
}

// WithTitle sets title to Parameter.Title
func WithTitle(title string) ParameterOption {
	return func(p *Parameter) {
		p.Title = title
	}
}

// Parameter is golang template parameter
type Parameter struct {
	Title   string `json:"title,omitempty"`
	Favicon string `json:"favicon,omitempty"`
	SpecURL string `json:"spec_url,omitempty"`
}
