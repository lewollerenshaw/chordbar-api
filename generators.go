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

func generateModalSteps(chosenMode int) map[int]int {
	if chosenMode == 1 {
		return scaleStepsMap
	}

	var modalScaleSteps = map[int]int{}
	var degree = 1

	for i := chosenMode; i <= len(majorScaleChordMap); i++ {
		modalScaleSteps[degree] = scaleStepsMap[i]
		degree++
	}

	for i := 1; i < chosenMode; i++ {
		modalScaleSteps[degree] = scaleStepsMap[i]
		degree++
	}
	return modalScaleSteps
}

func generateModalNotes(chosenKey int, modalScaleSteps map[int]int) map[int]string {
	var currentStep = chosenKey
	var modalNotes = map[int]string{}

	for i := 1; i < 8; i++ {
		modalNotes[i] = noteMap[currentStep]
		currentStep += modalScaleSteps[i]
	}

	return modalNotes
}

// This maps the generetaed notes and chords together, outputting a modal scale
func mapModalScale(modalScaleNotes map[int]string, modalScaleChords map[int]string) map[int]chord {
	var modalScale = map[int]chord{}

	for i := 1; i < 8; i++ {
		modalScale[i] = chord{
			note:  modalScaleNotes[i],
			chord: modalScaleChords[i],
		}
	}

	return modalScale
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
