package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/kg0r0/authzen-interop-pdp-example/pdp"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/access/v1/evaluation", pdp.Evaluation)
	r.HandleFunc("/access/v1/evaluations", pdp.Evaluation)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Not found", "path", r.URL.Path)
		http.Error(w, "Not found", http.StatusNotFound)
	})

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	slog.Info("Starting server", "port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
