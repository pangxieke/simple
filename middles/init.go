package middles

import (
	"github.com/pangxieke/simple/log"
	"net/http"
	"time"
)

func Handler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info(map[string]interface{}{
			"access_time": time.Now(),
			"ip":          r.RemoteAddr,
			"method":      r.Method,
			"path":        r.RequestURI,
			"query":       r.URL.RawQuery,
		})
		next.ServeHTTP(w, r)
	})
}
