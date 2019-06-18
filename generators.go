package main

import (
	"math/rand"
	"strconv"
	"time"
)

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

// Generates the steps of the chosen modal scale
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

// Generates the notes of the corresponding modal scale based on a chosen key
func generateModalNotes(chosenKey int, modalScaleSteps map[int]int) map[int]string {
	var currentStep = chosenKey
	var modalNotes = map[int]string{}

	for i := 1; i < 8; i++ {
		modalNotes[i] = noteMap[currentStep]
		currentStep += modalScaleSteps[i]
	}

	return modalNotes
}

// This maps the generated notes and chords together, outputting a modal scale
func mapModalScale(modalScaleNotes map[int]string, modalScaleChords map[int]string) map[int]Chord {
	var modalScale = map[int]Chord{}

	for i := 1; i < 8; i++ {
		modalScale[i] = Chord{
			Note:  modalScaleNotes[i],
			Chord: modalScaleChords[i],
		}
	}

	return modalScale
}

// Generates a chord progression based on a passed scale, and progression length
func generateChordProgression(modalScale map[int]Chord, progressionLength int) map[int]Chord {
	var chordProgression = map[int]Chord{}
	max := 7
	min := 1

	rand.Seed(time.Now().UnixNano())
	chordProgression[1] = modalScale[1]

	for i := 2; i <= progressionLength; i++ {
		chordProgression[i] = modalScale[rand.Intn(max-min)+min]
	}

	return chordProgression
}

// Generates a progression object based on passed paramaters, contains scale, key, and mode types
func generateModalScale(scale map[int]Chord, key int, mode int) Scale {
	var newScale Scale

	newScale.Key = noteMap[key]
	newScale.Mode = modeMap[mode]

	for i := 1; i < 8; i++ {
		newScale.Chords = append(newScale.Chords, scale[i])
	}

	println(newScale.Chords)

	return newScale
}

// Create a progression object, contains key, mode and chord progression
func createProgressionObject(progression map[int]Chord, key int, mode int, progressionLength int) Progression {
	var newProgression Progression

	newProgression.Key = noteMap[key]
	newProgression.Mode = modeMap[mode]
	newProgression.ProgressionLength = strconv.Itoa(progressionLength)

	for i := 1; i <= progressionLength; i++ {
		newProgression.Chords = append(newProgression.Chords, progression[i])
	}

	return newProgression
}
