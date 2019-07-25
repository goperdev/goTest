package main
import (
	. "goTest/Router"
)

func main() {
	router := InitRouter()
	router.Run(":7015")
}
