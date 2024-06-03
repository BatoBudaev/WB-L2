package main

import (
	"calendar/internal"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create_event", internal.CreateEvent)
	http.HandleFunc("/update_event", internal.UpdateEvent)
	http.HandleFunc("/delete_event", internal.DeleteEvent)
	http.HandleFunc("/events_for_day", internal.GetEventsForDay)
	http.HandleFunc("/events_for_week", internal.GetEventsForWeek)
	http.HandleFunc("/events_for_month", internal.GetEventsForMonth)

	loggedRouter := internal.LoggingMiddleware(http.DefaultServeMux)

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(port, loggedRouter))
}
