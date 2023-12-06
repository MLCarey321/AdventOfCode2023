package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type RangeMap struct {
	destStart   int
	sourceStart int
	length      int
}

func (rm RangeMap) destEnd() int {
	return rm.destStart + rm.length - 1
}

func (rm RangeMap) sourceEnd() int {
	return rm.sourceStart + rm.length - 1
}

func (rm RangeMap) shiftAmount() int {
	return rm.destStart - rm.sourceStart
}

func mapRanges(r []Range, m []RangeMap) []Range {
	if len(r) == 0 {
		return r
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i].start < r[j].start
	})
	sort.Slice(m, func(i, j int) bool {
		return m[i].sourceStart < m[j].sourceStart
	})
	var retval []Range
	ri := 0
	mi := 0
	for ri < len(r) && mi < len(m) {
		if r[ri].start < m[mi].sourceStart {
			// range starts before current mapping starts
			if r[ri].end < m[mi].sourceStart {
				// range ends before current mapping starts
				retval = append(retval, r[ri])
				ri++
			} else {
				// range ends after current mapping starts
				// range is partially before mapping, partially within
				retval = append(retval, Range{r[ri].start, m[mi].sourceStart - 1})
				r[ri].start = m[mi].sourceStart
			}
		} else if r[ri].start <= m[mi].sourceEnd() {
			// range starts on or after current mapping start
			// range starts before current mapping end
			if r[ri].end <= m[mi].sourceEnd() {
				// range ends on or before current mapping end
				// range is wholly within the mapping
				retval = append(retval, Range{r[ri].start + m[mi].shiftAmount(), r[ri].end + m[mi].shiftAmount()})
				if r[ri].end == m[mi].sourceEnd() {
					mi++
				}
				ri++
			} else {
				// range ends after current mapping end
				// range spills over
				retval = append(retval, Range{r[ri].start + m[mi].shiftAmount(), m[mi].sourceEnd() + m[mi].shiftAmount()})
				r[ri].start = m[mi].sourceEnd() + 1
				mi++
			}
		} else {
			// mapping starts and ends prior to range
			mi++
		}
	}
	for ri < len(r) {
		// leftover ranges
		retval = append(retval, r[ri])
		ri++
	}
	return retval
}

func main() {
	file, err := os.Open("day05.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ranges1 []Range
	var ranges2 []Range
	var rangeMaps []RangeMap
	for scanner.Scan() {
		line := scanner.Text()
		if len(ranges1) == 0 {
			seeds := strings.Split(line, " ")[1:]
			for i := 0; i < len(seeds); i += 2 {
				first, err := strconv.Atoi(seeds[i])
				if err != nil {
					panic(err)
				}
				second, err := strconv.Atoi(seeds[i+1])
				if err != nil {
					panic(err)
				}
				ranges1 = append(ranges1, Range{first, first})
				ranges1 = append(ranges1, Range{second, second})
				ranges2 = append(ranges2, Range{first, first + second - 1})
			}
		} else if line == "" {
			ranges1 = mapRanges(ranges1, rangeMaps)
			ranges2 = mapRanges(ranges2, rangeMaps)
			rangeMaps = []RangeMap{}
		} else if strings.Contains(line, ":") {
			continue
		} else {
			parts := strings.Split(line, " ")
			d, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			s, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			l, err := strconv.Atoi(parts[2])
			if err != nil {
				panic(err)
			}
			rangeMaps = append(rangeMaps, RangeMap{d, s, l})
		}
	}
	ranges1 = mapRanges(ranges1, rangeMaps)
	ranges2 = mapRanges(ranges2, rangeMaps)
	rangeMaps = []RangeMap{}

	sort.Slice(ranges1, func(i, j int) bool {
		return ranges1[i].start < ranges1[j].start
	})
	sort.Slice(ranges2, func(i, j int) bool {
		return ranges2[i].start < ranges2[j].start
	})

	fmt.Printf("Part 1: %d\n", ranges1[0].start)
	fmt.Printf("Part 2: %d\n", ranges2[0].start)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
