package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	var events Events
	ev1 := NewEvent("test event A", time.Now())
	// fmt.Printf("%#v", ev1)
	events.AddEvent(ev1)

	ev3 := NewEvent("test event B", time.Now())
	events.AddEvent(ev3)

	ev1.AddTask(NewTask("task A", time.Now()))
	// fmt.Printf("%#v", ev1)

	ev1.AddTask(NewTask("task B", time.Now()))
	// fmt.Printf("%#v", ev1)

	ev1.Tasks[0].AddItem(NewItem("item A"))
	// fmt.Printf("%#v", ev1)
	ev1.Tasks[0].AddItem(NewItem("item B"))
	fmt.Printf("%#v", ev1)
	fmt.Printf("%#v", ev1.Tasks[0])

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
