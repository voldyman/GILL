package webserver

import (
	"github.com/go-martini/martini"
)

func StartServer() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "hello world"
	})

	m.Run()
}
