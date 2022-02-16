package PostRequest1

import (
	pr "FinalApi/post"
	"fmt"
	"log"
	"net/http"
)

func homepage(m http.ResponseWriter, r *http.Request) {
	fmt.Fprint(m, "Welcome to my API! - Madhosh Yagnik")

}
func PostRequest1() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/Word", pr.PostRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
