package web

import (
	"fmt"
	"net/http"
	"time"
)

// Server ...
func Server(port int) {

	addr := fmt.Sprintf(`localhost:%d`, port)

	mux := http.NewServeMux()
	mux.HandleFunc(`/json`, fallback)
	mux.HandleFunc(`/pb`, fallback)

	s := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Println(`start web server`, addr)

	s.ListenAndServe()
}

func fallback(w http.ResponseWriter, r *http.Request) {
}
