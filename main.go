package main

import (
	"github.com/slyxh2/golang-blog/router"
	"github.com/slyxh2/golang-blog/settings"
)

func main() {
	settings.Init()
	db := settings.GetDbClient()
	router.Init(db)
}
