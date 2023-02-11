// https://adventofcode.com/2022/day/13#part2

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"golang.org/x/exp/slices"
)

const (
	puzzleID string = "2022-13"
)

var (
	dividerPackets []json.RawMessage = []json.RawMessage{
		json.RawMessage("[[2]]"),
		json.RawMessage("[[6]]"),
	}
)

func GetData[T any](d json.RawMessage) (T, bool) {
	var i T
	if err := json.Unmarshal(d, &i); err == nil {
		return i, true
	}
	return i, false
}

func solve(lines []string) int {
	packetPairs := getPackets(lines)
	slices.SortFunc(packetPairs, func(a, b json.RawMessage) bool {
		c, r := isCorrupt(0, a, b)
		if !r {
			panic("unresolved")
		}
		return !c
	})

	key := 1
	for i, p := range packetPairs {
		if string(p) == string(dividerPackets[0]) || string(p) == string(dividerPackets[1]) {
			key *= i + 1
		}
	}

	return key
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, true, false))))
}

func getPackets(lines []string) []json.RawMessage {
	packets := []json.RawMessage{}
	for _, ln := range lines {
		if ln == "" {
			continue
		}
		packet := json.RawMessage{}
		_ = json.Unmarshal([]byte(ln), &packet)
		packets = append(packets, packet)
	}
	packets = append(packets, dividerPackets...)
	return packets
}

func isCorrupt(depth int, left json.RawMessage, right json.RawMessage) (bool, bool) {
	prefix := ""
	depth = depth + 1
	for i := 0; i < depth; i++ {
		prefix = prefix + " "
	}

	leftInt, leftIsInt := GetData[int](left)
	rightInt, rightIsInt := GetData[int](right)
	leftList, leftIsList := GetData[[]json.RawMessage](left)
	rightList, rightIsList := GetData[[]json.RawMessage](right)

	if leftIsInt && rightIsInt {
		if leftInt > rightInt {
			return true, true
		} else if leftInt < rightInt {
			return false, true
		} else {
			return false, false
		}
	} else if leftIsList && rightIsList {
		shorterLen := shorterLength(leftList, rightList)
		for i := 0; i < shorterLen; i++ {
			left := leftList[i]
			right := rightList[i]
			if corrupt, resolved := isCorrupt(depth, left, right); resolved {
				return corrupt, true
			}
		}
		if len(left) < len(right) {
			return false, true
		} else if len(left) > len(right) {
			return true, true
		} else {
			return false, false
		}
	} else if leftIsList && rightIsInt {
		return isCorrupt(depth, left, json.RawMessage(fmt.Sprintf("[%d]", rightInt)))
	} else if !leftIsList && rightIsList {
		return isCorrupt(depth, json.RawMessage(fmt.Sprintf("[%d]", leftInt)), right)
	}

	return false, false
}

func shorterLength(left, right []json.RawMessage) int {
	if len(left) < len(right) {
		return len(left)
	}
	return len(right)
}
