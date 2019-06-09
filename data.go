package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"strings"
	//	"errors"
	"os"
)

type Events []*Event

type Event struct {
	Name string
	//Id    string
	Note string
	//	Begin time.Time
	//	End   time.Time
	Tasks []*Task
}

type Task struct {
	Name string
	//Id    string
	Note  string
	Begin time.Time
	End   time.Time
	Done bool
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

// removeEventはイベントを削除する。
// このメソッドは境界をチェックしないので直接呼び出さない。
func (el *Events) removeEvent(i int) bool {
	l := len(*el)
	*el = append((*el)[0:i], (*el)[i+1:l]...)
	return true
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

// removeTaskはタスクを削除する。
// このメソッドは境界をチェックしないので直接呼び出さない。
func (e *Event) removeTask(i int) bool {
	l := len(e.Tasks)
	e.Tasks = append((e.Tasks)[0:i], (e.Tasks)[i+1:l]...)
	return true
}

// RemoveItemはタスクまたはサブタスクを削除する。
// removeTask, removeSubtaskは境界をチェックしない安全でないメソッドなので
// これらを直接呼ばずにRemoveItemを使うこと。
func (el *Events) RemoveItem(mainNum, subNum int) bool {
	if subNum == 0 {
		return el.removeEvent(mainNum - 1)
	}
	if mainNum < 1 || mainNum > len(*el) {
		return false
	}
	return (*el)[mainNum-1].removeTask(subNum - 1)
}

// jsonファイルに保存されたイベントを表示する関数
func ListEvents(el Events) (err error) {
	// 読み込み用の構造体スライスを宣言
	var t0 time.Time
	const t_fmt = "01/02"


	// イベント構造体スライス，タスク構造体スライスごとにfor文を回して中身を表示
	for i, ev := range el {
		evln := fmt.Sprintf("%2d. %s", i+1, ev.Name)
		fmt.Println(evln)

		for j, ts := range ev.Tasks {
			var stat string
			if ts.Done==true {
				stat = "[Done]"
			} else {
				if ts.Begin.Equal(t0) {
					if ts.End.Equal(t0) {
						// 日付未設定のとき
						stat = "[ - ]"
					} else {
						// 終了日だけあるとき
						stat = fmt.Sprintf("[-%s]", ts.End.Format(t_fmt))
					}
				} else {
					if ts.End.Equal(t0) {
						// 開始日だけあるとき
						stat = fmt.Sprintf("[%s-]", ts.Begin.Format(t_fmt))
					} else if ts.End.Equal(ts.Begin) {
						// 開始日と終了日が等しいとき
						stat = fmt.Sprintf("[%s]", ts.Begin.Format(t_fmt))
						//tsln = tsln + " [ " + ts.Begin.Format(t_fmt) + " ]"
					} else {
						// 開始日と終了日が別日に設定されているとき
						stat = fmt.Sprintf("[%s-%s]", ts.Begin.Format(t_fmt), ts.End.Format(t_fmt))
						//tsln = fmt.Sprintf("%2d: [%s-%s] %-s", j+1, ts.Begin.Format(t_fmt), ts.End.Format(t_fmt), tsln)
					}
				}
			}
			tsln := fmt.Sprintf("  %2d: %s %s", j+1, stat, ts.Name)
			fmt.Println(tsln)
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

func genBeginEnd(st string) (bt time.Time, et time.Time, err error) {
	const Nmax = 30
	var bgn string
	var end string
	// 引数の長さをチェック
	if len(st) > Nmax {
		//fmt.Fprintln(os.Stderr, "Invalid Date (too long)")
		return bt, et, err
	}
	// "-"を含むか判定
	if strings.Contains(st, "-") {
		// "-"で開始日と終了日を分割
		slice := strings.Split(st, "-")
		bgn = slice[0]
		end = slice[1]
	} else {
		// 単一の日付のとき
		bgn = st
		end = st
	}
	
	// 開始日と終了日をパースして時間型に変換
	bt, err = ParseDate(bgn)
	if err != nil {
		//fmt.Println("invalid begin date")
		return bt, et, err
	}
	//fmt.Println(bdat)

	et, err = ParseDate(end)
	if err != nil {
		//fmt.Println("invalid end date")
		//os.Exit(1)
		return bt, et, err
	}
	return bt, et, err
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
	//fp, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	return err
	//}
	fp, err := os.Create(fname)
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
