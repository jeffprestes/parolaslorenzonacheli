package main

import "gopkg.in/macaron.v1"

func main()  {
	app := macaron.Classic()
	app.Run()
}