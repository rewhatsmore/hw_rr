package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	word      string
	frequency int
}

// looking for dashes and apostrophes if they are not inside a word
// as well as other punctuation marks
// and also digits, because they are not words.
var re = regexp.MustCompile(`\s-|-\s|\s'|'\s|[,\.!?;:()"0-9]`)

func Top10(s string) []string {
	s = re.ReplaceAllString(strings.ToLower(s), " ")
	segments := strings.Fields(s)
	dict := make(map[string]int)
	for _, word := range segments {
		dict[word]++
	}

	words := make([]Word, 0, len(dict))
	for word, frequency := range dict {
		words = append(words, Word{word, frequency})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].frequency != words[j].frequency {
			return words[i].frequency > words[j].frequency
		}
		return words[i].word < words[j].word
	})

	n := 10
	if len(words) < 10 {
		n = len(words)
	}

	result := make([]string, n)
	for i := range result {
		result[i] = words[i].word
	}
	return result
}
