package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"os"
	"path"
)

func isExist(path string) bool { //copy from  phpgo's csdnBlog
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func getConfig(fn string) (*simplejson.Json, error) {
	var fp string
	wd, err := os.Getwd()
	if err == nil {
		fp = path.Join(wd, fn)
		fmt.Println("path: ", fp)
	} else {
		panic(err)
	}
	if !isExist(fp) {
		fmt.Println("error: the configfile is not exist")
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("error:file read error")
		panic(err)
	} else {
		return simplejson.NewJson(data)
	}

}
