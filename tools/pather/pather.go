package pather

import (
	"fmt"
	"strings"
)

func Path(puzzleID string, sample bool, test bool) string {
	p := ""
	if test {
		p += "../../"
	}
	p += fmt.Sprintf("inputs/%s", strings.Replace(puzzleID, "-", "/", 1))
	if sample {
		p += "-sample"
	}
	p += ".txt"

	return p
}
