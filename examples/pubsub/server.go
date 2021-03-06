package main

import (
	"github.com/jcelliott/turnpike"
	"net/http"
)

func main() {
	s := turnpike.NewServer()

	http.Handle("/ws", s.Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
