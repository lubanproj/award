package main

import (
	"fmt"
	"net/http"
)

func handleDrawRequests(rsp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	var username string
	usernames, ok := params["username"]
	if ok {
		username = usernames[0]
	} else {
		fmt.Println("username is empty")
		rsp.Write([]byte("username cannot be empty"))
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

