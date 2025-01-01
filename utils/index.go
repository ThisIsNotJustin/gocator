package utils

import "sort"

type Index map[string][]int

func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}

			idx[token] = append(ids, doc.ID)
		}
	}
}

func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[i] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	return r
}

func (idx Index) Search(text string) []int {
	tokens := analyze(text)
	if len(tokens) == 0 {
		return nil
	}

	lists := make([][]int, 0, len(tokens))
	for _, token := range tokens {
		if ids, ok := idx[token]; ok {
			lists = append(lists, ids)
		} else {
			return nil
		}
	}
	sort.Slice(lists, func(i, j int) bool {
		return len(lists[i]) < len(lists[j])
	})

	result := lists[0]
	for _, ids := range lists[1:] {
		result = Intersection(result, ids)
		if len(result) == 0 {
			return nil
		}
	}

	return result
}
