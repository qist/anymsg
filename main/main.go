package main

import (
	"fmt"

	"anymsg/http"
)

const cfgFileName = "cfg.json"

func main() {

	Jcfg, err := getConfig(cfgFileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	http.SrvStart(Jcfg)

	fmt.Println("msg-sender servise runing... ")
}
