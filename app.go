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
	r.HandleFunc("/api/test", testConnection).Methods("GET")
	r.HandleFunc("/api/progression", getChordProgression).Methods("GET")
	r.HandleFunc("/api/scale", getModalScale).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Test connection to API is available
// TODO: Remove this function when development finished
func testConnection(w http.ResponseWriter, r *http.Request) {
	chord := Chord{
		Note:  "3",
		Chord: "3",
	}
	log.Println("Test: Hit by request")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chord)
}

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

func getModalScale(w http.ResponseWriter, r *http.Request) {

}
