// +build !test

package mongodb

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/mainflux/mainflux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MakeHandler returns a HTTP handler for API endpoints.
func MakeHandler() http.Handler {
	r := bone.New()
	r.GetFunc("/version", mainflux.Version("mongodb-writer"))
	r.Handle("/metrics", promhttp.Handler())

	return r
}
