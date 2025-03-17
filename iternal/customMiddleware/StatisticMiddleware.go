package custommiddleware

import (
	"net/http"

	"github.com/dixxe/personal-website/iternal"
)

func Statistics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iternal.VisitorAmount++
		h.ServeHTTP(w, r)
	})
}
