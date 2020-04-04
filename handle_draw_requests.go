package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleDrawRequests(rsp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	var username string
	usernames, ok := params["username"]
	if ok {
		username = usernames[0]
	} else {
		log.Println("username is nil")
		return
	}
	awardBatch := GetAward(username)
	if awardBatch == nil {
		rsp.Write([]byte("sorry you didn't win any prize"))
	} else {
		words := fmt.Sprintf("congratutions ! you won a %s", awardBatch.GetName())
		rsp.Write([]byte(words))
	}
}

