package web

import (
	"log"
	"net/http"
	"strconv"
)

var counter = 0

func MyCompHandler(w http.ResponseWriter, r *http.Request) {
	counter++
	component := MyComp(strconv.Itoa(counter))
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in MyCompHandler: %e", err)
	}
}
