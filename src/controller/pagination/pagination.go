package pagination

import (
	"io"
	"net/http"
)

func PrevPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "prevPage")
}

func NextPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "nextPage")
}

func GoToPage(w http.ResponseWriter, r *http.Request) {
	pageValue := r.FormValue("pageValue")
	io.WriteString(w, pageValue)
}
