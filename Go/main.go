package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bbalet/stopwords"
)

var tpl *template.Template
var words string

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
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
	words = text

	out := topten(words)
	fmt.Println(out)
	outWithRemoval := toptenusestopwords(words)
	fmt.Println(outWithRemoval)

	// fmt.Println("hello ", text)
}

// word counter

func topten(words string) []string {
	input := strings.Fields(words)
	count := make(map[string]int)

	for _, word := range input {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		word := reg.ReplaceAllString(word, "")
		_, match := count[word]
		if match {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}

	keys := make([]string, 0, len(count))
	out := make([]string, 0)
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	for i, name := range keys {
		var s = strconv.Itoa(count[name])
		if i != 10 {
			out = append(out, name+":"+s)
		} else {
			break
		}
	}
	return out
}

func toptenusestopwords(words string) []string {
	cleanContent := stopwords.CleanString(words, "en", true)

	input := strings.Fields(cleanContent)
	count := make(map[string]int)

	for _, word := range input {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		word := reg.ReplaceAllString(word, "")
		_, match := count[word]
		if match {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}

	keys := make([]string, 0, len(count))
	out := make([]string, 0)
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	for i, name := range keys {
		var s = strconv.Itoa(count[name])
		if i != 10 {
			out = append(out, name+":"+s)
		} else {
			break
		}
	}
	return out
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)

}
