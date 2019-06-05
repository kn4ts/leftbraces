package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
//	"errors"
	"os"
)

type Events []*Event

type Event struct {
	Name  string
	//Id    string
	Note  string
//	Begin time.Time
//	End   time.Time
	Tasks []*Task
}

type Task struct {
	Name  string
	//Id    string
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



// イベント追加メソッド
// 引数：イベント名，開始時間（Now or 指定）
func NewEvent(name string) *Event {
	return &Event{
		Name: name,
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
		Name:  name,
		Begin: begin,
		End:   end,
	}
}

// jsonファイルに保存されたイベントを表示する関数
func ListEvents() error {
	// 読み込み用の構造体スライスを宣言
	var t0 time.Time
	const t_fmt = "01/02"

	// jsonファイルを読み込んでelに入れる
	el, err := ReadEvents("./event.json")
	if err != nil {
		return err
	}

	// イベント構造体スライス，タスク構造体スライスごとにfor文を回して中身を表示
	for i, ev := range el {
		evln := fmt.Sprintf("#%-2d  %s", i+1, ev.Name)
		fmt.Println(evln)

		for j, ts := range ev.Tasks {
			tsln := ts.Name
			if ts.Begin.Equal(t0) {
				if ts.End.Equal(t0) {
					// 日付未設定のとき
					tsln = fmt.Sprintf(".%-2d  %-s  [ - - - ]", j+1, tsln)
					//tsln = tsln + " "
				} else {
					// 終了日だけあるとき
					tsln = fmt.Sprintf(".%-2d  %-s  [ - %s ]", j+1, tsln, ts.End.Format(t_fmt))
					//tsln = tsln + " [ -" + ts.End.Format(t_fmt) + " ]"
				}
			} else {
				if ts.End.Equal(t0) {
					// 開始日だけあるとき
					tsln = fmt.Sprintf(".%-2d  %-s  [ %s - ]", j+1, tsln, ts.Begin.Format(t_fmt))
					//tsln = tsln + " [ " + ts.Begin.Format(t_fmt) + "- ]"
				} else if ts.End.Equal(ts.Begin) {
					// 開始日と終了日が等しいとき
					tsln = fmt.Sprintf(".%-2d  %-s  [ %s ]", j+1, tsln, ts.Begin.Format(t_fmt))
					//tsln = tsln + " [ " + ts.Begin.Format(t_fmt) + " ]"
				} else {
					// 開始日と終了日が別日に設定されているとき
					tsln = fmt.Sprintf(".%-2d  %-s  [ %s - %s ]", j+1, tsln, ts.Begin.Format(t_fmt), ts.End.Format(t_fmt))
					//tsln = tsln + " [ " + ts.Begin.Format(t_fmt) + "-" + ts.End.Format("01/02") + " ]"
				}
			}
			fmt.Println("  +", tsln)
		}
	}
	return err
}

func ParseDate(ts string) (dat time.Time, err error) {
	// 読み取る日付のパターンを定義
	const tfmt_l = "2006/01/02"
	const tfmt_s = "0102"

	tn := time.Now()
	//var dat time.Time
	//dat = time.Date(0,0,0,0,0,0,0,time.UTC)
	// 引数を日付としてパース
	if len(tfmt_l) == len(ts) {
		dat, err = time.Parse(tfmt_l, ts)
		if err != nil {
			return dat, err
		}
	} else if len(tfmt_s) == len(ts) {
		dat, err = time.Parse(tfmt_s, ts)
		if err != nil {
			return dat, err
		}
		dat = time.Date(tn.Year(), dat.Month(), dat.Day(), 0, 0, 0, 0, time.UTC)
	}
	return dat, nil
}

// jsonファイルを読み込みEvents構造体に内容を転写
func ReadEvents(fname string) (evs Events, err error) {
	// jsonファイルの読み込み
	var raw []byte
	raw, err = ioutil.ReadFile(fname)
	if err != nil {
		return evs, err
	}
	
	// 読み込んだjsonファイルを整列してeventsに入れる
	err = json.Unmarshal(raw, &evs)
	if err != nil {
		return evs, err
	}
	return evs, err
}

// Events構造体の内容をjsonファイルに保存
func SaveEvents(evs Events, fname string) (err error) {
	var wr []byte
	//var fp *File
	wr, err = json.MarshalIndent(evs, "", "  ")
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	fp.Write(wr)
	return err
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
