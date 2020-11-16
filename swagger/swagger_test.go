package swagger

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_createHTML(t *testing.T) {
	type args struct {
		param *Parameter
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "All Parameter",
			args: args{
				param: &Parameter{
					Favicon: "favicon.ico",
					SpecURL: "http://example.com",
					Title:   "Test Document",
				},
			},
			want:    []byte(`<html><head><meta charset="UTF-8"/><title>Test Document</title><link rel="icon" href="favicon.ico"><style>body{margin:0;padding:0;}redoc{display:block;}</style><link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet"></head><body><redoc spec-url='http://example.com'></redoc><script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script></body></html>`),
			wantErr: false,
		},
		{
			name: "No Title",
			args: args{
				param: &Parameter{
					Favicon: "favicon.ico",
					SpecURL: "http://example.com",
					Title:   "",
				},
			},
			want:    []byte(`<html><head><meta charset="UTF-8"/><title>API Documentation</title><link rel="icon" href="favicon.ico"><style>body{margin:0;padding:0;}redoc{display:block;}</style><link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet"></head><body><redoc spec-url='http://example.com'></redoc><script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script></body></html>`),
			wantErr: false,
		},
		{
			name: "No Favicon",
			args: args{
				param: &Parameter{
					Favicon: "",
					SpecURL: "http://example.com",
					Title:   "",
				},
			},
			want:    []byte(`<html><head><meta charset="UTF-8"/><title>API Documentation</title><style>body{margin:0;padding:0;}redoc{display:block;}</style><link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet"></head><body><redoc spec-url='http://example.com'></redoc><script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script></body></html>`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createHTML(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("createHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func ExampleHTTPHandleFunc() {
	if err := http.ListenAndServe(":1202", HTTPHandleFunc("http://localhost:1324/docs/spec")); err != nil {
		panic(err)
	}
}

func ExampleEchoHandleFunc() {
	e := echo.New()
	e.GET("/docs", EchoHandleFunc("/docs/spec", WithTitle("Example")))
	if err := e.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
