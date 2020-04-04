package main

import "net/http"

func init() {
	InitLogConf()
	InitAwardPool()
}

func main() {

	parseConf()

	Start()
}


// Start starts to handle http requests
func Start() {
	http.HandleFunc("/draw", handleDrawRequests)
	http.ListenAndServe(":8080", nil)
}