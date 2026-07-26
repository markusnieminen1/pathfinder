package visualising

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pathfinder/data"
	"time"
)

func InitWeb(ctx context.Context, start, end *data.Station, fastest_route *[]data.Station) {

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("GET /", Roothandler(start, end, fastest_route))
	http.HandleFunc("GET /events", EventsHandler)

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(shutdownCtx)
	}()

	fmt.Println("Visualisation visibile at: http://localhost:8080 \nUse ^C to shutdown.")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
