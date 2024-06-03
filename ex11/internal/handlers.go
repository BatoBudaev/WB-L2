package internal

import (
	"net/http"
	"strconv"
	"time"
)

var events = make(map[int]Event)
var eventID = 1

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if err := parseForm(r); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid form data"})
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid date"})
		return
	}

	event := Event{
		ID:          eventID,
		UserID:      userID,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Date:        date,
	}

	events[eventID] = event
	eventID++

	writeJSON(w, http.StatusOK, SuccessResponse{Result: "Event created"})
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if err := parseForm(r); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid form data"})
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid event ID"})
		return
	}

	event, exists := events[id]
	if !exists {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "Event not found"})
		return
	}

	if userID := r.FormValue("user_id"); userID != "" {
		event.UserID, err = strconv.Atoi(userID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
			return
		}
	}

	if title := r.FormValue("title"); title != "" {
		event.Title = title
	}

	if description := r.FormValue("description"); description != "" {
		event.Description = description
	}

	if date := r.FormValue("date"); date != "" {
		event.Date, err = time.Parse("2006-01-02", date)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid date"})
			return
		}
	}

	events[id] = event
	writeJSON(w, http.StatusOK, SuccessResponse{Result: "Event updated"})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if err := parseForm(r); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid form data"})
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid event ID"})
		return
	}

	if _, exists := events[id]; !exists {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "Event not found"})
		return
	}

	delete(events, id)
	writeJSON(w, http.StatusOK, SuccessResponse{Result: "Event deleted"})
}

func GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid date"})
		return
	}

	var result []Event
	for _, event := range events {
		if event.Date.Equal(date) {
			result = append(result, event)
		}
	}

	writeJSON(w, http.StatusOK, result)
}

func GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid date"})
		return
	}

	year, week := date.ISOWeek()
	var result []Event
	for _, event := range events {
		eventYear, eventWeek := event.Date.ISOWeek()
		if eventYear == year && eventWeek == week {
			result = append(result, event)
		}
	}

	writeJSON(w, http.StatusOK, result)
}

func GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid date"})
		return
	}

	year, month, _ := date.Date()
	var result []Event
	for _, event := range events {
		eventYear, eventMonth, _ := event.Date.Date()
		if eventYear == year && eventMonth == month {
			result = append(result, event)
		}
	}

	writeJSON(w, http.StatusOK, result)
}
