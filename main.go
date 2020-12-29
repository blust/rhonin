package main

import (
	"bw/app/router"

	eagle "github.com/bridge-wall/bw-eagle"
)

func main()  {
	err := eagle.New()
	if err != nil {
		panic(err)
	}

	router.Register()

	eagle.LogInfo(nil, "bw server start")

	eagle.Run()
}