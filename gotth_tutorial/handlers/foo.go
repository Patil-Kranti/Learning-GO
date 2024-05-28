package handlers

import "net/http"

func Handlefoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World! from foo1"))
}
