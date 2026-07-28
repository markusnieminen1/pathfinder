package visualising

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pathfinder/data"
	"time"
)

func InitWeb(ctx context.Context, start, end *data.Station, fast_path []data.Station, turns [][]string) {

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("GET /", Roothandler(start, end, &fast_path, turns))
	http.HandleFunc("GET /events", EventsHandler)
	http.HandleFunc("GET /style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "visualising/style.css")
	})

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(shutdownCtx)
	}()

	fmt.Println("Visualisation available at: http://localhost:8080 \nUse ^C to shutdown.")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
