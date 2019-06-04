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
		//case "add":
		//	c = len(subargs) -1
		//	if c < 1 {
		//		fmt.Fprintln(os.Stderr, "[usage] %s add TaskName", os.Args[0])
		//		fmt.Fprintln(os.Stderr, "[usage] %s add TaskName Time", os.Args[0])
		//		os.Exit(1)
		//	}
		//	ev1 := NewEvent(subargs[0], time.Now()) // イベント"A"を新規作成
		//	// fmt.Printf("%#v", ev1)
		//	events.AddEvent(ev1) // "A"をイベントリストに追加
		//	ev1.AddTask(NewTask("task A", time.Now())) // "A"にタスク"A"(T-A)を追加
	}

}
