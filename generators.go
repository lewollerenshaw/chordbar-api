package main

import "fmt"

/* Contains generators for modal scales and progressions */

func generateModalScale(chosenMode int) map[int]string {
	if chosenMode == 0 {
		return majorscalelist
	}

	println("Mode: ", chosenMode, "(", modelist[chosenMode], ")")

	var newModalScale = map[int]string{}

	for i := chosenMode; i <= len(majorscalelist); i++ {
		newModalScale[i] = majorscalelist[i]
		fmt.Println(newModalScale[i])
	}

	for i := 1; i < chosenMode; i++ {
		newModalScale[i] = majorscalelist[i]
		fmt.Println(newModalScale[i])
	}

	return newModalScale
}
