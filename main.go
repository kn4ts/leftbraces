package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"os"
	//	"time"
)

func main() {

	// 引数なしだと終了
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s list", os.Args[0])
		os.Exit(1)
	}

	var subcmd = os.Args[1]
	//var subargs = os.Args[2:]

	switch subcmd {
	case "list":
		err := ListEvents()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
