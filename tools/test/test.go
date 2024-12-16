package test

import (
	"regexp"
	"strconv"
)

var SplitPattern = regexp.MustCompile(`\nEND-INPUT-SEPARATOR-WAFER-BW-END-INPUT-SEPARATOR-WAFER-BW (\d+)\n`)

type Case struct {
	Input  string
	Answer int
}

type Cases []Case

func GetCases(inputs string) Cases {
	cases := Cases{}
	lastIndex := 0
	matches := SplitPattern.FindAllStringSubmatchIndex(inputs, -1)
	for _, match := range matches {
		input := inputs[lastIndex:match[0]]
		answerStr := inputs[match[2]:match[3]]
		lastIndex = match[1]

		answer, err := strconv.Atoi(answerStr)
		if err != nil {
			panic(err)
		}

		cases = append(cases, Case{Input: input, Answer: answer})
	}

	return cases
}
