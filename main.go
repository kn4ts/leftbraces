package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {

	// 引数なしだと終了
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s list", os.Args[0])
		os.Exit(1)
	}

	var subcmd = os.Args[1]
	var subargs = os.Args[2:]
	const Nmax = 100

	switch subcmd {
	case "list":
		// jsonファイルの内容を一覧表示
		err := ListEvents()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "new":
		// 部分引数の数をチェック
		c = len(subargs) -1
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

		var bgn string
		var end string
		// 第二引数の長さをチェック
		if len(subargs[1]) > Nmax {
			fmt.Fprintln(os.Stderr, "Too long Date")
			os.Exit(1)
		}
		if strings.Contains(subargs[1], "-") {
			// "-"で開始日と終了日を分割
			slice := strings.Split(subargs[1], "-")
			bgn = slice[0]
			end = slice[1]
		} else {
			// 単一の日付のとき
			bgn = subargs[1]
			end = subargs[1]
		}

//		var bdat = time.Time{}
//		var edat = time.Time{}
		// 第二引数を日付としてパース
		bdat, err := time.Parse("20200101", bgn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		edat, err := time.Parse("20200101", end)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// jsonファイルの読み込み
		raw, err := ioutil.ReadFile("./event.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var events Events
		// 読み込んだjsonファイルを整列してeventsに入れる
		err = json.Unmarshal(raw, &events)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ev1 := NewEvent(evn) // イベントを新規作成
		// fmt.Printf("%#v", ev1)
		ev1.AddTask(NewTask(tsn, bdat, edat)) // イベントにタスクを追加
		events.AddEvent(ev1) // イベントリストに追加

		wr, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			fmt.Printf("error:", err)
		}
		fp, err := os.OpenFile("event.json", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fp.Close()
		fp.Write(wr)
	}

}
