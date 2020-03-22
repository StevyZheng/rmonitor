package main

import "rmonitor/internals/app/routers"

func main() {
	println("rmonitor service is running...")
	router := routers.InitRouter()
	_ = router.Run(":80")
}