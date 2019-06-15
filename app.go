package main

import "fmt"

func main() {
	chosenMode := 2
	chosenKey := 3

	println("Mode:", noteMap[chosenKey], "", modeMap[chosenMode])

	fmt.Println(mapNotesToChords(generateModalChords(chosenMode),
		generateScaleNotes(generateScaleSteps(chosenKey), chosenMode)))

}
