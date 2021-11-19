package main

import (
	"anymsg/http"
	"flag"
	"fmt"
	"os"
)

const cfgFileName = "cfg.json"



func main() {
	argsWithProg := os.Args
	if(len(argsWithProg) < 2) {
		fmt.Println("usage : ",argsWithProg[0],"-h")
		return
	}
	var path string
	flag.StringVar(&path,"f","/etc/cfg.json","config")
	flag.Parse()
	Jcfg, err := getConfig(cfgFileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	http.SrvStart(Jcfg)

	fmt.Println("msg-sender servise runing... ")
}
