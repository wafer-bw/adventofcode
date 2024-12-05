package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Update struct {
	Pages []*Page
}

func (u *Update) Push(page *Page) {
	if len(u.Pages) == 0 {
		u.Pages = append(u.Pages, page)
		return
	}
	last := u.Pages[len(u.Pages)-1]
	last.Next = page
	page.Prev = last
	u.Pages = append(u.Pages, page)
}

type Page struct {
	Number int
	Next   *Page
	Prev   *Page
}

func (p *Page) Before(n int) bool {
	seen := map[int]struct{}{}
	pg := p.Next
	for pg != nil {
		seen[pg.Number] = struct{}{}
		if pg.Number == n {
			return true
		}
		pg = pg.Next
	}
	pg = p.Prev
	for pg != nil {
		seen[pg.Number] = struct{}{}
		if pg.Number == n {
			return false
		}
		pg = pg.Prev
	}

	if _, ok := seen[n]; !ok {
		return true
	}

	return false
}

func Solve(input string) int {
	s := 0

	parts := strings.Split(input, "\n\n")
	rulesLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	updatesLines := strings.Split(strings.TrimSpace(parts[1]), "\n")

	rulesMap := map[int]map[int]struct{}{}
	for _, rule := range rulesLines {
		parts := strings.Split(strings.TrimSpace(rule), "|")
		subject, _ := strconv.Atoi(parts[0])
		before, _ := strconv.Atoi(parts[1])
		if _, ok := rulesMap[subject]; !ok {
			rulesMap[subject] = map[int]struct{}{}
		}
		rulesMap[subject][before] = struct{}{}
	}

	updates := make([]*Update, 0, len(updatesLines))
	for _, updateStr := range updatesLines {
		parts := strings.Split(strings.TrimSpace(updateStr), ",")
		update := &Update{Pages: []*Page{}}
		for _, page := range parts {
			number, _ := strconv.Atoi(page)
			page := &Page{Number: number}
			update.Push(page)
		}
		updates = append(updates, update)
	}

	for _, update := range updates {
		if valid(update, rulesMap) {
			s += middleOfSlice(update.Pages).Number
		}
	}

	return s
}

func middleOfSlice[T any](s []T) T {
	return s[len(s)/2]
}

func valid(update *Update, rulesMap map[int]map[int]struct{}) bool {
	for _, page := range update.Pages {
		befores, ok := rulesMap[page.Number]
		if !ok {
			continue
		}
		for before := range befores {
			if !page.Before(before) {
				return false
			}
		}
	}
	return true
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
