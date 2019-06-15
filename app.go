package main

import "fmt"

func main() {
	chosenMode := 2
	chosenKey := 3
	progressionLength := 4

	println("Mode:", noteMap[chosenKey], "", modeMap[chosenMode])

	scale := mapNotesToChords(generateModalChords(chosenMode),
		generateScaleNotes(generateScaleSteps(chosenKey), chosenMode))

	chordProgression := generateChordProgression(scale, progressionLength)

	fmt.Println(chordProgression)
}
