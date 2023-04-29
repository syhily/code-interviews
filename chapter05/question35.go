package chapter05

import (
	"strconv"
	"strings"
)

func minimalTimeDelta(t ...string) int {
	var times [1441]bool
	for _, s := range t {
		time := parseTime(s)
		if time == 0 {
			if times[0] || times[1440] {
				return 0
			}

			times[0] = true
			times[1440] = true

			continue
		}

		if times[time] {
			return 0
		}

		times[time] = true
	}

	prev, minimal := -1, -1
	for i := 0; i < len(times); i++ {
		if !times[i] {
			continue
		}

		if prev == -1 {
			prev = i
			continue
		}

		delta := i - prev
		if minimal == -1 || delta < minimal {
			minimal = delta
		}
		prev = i
	}

	return minimal
}

func parseTime(s string) int {
	parts := strings.Split(s, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	return h*60 + m
}
