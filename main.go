package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
//	"time"
)

func main() {

	// jsonファイルの読み込み
	raw, err := ioutil.ReadFile("./event.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// 読み込み用の構造体スライスを宣言
	var el EventsR

	// 読み込んだjsonファイルを整列してelに入れる
	err = json.Unmarshal(raw, &el)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// イベント構造体スライス，タスク構造体スライスごとにfor文を回して中身を表示
	for _, ev := range el {
		fmt.Println(ev.Name)
		for _, ts := range ev.Tasks {
			fmt.Println(ts.Name)
		}
	}
}
