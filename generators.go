package main

/* Contains generators for modal scales and progressions */

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

func generateScaleSteps(chosenKey int) map[int]string {
	var note = chosenKey
	var newModalScaleSteps = map[int]string{}

	for i := 1; i < 8; i++ {
		newModalScaleSteps[i] = noteMap[note]
		note += scaleStepsMap[i]
	}

	return newModalScaleSteps
}

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

func mapNotesToChords(modalScale map[int]string, modalNotes map[int]string) map[int]map[string]string {
	var newMap = map[int]map[string]string{}

	for i := 1; i < 8; i++ {
		newMap[i] = map[string]string{}
		newMap[i][modalNotes[i]] = modalScale[i]
	}

	return newMap
}
