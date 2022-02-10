// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"net/http"
// 	"regexp"
// 	"sort"
// 	"strings"

// 	"github.com/ledongthuc/pdf"
// )

// type wordStruct struct {
// 	string
// 	int
// }

// func readPdf(path string) (string, error) {
// 	f, r, err := pdf.Open(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()
// 	var buf bytes.Buffer
// 	b, err := r.GetPlainText()
// 	if err != nil {
// 		return "", err
// 	}
// 	buf.ReadFrom(b)

// 	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
// 	if err != nil {
// 		return "", err
// 	}
// 	content := reg.ReplaceAllString(buf.String(), " ")
// 	return content, nil
// }

// func wordCount(content string) []wordStruct {
// 	wordSlice := strings.Fields(content)
// 	wordMap := make(map[string]int)

// 	// to lower character
// 	for i := 0; i < len(wordSlice); i++ {
// 		wordMap[strings.ToLower(wordSlice[i])] += 1
// 	}
// 	words := make([]string, 0, len(wordMap))
// 	for word := range wordMap {
// 		words = append(words, word)
// 	}
// 	// sorting the words before counting their occurrences
// 	sort.SliceStable(words, func(i, j int) bool {
// 		return wordMap[words[i]] > wordMap[words[j]]
// 	})

// 	topTenWords := make([]wordStruct, 0, 10)

// 	for i := 0; i < 10; i++ {
// 		//fmt.Println(words[i], " => "wordMap[words[i]])
// 		topTenWords = append(topTenWords, wordStruct{words[i], wordMap[words[i]]})
// 	}
// 	return topTenWords
// }

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

// func main() {
// 	http.HandleFunc("/", HelloServer)
// 	http.ListenAndServe(":8080", nil)

// 	content, err := readPdf("pdf-sample.pdf")
// 	if err != nil {
// 		panic(err)
// 	}
// 	frequentWords := wordCount(content)
// 	for _, element := range frequentWords {
// 		fmt.Println(element.string, " => ", element.int)
// 	}
// }
package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	text := r.FormValue("text")
	d := struct {
		Text string
	}{
		Text: text,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)

	// fmt.Println("hello ", text)
}
