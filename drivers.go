package main

// Function containing sequence of calls to generate chord progression, and return
func serveChordProgression(key int, mode int, progressionLength int) map[int]Chord {
	modalScaleChords := generateModalChords(mode)
	modalScaleSteps := generateModalSteps(mode)
	modalScaleNotes := generateModalNotes(key, modalScaleSteps)
	modalScale := mapModalScale(modalScaleNotes, modalScaleChords)
	return generateChordProgression(modalScale, progressionLength)
}

func serveModalScale(key int, mode int, Progression int)
