package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"os"
	"time"
)

func main() {
	var events Events
	ev1 := NewEvent("test event A", time.Now()) // イベント"A"を新規作成
	// fmt.Printf("%#v", ev1)
	events.AddEvent(ev1) // "A"をイベントリストに追加

	ev3 := NewEvent("test event B", time.Now()) // イベント"B"を新規作成
	events.AddEvent(ev3)                        // "B"をイベントリストに追加

	ev1.AddTask(NewTask("task A", time.Now())) // "A"にタスク"A"(T-A)を追加
	// fmt.Printf("%#v", ev1)

	ev1.AddTask(NewTask("task B", time.Now())) // "B"にタスク"B"(T-B)を追加
	// fmt.Printf("%#v", ev1)

	ev1.Tasks[0].AddItem(NewItem("item A")) // "A"のタスク"T-A"にアイテム"I-A"を追加
	// fmt.Printf("%#v", ev1)
	ev1.Tasks[0].AddItem(NewItem("item B")) // "A"のタスク"T-A"にアイテム"I-B"を追加
	fmt.Printf("%#v", ev1)
	fmt.Printf("%#v", ev1.Tasks[0])

	var t = time.Time{}
	t, _ = time.Parse("2020-01-01", "2020-02-01")
	ev4 := NewEvent("future event T", t)
	events.AddEvent(ev4)

	//	raw, err := ioutil.ReadFile("./event.json")
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		os.Exit(1)
	//	}
	//	var fc FeatureCollection
	//
	//	json.Unmarshal(raw, &fc)
	//
	//	for _, ft := range fc.Features {
	//		fmt.Println(ft.Type)
	//		fmt.Println(ft.Properties.Pref, ft.Properties.City1, ft.Properties.City2)
	//		fmt.Println(ft.Geometory.Type)
	//		fmt.Println(ft.Geometory.Coordinates[0][0][0], ft.Geometory.Coordinates[0][0][1])
	//	}

	b, err := json.MarshalIndent(events, "", "  ") // インデント整列
	if err != nil {
		fmt.Printf("error:", err)
	}
	f, err := os.Create("event.json") // jsonファイルを作成
	if err != nil {
		fmt.Printf("error:", err)
	}
	defer f.Close()
	f.Write(b) // 整列させた内容をjsonに書き込み

	var eventss Events
	err = json.Unmarshal(b, &eventss)
	if err != nil {
		fmt.Printf("error:", err)
	}
	//fmt.Printf("%#v", eventss)
	fmt.Println("%v", eventss)

}
