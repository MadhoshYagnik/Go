package post

import (
	"log"
	"regexp"
	"strings"
)

func WordC(s []string) map[string]int {
	words := strings.Split(strings.Join(s, ""), " ")
	m := make(map[string]int)

	for _, word := range words {
		re, err := regexp.Compile(`[^\w]`)
		if err != nil {
			log.Fatal(err)
		}
		word = re.ReplaceAllString(word, "")
		_, ok := m[word]
		if ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	return m
}
