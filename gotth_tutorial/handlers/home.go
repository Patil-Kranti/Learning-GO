package handlers

import (
	"gotth_tutorial/views/home"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
