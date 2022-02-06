// Please create a small service that accepts as input a body of text, such as that from a book, and
// return the top ten most-used words along with how many times they occur in the text.

package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

type wordStruct struct {
	string
	int
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var buf bytes.Buffer
	b, err := r.GetPlainText() //ignoring fonts and styles
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	//
	content := reg.ReplaceAllString(buf.String(), " ")
	return content, nil
}

func wordCount(content string) []wordStruct {
	wordSlice := strings.Fields(content) // "a b cd"=> ["a","b","cd"]
	wordMap := make(map[string]int)

	for i := 0; i < len(wordSlice); i++ {
		wordMap[strings.ToLower(wordSlice[i])] += 1
	}
	words := make([]string, 0, len(wordMap))
	for word := range wordMap {
		words = append(words, word)
	}
	//
}

func main() {
	content, err := readPdf("pdf_sample.pdf")
	if err != nil {
		panic(err)
	}
	frequentWords := wordCount(content)
	for _, element := range frequentWords {
		fmt.Println(element.string, " => ", element.int)
	}
}
