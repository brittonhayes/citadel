// Code generated by goa v3.2.6, DO NOT EDIT.
//
// vulnerabilities go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen citadel/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountFindHandler configures the mux to serve the "vulnerabilities" service
// "find" endpoint.
func MountFindHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/vulnerabilities/{id}", f)
}

// MountListHandler configures the mux to serve the "vulnerabilities" service
// "list" endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/vulnerabilities", f)
}

// MountSubmitHandler configures the mux to serve the "vulnerabilities" service
// "submit" endpoint.
func MountSubmitHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/vulnerabilities", f)
}
