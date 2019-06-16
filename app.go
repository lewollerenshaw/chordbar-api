package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers / endpoints
	r.HandleFunc("/api/generate-progression", serveChordProgression).Methods("GET")
	r.HandleFunc("/api/generate-modalscale", serveModalScale).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func serveChordProgression(w http.ResponseWriter, r *http.Request) {
	chosenMode := 4
	chosenKey := 4
	progressionLength := 4

	modalScaleChords := generateModalChords(chosenMode)
	modalScaleSteps := generateModalSteps(chosenMode)
	modalScaleNotes := generateModalNotes(chosenKey, modalScaleSteps)
	modalScale := mapModalScale(modalScaleNotes, modalScaleChords)
	chordProgression := generateChordProgression(modalScale, progressionLength)

	println(chordProgression)
}

func serveModalScale(w http.ResponseWriter, r *http.Request) {

}
