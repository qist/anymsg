package main

import (
	"fmt"

	"github.com/qist/anymsg/http"
)

const cfgFileName = "cfg.json"

func main() {

	Jcfg, err := getConfig(cfgFileName)
	if err != nil {
		panic(err)
	}

	http.SrvStart(Jcfg)

	fmt.Println("msg-sender servise runing... ")
}
