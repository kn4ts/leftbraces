package main

import "time"

type Events []*Event

type Event struct {
	Name  string
	Note  string
	Begin time.Time
	End time.Time
	Tasks []*Task
}

type Task struct {
	Name  string
	Note  string
	Begin time.Time
	End   time.Time
	//Items []*Item
}

//type Item struct {
//	Name string
//	Note string
//	Done bool
//}

type EventsR []EventR

type EventR struct {
	Name  string  `json:"Name"`
	Note  string  `json:"Note"`
	Begin time.Time  `json:"Begin"`
	End   time.Time  `json:"End"`
	Tasks []TaskR  `json:"Tasks"`
}

type TaskR struct {
	Name  string  `json:"Name"`
	Note  string  `json:"Note"`
	Begin time.Time  `json:"Begin"`
	End   time.Time  `json:"End"`
	//Items []*Item
}

//type ItemR struct {
//	Name string
//	Note string
//	Done bool
//}
// func ShowEvents

// イベント追加メソッド
// 引数：イベント名，開始時間（Now or 指定）
func NewEvent(name string, begin time.Time) *Event {
	return &Event{
		Name:  name,
		Begin: begin,
	}
}

// イベント追加メソッド
func (es *Events) AddEvent(e *Event) {
	*es = append(*es, e)
}

// タスク追加メソッド
func (e *Event) AddTask(t *Task) {
	e.Tasks = append(e.Tasks, t)
}

// 新規タスク生成メソッド
func NewTask(name string, end time.Time) *Task {
	return &Task{
		Name: name,
		End:  end,
	}
}

// アイテム追加メソッド
//func (t *Task) AddItem(i *Item) {
//	t.Items = append(t.Items, i)
//}

// 新規アイテム生成メソッド
//func NewItem(name string) *Item {
//	return &Item{
//		Name: name,
//	}
//}
