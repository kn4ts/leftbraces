package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"strconv"
	"os"
	"strings"
//	"time"
)

var fname = "./event.json"

func main() {
	ef := Exists(fname)
	if ef != true {
		err := initJson(fname)
		if err != true {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	// 引数なしだと終了
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s list\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "[usage] %s new", os.Args[0])
		os.Exit(1)
	}

	var subcmd = os.Args[1]
	var subargs = os.Args[2:]
	const Nmax = 100

	switch subcmd {
	case "list":
		// jsonファイルを読み込んでelに入れる
		el, err := ReadEvents(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// jsonファイルの内容を一覧表示
		err = ListEvents(el)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "new":
		// 部分引数の数をチェック
		c = len(subargs) - 1
		if c < 1 {
			fmt.Fprintf(os.Stderr, "[usage] %s new EventName/TaskName Date\n", os.Args[0])
			fmt.Fprintf(os.Stderr, "[usage] %s new TaskName Date", os.Args[0])
			os.Exit(1)
		}
		var evn string
		var tsn string
		// 第一引数が"/"を含むかを判定
		if strings.Contains(subargs[0], "/") {
			// 第一引数の長さをチェック
			if len(subargs[0]) > 2*Nmax {
				fmt.Fprintln(os.Stderr, "Too long EventName/TaskName")
				os.Exit(1)
			}
			// "/"でイベント名とタスク名を分割
			slice := strings.Split(subargs[0], "/")
			evn = slice[0]
			tsn = slice[1]
		} else {
			// イベント名なし，タスク名のみを設定
			evn = ""
			tsn = subargs[0]
		}
		// イベント名の長さをチェック
		if len(evn) > Nmax {
			fmt.Fprintln(os.Stderr, "Too long EventName")
			os.Exit(1)
		}
		// タスク名の長さをチェック
		if len(tsn) > Nmax {
			fmt.Fprintln(os.Stderr, "Too long TaskName")
			os.Exit(1)
		}

		bdat, edat, err := genBeginEnd(subargs[1])
		//fmt.Println(bdat)
		//fmt.Println(edat)
		// 保存されているイベントをEvents構造体に読み込む
		events, err := ReadEvents(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		ev1 := NewEvent(evn) // イベントを新規作成
		// fmt.Printf("%#v", ev1)
		ev1.AddTask(NewTask(tsn, bdat, edat)) // イベントにタスクを追加
		events.AddEvent(ev1)                  // イベントリストに追加

		// イベントをjsonへ保存する
		err = SaveEvents(events, fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "rm":
		c = len(subargs) - 1
		if c < 0 {
			fmt.Fprintf(os.Stderr, "[usage] %s rm EventNum.TaskNum\n", os.Args[0])
			fmt.Fprintf(os.Stderr, "[usage] %s rm EventNum.0", os.Args[0])
			os.Exit(1)
		}
		var enum int
		var tnum int
		var err error
		if strings.Contains(subargs[0], ".") {
			// 引数の長さをチェック
			if len(subargs[0]) > 5 {
				fmt.Fprintln(os.Stderr, "too long args")
				os.Exit(1)
			}
			// "."でイベント番号とタスク番号を分割
			slice := strings.Split(subargs[0], ".")
			enum, err = strconv.Atoi(slice[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tnum, err = strconv.Atoi(slice[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("command 'rm' needs commma .")
			fmt.Fprintf(os.Stderr, "[usage] %s rm EventNum.TaskNum\n", os.Args[0])
			fmt.Fprintf(os.Stderr, "[usage] %s rm EventNum.0", os.Args[0])
			os.Exit(1)
		}

		// 保存されているイベントをEvents構造体に読み込む
		events, err := ReadEvents(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b := events.RemoveItem(enum,tnum)
		if b == false {
			fmt.Println(b)
			os.Exit(1)
		}

		//err = ListEvents(events);
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		// イベントをjsonへ保存する
		err = SaveEvents(events, fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "add":
		// 部分引数の数をチェック
		c = len(subargs) - 1
		if c < 1 {
			fmt.Fprintf(os.Stderr, "[usage] %s add EventNum/TaskName Date\n", os.Args[0])
			//fmt.Fprintf(os.Stderr, "[usage] %s add TaskName Date", os.Args[0])
			os.Exit(1)
		}
		var evnum int
		var tsn string
		var err error
		// 第一引数が"/"を含むかを判定
		if strings.Contains(subargs[0], "/") {
			// 第一引数の長さをチェック
			if len(subargs[0]) > 2*Nmax {
				fmt.Fprintln(os.Stderr, "Too long EventNum/TaskName")
				os.Exit(1)
			}
			// "/"でイベント名とタスク名を分割
			slice := strings.Split(subargs[0], "/")
			evnum, err = strconv.Atoi(slice[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tsn = slice[1]
		} else {
			fmt.Println("Separater '/' is needed")
			os.Exit(1)
		}

		// 保存されているイベントをEvents構造体に読み込む
		events, err := ReadEvents(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		// 日付を抽出
		bdat, edat, err := genBeginEnd(subargs[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 境界値判定
		if evnum < 1 || evnum > len(events) {
			fmt.Println("invalid event id ")
			os.Exit(1)
		}
		events[evnum-1].AddTask(NewTask(tsn, bdat, edat))

		// イベントをjsonへ保存する
		err = SaveEvents(events, fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "done":
		c = len(subargs) - 1
		if c < 0 {
			fmt.Fprintf(os.Stderr, "[usage] %s done EventNum.TaskNum", os.Args[0])
			//fmt.Fprintf(os.Stderr, "[usage] %s done EventNum.0", os.Args[0])
			os.Exit(1)
		}
		var enum int
		var tnum int
		var err error
		if strings.Contains(subargs[0], ".") {
			// 引数の長さをチェック
			if len(subargs[0]) > 5 {
				fmt.Fprintln(os.Stderr, "too long args")
				os.Exit(1)
			}
			// "."でイベント番号とタスク番号を分割
			slice := strings.Split(subargs[0], ".")
			enum, err = strconv.Atoi(slice[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tnum, err = strconv.Atoi(slice[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("command 'done' needs commma . to separate EventNum and TaskNum")
			fmt.Fprintf(os.Stderr, "[usage] %s done EventNum.TaskNum", os.Args[0])
			//fmt.Fprintf(os.Stderr, "[usage] %s done EventNum.0", os.Args[0])
			os.Exit(1)
		}

		// 保存されているイベントをEvents構造体に読み込む
		events, err := ReadEvents(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b := events.DoneItem(enum,tnum)
		if b == false {
			fmt.Println(b)
			os.Exit(1)
		}

		// イベントをjsonへ保存する
		err = SaveEvents(events, fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("no such a command")
		os.Exit(1)
	}
}
