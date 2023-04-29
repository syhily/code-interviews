package chapter05

import "sort"

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func groupingAnagrams(s ...string) [][]string {
	grouping := make(map[string][]string)

	for _, z := range s {
		// Calculate a unique key.
		slice := RuneSlice(z)
		sort.Sort(slice)
		key := string(slice)

		ss := grouping[key]
		ss = append(ss, z)
		grouping[key] = ss
	}

	var res [][]string
	for _, s := range grouping {
		res = append(res, s)
	}
	return res
}
