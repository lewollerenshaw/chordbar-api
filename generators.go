package main

import (
	"math/rand"
	"time"
)

/* Contains generators for modal scales and progressions */

// Generates the chords for each degree in the scale
func generateModalChords(chosenMode int) map[int]string {
	if chosenMode == 1 {
		return majorScaleChordMap
	}

	var newModalScale = map[int]string{}
	var degree = 1

	for i := chosenMode; i <= len(majorScaleChordMap); i++ {
		newModalScale[degree] = majorScaleChordMap[i]
		degree++
	}

	for i := 1; i < chosenMode; i++ {
		newModalScale[degree] = majorScaleChordMap[i]
		degree++
	}

	return newModalScale
}

// Generates the scale of the given key
func generateKeyScale(chosenKey int) map[int]string {
	var note = chosenKey
	var newModalScaleSteps = map[int]string{}

	for i := 1; i < 8; i++ {
		newModalScaleSteps[i] = noteMap[note]
		note += scaleStepsMap[i]
	}

	return newModalScaleSteps
}

// Generates the modal scale of the chosen mode
// This and generateModalChords can be refactored into a generic function
func generateScaleNotes(majorScale map[int]string, chosenMode int) map[int]string {
	if chosenMode == 1 {
		return majorScale
	}

	var modalScaleNotes = map[int]string{}
	var degree = 1

	for i := chosenMode; i <= len(majorScale); i++ {
		modalScaleNotes[degree] = majorScale[i]
		degree++
	}

	for i := 1; i < chosenMode; i++ {
		modalScaleNotes[degree] = majorScale[i]
		degree++
	}

	return modalScaleNotes
}

// This maps the generetaed notes and chords together, outputting a modal scale
func mapNotesToChords(modalScale map[int]string, modalNotes map[int]string) map[int]chord {
	var newModalScale = map[int]chord{}

	for i := 1; i < 8; i++ {
		newModalScale[i] = chord{
			note:  modalNotes[i],
			chord: modalScale[i],
		}
	}

	return newModalScale
}

// Generates a chord progression based on a passed scale, and progression length
func generateChordProgression(modalScale map[int]chord, progressionLength int) map[int]chord {
	var chordProgression = map[int]chord{}
	max := 7
	min := 1

	rand.Seed(time.Now().UnixNano())
	chordProgression[1] = modalScale[1]

	for i := 2; i <= progressionLength; i++ {
		chordProgression[i] = modalScale[rand.Intn(max-min)+min]
	}

	return chordProgression
}
