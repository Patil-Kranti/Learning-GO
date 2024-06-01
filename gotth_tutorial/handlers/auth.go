package handlers

import (
	"gotth_tutorial/views/auth"
	"net/http"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
