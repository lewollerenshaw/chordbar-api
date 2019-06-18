package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers / endpoints
	r.HandleFunc("/api/progression", getChordProgression).Methods("GET")
	r.HandleFunc("/api/scale", getModalScale).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Get chord progressions
func getChordProgression(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Progression:  Hit by request")
	w.Header().Set("Content-Type", "application/json")

	var progression Progression

	// Decode request body and stores it in progression
	_ = json.NewDecoder(r.Body).Decode(&progression)

	// Converts request body to integers
	key, err := strconv.Atoi(progression.Key)
	mode, err := strconv.Atoi(progression.Mode)
	progressionLength, err := strconv.Atoi(progression.ProgressionLength)

	// Null check
	if err != nil {
		log.Fatal("Handling null check")
	}

	// Encode chord progression and return json object
	json.NewEncoder(w).Encode(serveChordProgression(key, mode, progressionLength))
}

// Get modal scales
func getModalScale(w http.ResponseWriter, r *http.Request) {
	log.Println("Get modal scale: Hit by request")
	w.Header().Set("Content-Type", "application/json")

	var scale Scale

	_ = json.NewDecoder(r.Body).Decode(&scale)

	// Converts request body to integers
	key, err := strconv.Atoi(scale.Key)
	mode, err := strconv.Atoi(scale.Mode)

	// Null check
	if err != nil {
		log.Fatal("Handling null check")
	}

	// Encode scale and return json object
	json.NewEncoder(w).Encode(serveModalScale(key, mode))
}
