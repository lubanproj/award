package main

import "net/http"

func init() {
	// parseConf must before InitAwardPool
	parseConf()
	InitLogConf()
	InitAwardPool()
}

func main() {
	Start()
}


// Start starts to handle http requests
func Start() {
	http.HandleFunc("/draw", handleDrawRequests)
	http.ListenAndServe(":8080", nil)
}