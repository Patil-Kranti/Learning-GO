//go:build !dev
// +build !dev

package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
