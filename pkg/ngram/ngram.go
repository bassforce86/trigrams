package ngram

import (
	"strings"
	"unicode"
)

type Ngrams []Ngram

type Ngram struct {
	Words     string
	Frequency uint32
}

func (n Ngrams) Len() int      { return len(n) }
func (n Ngrams) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Ngrams) Less(i, j int) bool {
	less := n[i].Frequency > n[j].Frequency
	if n[i].Frequency == n[j].Frequency {
		less = n[i].Words < n[j].Words
	}
	return less
}

func BuildTrigram(sequence string) Ngrams {
	return build(sequence, 3)
}

func build(sequence string, n int) Ngrams {
	var result Ngrams
	words := SplitOnNonLetters(sequence)
	ngrams := noOfNgrams(words, 3)

	for word, count := range ngrams {
		result = append(result, Ngram{Words: word, Frequency: count})
	}

	return result
}

func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func noOfNgrams(words []string, size int) (count map[string]uint32) {

	count = make(map[string]uint32, 0)
	offset := int(size / 2)

	max := len(words)
	for i, _ := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		count[gram]++
	}

	return count
}
