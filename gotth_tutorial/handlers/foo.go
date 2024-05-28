package handlers

import (
	"gotth_tutorial/views/foo"
	"net/http"
)

func Handlefoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, foo.Index())
}
