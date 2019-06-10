# leftbraces

タスク記録・管理ツール  

## Windows環境でのビルド  

	$ go build -o lb.exe


## 仕様  

データ構造（jsonで管理）

	1.Event1
	  |-- 1:Task1 (Date1, Status1)
	  |-- 2:Task2 (Date2, Status2)
	  |-- ...
	  ~

	2.Event2
	  |-- 1:Task3 (Date3, Status3)

	  ~

	3.
	  |--- 1:Task4 (Date4, Status4)

ここで，各行先頭の`Event1`, `Task1` などはイベント名，タスク名を表す．
タスクは日付（`Date`，日単位）と状態（`Status`，`0:未完, 1:完了`の二値）を有する．

上のデータ構造に示されるように，原則としてタスクはいずれかのイベントに属する．タスクを単体で生成することも可能である（ようにする）が，その場合も`無名のイベント`に内包される．

## コマンド

### イベント・タスクの一覧表示（`lb list`）
	$ lb list	// 上のデータ構造のようなものを表示する

### イベント，タスクの新規生成（`lb new`）
	$ lb new [EventName]/[TaskName] [Date]
	$ lb new [TaskName] [Date]

ここで`[Date]`は以下のいずれかの形式とする．

	2019/06/10-2019/06/22	// `-`の前後で開始，終了年月日を指定
	0610-0622		// 上の例の日付は月日の4桁に省略可能
	0610-			// `-`を後ろに付けると終了日を指定しない
	-0622			// `-`を前に付けると開始日を指定しない
	0610			// `-`を付けない単一の日付はその日のみ（開始日=終了日）
	0			// 日付指定なし

また`[EventName]`，`[TaskName]`には文字数制限（50文字）がある．

### イベント，タスクの削除（`lb rm`）
`lb list`で表示されるイベント番号とタスク番号を指定して削除する  
（イベント番号とタスク番号は固有の識別子でないことに注意）

	$ lb rm [EventNum].[TaskNum]	// 指定タスクの削除
	$ lb rm [EventNum].0		// 指定イベントの削除（含まれるタスクも消える）

### 既存のイベントにタスクを追加（`lb add`）
`lb list`で表示されるイベント番号を指定してその中に新規タスクを追加する  
（イベント番号は固有の識別子でないことに注意）

	$ lb add [EventNum]/[TaskName] [Date]  // イベント番号を指定してその中にタスクを追加する

### 既存のタスクの日付を変更（上書き）（`lb mod`）
	$ lb mod [EventNum].[TaskNum] [Date]

### タスク完了（`lb done`）
	$ lb done [EventNum].[TaskNum]
