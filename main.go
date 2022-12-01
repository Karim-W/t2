package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	arg := args[1]
	//get substring from last / to end
	lastIndex := strings.LastIndex(arg, "/")
	last := arg[lastIndex+1:]
	dir := arg[:lastIndex]
	//split dir by / and check if each dir is a valid dir
	dirs := strings.Split(dir, "/")
	path := ""
	for _, d := range dirs {
		path += d + "/"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			//create dir
			os.Mkdir(path, 0777)
		}
	}
	//check if last contains .
	if strings.Contains(last, ".") {
		//create file
		f, err := os.Create(arg)
		if err != nil {
			os.Exit(1)
		}
		f.Close()

	} else {
		path += last + "/"
		os.Mkdir(path, 0777)
	}
}
