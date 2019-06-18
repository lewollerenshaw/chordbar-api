package main

// Function containing sequence of calls to generate chord progression, and return
func serveChordProgression(key int, mode int, progressionLength int) Progression {
	modalScaleChords := generateModalChords(mode)
	modalScaleSteps := generateModalSteps(mode)
	modalScaleNotes := generateModalNotes(key, modalScaleSteps)
	modalScale := mapModalScale(modalScaleNotes, modalScaleChords)
	generatedProgression := generateChordProgression(modalScale, progressionLength)
	return createProgressionObject(generatedProgression, key, mode, progressionLength)
}

// Function containing sequence of calls to general modal scale, and return object
func serveModalScale(key int, mode int) Scale {
	modalScaleChords := generateModalChords(mode)
	modalScaleSteps := generateModalSteps(mode)
	modalScaleNotes := generateModalNotes(key, modalScaleSteps)
	modalScale := mapModalScale(modalScaleNotes, modalScaleChords)
	return generateModalScale(modalScale, key, mode)
}
