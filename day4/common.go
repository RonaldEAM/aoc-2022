package main

import (
	"log"
	"strconv"
	"strings"
)

type sectionRange struct {
	start int
	end   int
}

func parseRanges(line string) (sectionRange, sectionRange) {
	sectionRangesStr := strings.Split(line, ",")
	parsedRanges := []sectionRange{}
	for _, rangeStr := range sectionRangesStr {
		sectionRangeSlice := strings.Split(rangeStr, "-")

		start, errStart := strconv.Atoi(strings.TrimSpace(sectionRangeSlice[0]))
		end, errEnd := strconv.Atoi(strings.TrimSpace(sectionRangeSlice[1]))
		if errStart != nil || errEnd != nil {
			log.Fatal("parseRanges: error parsing input", errStart, errEnd)
		}

		parsedRanges = append(parsedRanges, sectionRange{
			start: start,
			end:   end,
		})
	}
	return parsedRanges[0], parsedRanges[1]
}
