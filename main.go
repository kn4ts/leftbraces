package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	var events Events
	ev1 := NewEvent("test event A", time.Now())	// イベント"A"を新規作成
	// fmt.Printf("%#v", ev1)
	events.AddEvent(ev1)				// "A"をイベントリストに追加

	ev3 := NewEvent("test event B", time.Now())	// イベント"B"を新規作成
	events.AddEvent(ev3)				// "B"をイベントリストに追加

	ev1.AddTask(NewTask("task A", time.Now()))	// "A"にタスク"A"(T-A)を追加
	// fmt.Printf("%#v", ev1)

	ev1.AddTask(NewTask("task B", time.Now()))	// "B"にタスク"B"(T-B)を追加
	// fmt.Printf("%#v", ev1)

	ev1.Tasks[0].AddItem(NewItem("item A"))		// "A"のタスク"T-A"にアイテム"I-A"を追加
	// fmt.Printf("%#v", ev1)
	ev1.Tasks[0].AddItem(NewItem("item B"))		// "A"のタスク"T-A"にアイテム"I-B"を追加
	fmt.Printf("%#v", ev1)
	fmt.Printf("%#v", ev1.Tasks[0])

	var t = time.Time{}
	t, _ = time.Parse("2020-01-01", "2020-02-01")
	ev4 := NewEvent("future event T", t)
	events.AddEvent(ev4)

	b, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		fmt.Printf("error:", err)
	}
	f, err := os.Create("event.json")
	if err != nil {
		fmt.Printf("error:", err)
	}
	defer f.Close()
	f.Write(b)

	var ev2 Event
	err = json.Unmarshal(b, &ev2)
	if err != nil {
		fmt.Printf("error:", err)
	}
	fmt.Printf("%#v", ev2)

}
