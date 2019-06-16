package main

// Chord object ...
type Chord struct {
	Note  string `json:"note"`
	Chord string `json:"chord"`
}

// Progression object ...
type Progression struct {
	Key               string  `json:"key"`
	Mode              string  `json:"mode"`
	ProgressionLength string  `json:"progressionlength"`
	Chords            []Chord `json:"chords"`
}

// Scale object ...
type Scale struct {
	Key    string  `json:"key"`
	Mode   string  `json:"mode"`
	Chords []Chord `json:"chords"`
}
