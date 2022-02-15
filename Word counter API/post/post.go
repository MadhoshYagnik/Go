package post

import (
	"io/ioutil"
	"net/http"
)

type CountW struct {
	word  string
	count int
}

type countW_ struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type CountWs []countW_

func PostRequest(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadAll(r.Body)
	s := []string{string(content)}

}
