package main

import "fmt"

func main() {
	chosenMode := 2
	chosenKey := 3
	progressionLength := 4

	modalScaleChords := generateModalChords(chosenMode)
	modalScaleSteps := generateModalSteps(chosenMode)
	modalScaleNotes := generateModalNotes(chosenKey, modalScaleSteps)
	modalScale := mapModalScale(modalScaleNotes, modalScaleChords)
	chordProgression := generateChordProgression(modalScale, progressionLength)

	fmt.Println("Chosen Key: ", noteMap[chosenKey], modeMap[chosenMode])
	fmt.Println(modalScale)
	fmt.Println(chordProgression)
}
