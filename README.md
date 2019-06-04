# leftbraces

タスク記録・管理ツール  

## Windows環境でのビルド  

	$ go build -o lb.exe


## 仕様（仮）  

データ構造（jsonで管理）

	Event1 (EventId1)
	  |-- Task1 (TaskId1, Date1, Status1)
	  |-- Task2 (TaskId2, Date2, Status2)
	  |-- ...
	  ~

	Event2 (EventId2)
	  |-- Task3 (TaskId3, Date3, Status3)

	  ~

	Task4 ((EventId3), TaskId4, Date4, Status4)

ここで，各行先頭の`Event1`, `Task1` などはイベント名，タスク名を表し，それぞれ個別の識別子(`*Id`)を持つ．
タスクは識別子に加えて日付（`Date`，日単位）と状態（`Status`，`0:未完, 1:完了`の二値）を有する．

上のデータ構造に示されるように，原則としてタスクはいずれかのイベントに属する．タスクを単体で生成することも可能である（ようにする）が，その場合も`無名のイベント`に内包される形にする（したい）．

## コマンド（仮）
（これから実装したい）コマンド群

### イベント・タスクの一覧表示（`lb list`）
	$ lb list	// 上のデータ構造のようなものを表示する

### イベント，タスクの新規生成（`lb new`）
	$ lb new [EventName]/[TaskName] [Date]
	$ lb new /[TaskName] [Date]

ここで`[Date]`は以下のいずれかの形式とする．

	2019/06/10-2019/06/22	// `-`の前後で開始，終了年月日を指定
	0610-0622		// 上の例の日付は年月の4桁まで省略可能（にする）
	0610-			// `-`を後ろに付けると終了日を指定しない
	-0622			// `-`を前に付けると開始日を指定しない
	0610			// `-`を付けない単一の日付はその日のみ（開始日=終了日）

また`[EventName]`，`[TaskName]`には文字数制限を設ける（？）

### イベント，タスクの削除（`lb rm`）
	$ lb rm [EventId]
	$ lb rm /[TaskId]

### 既存のイベントにタスクを追加（`lb add`）
	$ lb add [EventId]/[TaskName] [Date]

### 既存のタスクの日付を変更，追加（`lb mod`）
	$ lb mod [EventId] [Date]
	$ lb mod /[TaskId] [Date]

### タスク完了（`lb done`）
	$ lb done /[TaskId]

