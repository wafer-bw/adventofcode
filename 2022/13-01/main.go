// https://adventofcode.com/2022/day/13

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID string = "2022-13"
)

func GetData[T any](d json.RawMessage) (T, bool) {
	var i T
	if err := json.Unmarshal(d, &i); err == nil {
		return i, true
	}
	return i, false
}

func solve(lines []string) int {
	corruptPairs := []int{}
	validPairs := []int{}
	pairs := getPairs(lines)
	packetPairs := getPacketPairs(pairs)

	for i, p := range packetPairs {
		idx := i + 1
		fmt.Println()
		log.Printf("pair %d:", idx)
		log.Printf(" left: %s", p[0])
		log.Printf(" right: %s", p[1])
		if corrupt, resolved := isCorrupt(idx, 0, p[0], p[1]); corrupt && resolved {
			corruptPairs = append(corruptPairs, idx)
		} else if !corrupt && resolved {
			validPairs = append(validPairs, idx)
		} else {
			return -1
		}
	}

	sum := 0
	for _, i := range validPairs {
		sum += i
	}

	return sum
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, true, false))))
}

func getPacketPairs(pairs [][]string) [][]json.RawMessage {
	packets := [][]json.RawMessage{}
	for _, pair := range pairs {
		packetPair := []json.RawMessage{}
		for _, pstr := range pair {
			packet := json.RawMessage{}
			_ = json.Unmarshal([]byte(pstr), &packet)
			packetPair = append(packetPair, packet)
		}
		packets = append(packets, packetPair)
	}

	return packets
}

func getPairs(lines []string) [][]string {
	pairs := [][]string{}
	last := 0
	for i := 3; i <= len(lines)+1; i = i + 3 {
		pair := []string{}
		for j := last; j < i-1; j++ {
			pair = append(pair, lines[j])
		}
		pairs = append(pairs, pair)
		last = i
	}

	for _, pair := range pairs {
		fmt.Println(pair)
	}

	return pairs
}

func isCorrupt(idx int, depth int, left json.RawMessage, right json.RawMessage) (bool, bool) {
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
		log.Printf("%s%d: comparing ints %d vs %d", prefix, idx, leftInt, rightInt)
		if leftInt > rightInt {
			log.Printf("%s %d: corrupt because left is greater than right", prefix, idx)
			return true, true
		} else if leftInt < rightInt {
			log.Printf("%s %d: valid because left is less than right", prefix, idx)
			return false, true
		} else {
			return false, false
		}
	} else if leftIsList && rightIsList {
		shorterLen := shorterLength(leftList, rightList)
		for i := 0; i < shorterLen; i++ {
			left := leftList[i]
			right := rightList[i]
			if corrupt, resolved := isCorrupt(idx, depth, left, right); resolved {
				return corrupt, true
			}
		}
		if len(left) < len(right) {
			log.Printf("%s %d: valid because list comparison has shorter left list", prefix, idx)
			return false, true
		} else if len(left) > len(right) {
			log.Printf("%s %d: corrupt because list comparison has shorter right list", prefix, idx)
			return true, true
		} else {
			return false, false
		}
	} else if leftIsList && rightIsInt {
		return isCorrupt(idx, depth, left, json.RawMessage(fmt.Sprintf("[%d]", rightInt)))
	} else if !leftIsList && rightIsList {
		return isCorrupt(idx, depth, json.RawMessage(fmt.Sprintf("[%d]", leftInt)), right)
	}

	log.Printf("%s%d: somehow unresolved at end %s vs %s", prefix, idx, left, right)
	return false, false
}

func shorterLength(left, right []json.RawMessage) int {
	if len(left) < len(right) {
		return len(left)
	}
	return len(right)
}
