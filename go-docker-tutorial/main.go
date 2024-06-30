package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {

		log.Fatal(err)
	}

	if err, err2 := store.Init(); err != nil {
		log.Fatal(err)
	} else if err2 != nil {
		log.Fatal(err2)
	}
	server := NewApiServer(":3000", store)
	// server := NewApiServer(":3001", nil)
	server.Run()
}
