package main

/* Various types and maps for scales and degrees */

// Chord object ...
type Chord struct {
	Note  string `json:"note"`
	Chord string `json:"chord"`
}

var modeMap = map[int]string{
	1: "Ionian",
	2: "Dorian",
	3: "Phrygian",
	4: "Lydian",
	5: "Mixolydian",
	6: "Aoelian",
	7: "Locrian",
}

var majorScaleChordMap = map[int]string{
	1: "Major",
	2: "Minor",
	3: "Minor",
	4: "Major",
	5: "Major",
	6: "Minor",
	7: "Diminished",
}

var noteMap = map[int]string{
	1:  "C",
	2:  "Db",
	3:  "D",
	4:  "Eb",
	5:  "E",
	6:  "F",
	7:  "F#",
	8:  "G",
	9:  "Ab",
	10: "A",
	11: "Bb",
	12: "B",
	13: "C",
	14: "Db",
	15: "D",
	16: "Eb",
	17: "E",
	18: "F",
	19: "F#",
	20: "G",
	21: "Ab",
	22: "A",
	23: "Bb",
	24: "B",
}

var scaleStepsMap = map[int]int{
	1: 2,
	2: 2,
	3: 1,
	4: 2,
	5: 2,
	6: 2,
	7: 1,
}
