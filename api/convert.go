package handler

import (
	"fmt"
	"net/http"
)

// Convert handler function to convert any web content from a certain URL to markdown
func Convert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WIP: Convert API is called!")
}
