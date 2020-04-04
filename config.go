package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Global config
var Conf = struct {
	// Award info
	Award struct {
		StartTime string `toml:"startTime"`  // start time of prizes
		EndTime string `toml:"endTime"`      // end time of prize delivery

	}

	// AwardMap store all prize categories and the initial number
	AwardMap map[string]int64

	// Global redis config
	Redis struct {
		Ip string `toml:"ip""`  // redis-server ip
		Port uint32 `toml:"port"`   // redis-server port
		Network string `toml:"network"`  // network
	}

	// Global mysql config
	Mysql struct {
		Dsn string `toml:"Dsn"`  // mysql dsn
	}
}{}

// parse conf
func parseConf() {
	if _, err := toml.DecodeFile("./conf/config.toml", &Conf); err != nil {
		fmt.Println("parse config error , ", err)
	}
 	fmt.Println("awardMap : ", Conf.AwardMap)
}
