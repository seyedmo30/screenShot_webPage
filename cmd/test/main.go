package main

import (
	"context"
	"screenShot-service/adapter/datastore/postgres"
)

func main() {
	// postgres.Setup()
	var gg = postgres.Postgres{}
	kk, _ := gg.CreateScreenShot(context.Background(), "dserf")
	println("___________________________")
	println(kk.Id)
	println(kk.Status)
	println(kk.FilePath)
	println(kk.Url)
	// println(err.Error())
	println("___________________________")
}
