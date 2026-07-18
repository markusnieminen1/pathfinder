package visualising

import (
	"context"
	"log"
	"net/http"
	"pathfinder/data"
	"time"
)

func InitWeb(ctx context.Context, start, end *data.Station) {

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("GET /", Roothandler(start, end))
	http.HandleFunc("GET /events", EventsHandler)

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(shutdownCtx)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
