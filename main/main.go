package main

import (
	//"anymsg/http"
	//"flag"
	"anymsg/cmd"
	"anymsg/http"
	"fmt"

	//"fmt"
	//"github.com/spf13/cobra/cobra/cmd"
	//"os"
	//"github.com/spf13/cobra"
)

const cfgFileName = "cfg.json"


func main() {
	//argsWithProg := os.Args
	//if(len(argsWithProg) < 2) {
	//	fmt.Println("usage : ",argsWithProg[0],"-h")
	//	return
	//}
	//var path string
	//flag.StringVar(&path,"f","/etc/cfg.json","config")
	//flag.Parse()


	err := cmd.Execute()
	if err !=nil{
		fmt.Println(err)
		return
	}
	Jcfg, err := getConfig(cmd.Config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	http.SrvStart(Jcfg)

	fmt.Println("msg-sender servise runing... ")
}
