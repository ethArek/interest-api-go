package handlers

import (
	"fmt"
	"net/http"
)

// Root handler
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Docker test server running!")
}
