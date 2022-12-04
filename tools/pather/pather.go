package pather

import "fmt"

func Path(puzzleID string, sample bool, test bool) string {
	p := ""
	if test {
		p += "../"
	}
	p += fmt.Sprintf("inputs/%s", puzzleID)
	if sample {
		p += "-sample"
	}
	p += ".txt"

	return p
}
