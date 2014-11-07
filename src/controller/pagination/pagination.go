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
