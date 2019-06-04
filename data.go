package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Events []*Event

type Event struct {
	Name  string
	Id    string
	Note  string
	Begin time.Time
	End   time.Time
	Tasks []*Task
}

type Task struct {
	Name  string
	Id    string
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
	Name  string    `json:"Name"`
	Id    string    `json:"Id"`
	Note  string    `json:"Note"`
	Begin time.Time `json:"Begin"`
	End   time.Time `json:"End"`
	Tasks []TaskR   `json:"Tasks"`
}

type TaskR struct {
	Name  string    `json:"Name"`
	Id    string    `json:"Id"`
	Note  string    `json:"Note"`
	Begin time.Time `json:"Begin"`
	End   time.Time `json:"End"`
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
func NewEvent(name string) *Event {
	return &Event{
		Name:  name,
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
func NewTask(name string, begin time.Time, end time.Time) *Task {
	return &Task{
		Name: name,
		Begin: begin,
		End:  end,
	}
}

// jsonファイルに保存されたイベントを表示する関数
func ListEvents() error {
	// jsonファイルの読み込み
	raw, err := ioutil.ReadFile("./event.json")
	if err != nil {
		return err
	}

	// 読み込み用の構造体スライスを宣言
	var el EventsR
	var t0 time.Time
	//var evln []string
	//var tsln []string
	//var flag1 = 0

	// 読み込んだjsonファイルを整列してelに入れる
	err = json.Unmarshal(raw, &el)
	if err != nil {
		return err
	}

	// イベント構造体スライス，タスク構造体スライスごとにfor文を回して中身を表示
	for _, ev := range el {
		evln := "# " + ev.Name
		fmt.Println(evln)

		for _, ts := range ev.Tasks {
			tsln := ts.Name
			if ts.Begin.Equal(t0) {
				if ts.End.Equal(t0) {
					// 日付未設定のとき
					tsln = tsln + " [ - - - ]"
				} else {
					// 終了日だけあるとき
					tsln = tsln + " [ - " + ts.End.Format("01/02") + " ]"
				}
			} else {
				if ts.End.Equal(t0) {
					// 開始日だけあるとき
					tsln = tsln + " [ " + ts.Begin.Format("01/02") + " - ]"
				} else if ts.End.Equal(ts.Begin) {
					// 開始日と終了日が等しいとき
					tsln = tsln + " [ " + ts.Begin.Format("01/02") + " ]"
				} else {
					// 開始日と終了日が別日に設定されているとき
					tsln = tsln + " [ " + ts.Begin.Format("01/02") + " - " + ts.End.Format("01/02") + " ]"
				}
			}
			fmt.Println("  +", tsln)
		}
	}

	return nil
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
