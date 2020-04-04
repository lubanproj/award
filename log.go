package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Use go log package
// Initialize a local log
func InitLogConf() error {
	y, m, d := time.Now().Date()
	fileName := fmt.Sprintf("%d-%d-%d_award.log", y, m, d)
	f , err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("file open error, %v", err)
		return err
	}

	// display line numbers
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)

	return nil
}