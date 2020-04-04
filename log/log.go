package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Use go log package
// Initialize a local log
func InitLogConf() {
	y, m, d := time.Now().Date()
	fileName := fmt.Sprintf("%d-%d-%d_award.log", y, m, d)
	f , err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("file open error, %v", err)
	}

	// display line numbers
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)
}