// https://adventofcode.com/2022/day/11#part2

package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"github.com/wafer-bw/adventofcode/tools/stack"
	"golang.org/x/exp/slices"
)

const (
	puzzleID  string = "2022-11"
	numRounds int    = 10000
)

func solve(lines []string) int {
	monkeys := setup(lines)

	modifier := 1
	for _, m := range monkeys {
		modifier = modifier * m.test
	}

	for i := 0; i < numRounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeyItems := monkeys[j].items
			for k := 0; k < len(monkeyItems); k++ {
				item := stack.PopBack(&monkeys[j].items)
				item = monkeys[j].inspect(item)
				item = relief(item, modifier)
				throwTo := monkeys[j].doTest(item)
				monkeys[throwTo].items = append(monkeys[throwTo].items, item)
			}
		}
	}

	monkeySlice := []monkey{}
	for _, m := range monkeys {
		monkeySlice = append(monkeySlice, *m)
	}

	slices.SortFunc(monkeySlice, func(a, b monkey) bool {
		return a.inspections > b.inspections
	})

	monkeyBusiness := monkeySlice[0].inspections * monkeySlice[1].inspections
	return monkeyBusiness
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

type monkey struct {
	id                int
	items             []int
	operationText     string
	test              int
	testTrueMonkeyID  int
	testFalseMonkeyID int
	inspections       int
}

func (m *monkey) inspect(oldWorry int) (newWorry int) {
	m.inspections++

	parts := strings.Split(m.operationText, " ")
	op := parts[1]

	a := 0
	if parts[0] == "old" {
		a = oldWorry
	} else {
		a, _ = strconv.Atoi(parts[0])
	}

	b := 0
	if parts[2] == "old" {
		b = oldWorry
	} else {
		b, _ = strconv.Atoi(parts[2])
	}

	switch op {
	case "+":
		newWorry = a + b
	case "*":
		newWorry = a * b
	default:
		log.Fatalf("unknown operation %s", op)
	}

	return
}

func (m *monkey) doTest(worry int) (monkeyID int) {
	if worry%m.test == 0 {
		monkeyID = m.testTrueMonkeyID
	} else {
		monkeyID = m.testFalseMonkeyID
	}

	return
}

func setup(lines []string) map[int]*monkey {
	var currentMonkey *monkey = nil
	monkeys := map[int]*monkey{}

	for _, ln := range lines {
		if strings.HasPrefix(ln, "Monkey") {
			monkeyNumberStr := strings.Split(ln, " ")[1]
			monkeyNumberStr = strings.ReplaceAll(monkeyNumberStr, ":", "")
			monkeyNumber, _ := strconv.Atoi(monkeyNumberStr)
			monkey := &monkey{id: monkeyNumber, items: []int{}}
			currentMonkey = monkey
			monkeys[monkeyNumber] = currentMonkey
		} else if strings.Contains(ln, "Starting items:") {
			itemsStr := strings.TrimSpace(strings.ReplaceAll(ln, "Starting items: ", ""))
			items := strings.Split(itemsStr, ", ")
			for _, item := range items {
				itemInt, _ := strconv.Atoi(item)
				currentMonkey.items = append(currentMonkey.items, itemInt)
			}
		} else if strings.Contains(ln, "Operation:") {
			operationStr := strings.TrimSpace(strings.ReplaceAll(ln, "Operation: new = ", ""))
			currentMonkey.operationText = operationStr
		} else if strings.Contains(ln, "Test:") {
			testParts := strings.Split(ln, " ")
			testStr := testParts[len(testParts)-1]
			test, _ := strconv.Atoi(testStr)
			currentMonkey.test = test
		} else if strings.Contains(ln, "If true:") {
			testTrueMonkeyParts := strings.Split(ln, " ")
			testTrueMonkeyStr := testTrueMonkeyParts[len(testTrueMonkeyParts)-1]
			testTrueMonkey, _ := strconv.Atoi(testTrueMonkeyStr)
			currentMonkey.testTrueMonkeyID = testTrueMonkey
		} else if strings.Contains(ln, "If false:") {
			testFalseMonkeyParts := strings.Split(ln, " ")
			testFalseMonkeyStr := testFalseMonkeyParts[len(testFalseMonkeyParts)-1]
			testFalseMonkey, _ := strconv.Atoi(testFalseMonkeyStr)
			currentMonkey.testFalseMonkeyID = testFalseMonkey
		}
	}

	return monkeys
}

func relief(oldWorry int, modifier int) (newWorry int) {
	newWorry = oldWorry % modifier
	return
}
