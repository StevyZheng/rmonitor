package main

import (
	"rmonitor/internals/app/routers"
	"rmonitor/internals/pkg/models/database"
)

func main() {
	println("rmonitor service is running...")
	err := database.InitDB("rmonitor", "127.0.0.1", "27017", "stevy", "000000")
	if err != nil {
		println(err.Error())
		return
	}
	router := routers.InitRouter()
	_ = router.Run(":80")
}
