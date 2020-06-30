package middleware

import (
	"net/http"
	"strconv"

	"github.com/eminetto/clean-architecture-go-v2/pkg/metric"

	"github.com/codegangsta/negroni"
)

//Metrics to prometheus
func Metrics(mService metric.Service) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		appMetric := metric.NewHTTP(r.URL.Path, r.Method)
		appMetric.Started()
		next(w, r)
		res := w.(negroni.ResponseWriter)
		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(res.Status())
		mService.SaveHTTP(appMetric)
	}
}
