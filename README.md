# miraiLABO-blog
技術ブログGo言語で制作<br>
https://protected-shore-69482.herokuapp.com/

#### Basic認証パス
SteinsGate<br>
ElPsyCongroo

## 機能
- マークダウン記法
- 新規作成・消去・編集機能
- Basic認証
- DBで管理
## 使い方
記事の作成・編集・保存・消去を利用するにはBasic認証を通す必要があります。
閲覧は必要ありません。

#### Basic認証パス
user: SteinsGate<br>
pass: ElPsyCongroo<br>

### 一覧画面
`新規`で空の記事編集画面、`ゴミ箱`で消去、`もっとみる`で表示を５項目づつ増やす
<br>上記の操作にBasic認証を通る必要があります。
<img width="1431" alt="スクリーンショット 2021-09-22 17 23 02" src="https://user-images.githubusercontent.com/80031888/134309926-80587d7c-35b5-47ab-bf8b-ae0d1edc608b.png">

<img width="1428" alt="スクリーンショット 2021-09-22 17 23 06" src="https://user-images.githubusercontent.com/80031888/134309950-808c7412-22f0-48c3-922a-ef90b652d0c9.png">

### 記事表示画面
`編集`で編集画面
<br>上記の操作にBasic認証を通る必要があります。
<img width="1430" alt="スクリーンショット 2021-09-22 17 24 35" src="https://user-images.githubusercontent.com/80031888/134309997-7545851d-1677-43b1-bcab-85d9a9a9ebcf.png">

### 記事編集画面
`保存`で上書き、`キャンセル`で前の画面に戻る、`目のマーク`で編集モードとプレビューモードを切り替え
#### 編集モード
<img width="1432" alt="スクリーンショット 2021-09-22 17 24 44" src="https://user-images.githubusercontent.com/80031888/134310017-9ff4e6fd-2813-40d9-bb2f-bf356e99a258.png">

#### プレビューモード

<img width="1434" alt="スクリーンショット 2021-09-22 17 24 47" src="https://user-images.githubusercontent.com/80031888/134310039-fa717f8b-f1e3-4b95-ac98-3c3f053781d8.png">

## 開発方法
### 開発環境やツール
- mac
- vscode
- Sequel Pro（DB GUI クライアント）
- heroku
- Go
- mysql
- envrc

### 利用したライブラリ
#### echo
軽量なWebフレームワーク
https://echo.labstack.com/

#### pongo2
テンプレートエンジン
https://github.com/flosch/pongo2

#### sqlx
DB操作
https://github.com/jmoiron/sqlx

#### goose
DBマイグレーションツール
https://bitbucket.org/liamstask/goose/src/master/

#### Remarkable 
マークダウン記法を解釈する
https://github.com/jonschlinkert/remarkable

## 説明
### ディレクトリ構造
```bash
.
├── README.md
├── db
│   ├── dbconf.yml
│   └── migrations
│       ├── 20210728204429_create_articles.sql
│       └── 20210807135910_alter_articles_add_columns.sql
├── go.mod
├── go.sum
├── handler
│   ├── article_handler.go
│   └── render.go
├── main.go
├── model
│   └── article.go
├── repository
│   ├── article_repository.go
│   └── repository.go
├── runner.conf
├── src
│   ├── css
│   │   ├── article
│   │   │   ├── _article.css
│   │   │   ├── _form.css
│   │   │   ├── edit.css
│   │   │   ├── index.css
│   │   │   ├── new.css
│   │   │   └── show.css
│   │   ├── layout.css
│   │   └── vendor
│   │       └── prism.css
│   ├── js
│   │   ├── article
│   │   │   ├── edit.js
│   │   │   ├── index.js
│   │   │   ├── new.js
│   │   │   └── show.js
│   │   ├── module
│   │   │   ├── form.js
│   │   │   └── parser.js
│   │   └── vendor
│   │       └── prism.js
│   └── template
│       ├── article
│       │   ├── _form.html
│       │   ├── edit.html
│       │   ├── index.html
│       │   ├── new.html
│       │   └── show.html
│       └── layout.html
└── tmp
    └── runner-build
```
## DB
![article_table](https://user-images.githubusercontent.com/80031888/137452917-1f9a250a-26fc-4338-9d2a-66c3ecfd8638.png)


## main.go
```go
package main

//利用するパッケージの宣言
import 
	"log"
	"os"

	"miraiLABO-blog/handler"
	"miraiLABO-blog/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

//グローバル変数eにcreateMux()の関数の戻り値を格納
var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	//ルーディングのグループ
	auth := e.Group("")

	auth.Use(basicAuth())

	// TOPページに一覧を表示
	e.GET("/", handler.Articleindex)

	e.GET("/articles", handler.Articleindex) //一覧画面
	//e.GET("/articles/new", handler.ArticleNew)              // 新規作成画面
	auth.GET("/articles/new", handler.ArticleNew)      // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow) // 詳細画面
	//e.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面
	auth.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

	e.GET("/api/articles", handler.ArticleList) // 一覧
	//e.POST("/api/articles", handler.ArticleCreate)              // 作成
	auth.POST("/api/articles", handler.ArticleCreate) // 作成
	//e.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 消去
	auth.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 消去
	//e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)  // 更新
	auth.PATCH("/api/articles/:articleID", handler.ArticleUpdate) // 更新

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New() //アプリケーションインスタンスを生成

	//アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())
	//e.Use(basicAuth())

	e.Static("/css", "src/css")
	e.Static("js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

// CustomValidator...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func basicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "SteinsGate" && password == "ElPsyCongroo" {
			return true, nil
		}
		return false, nil
	})
}

```
### Webサーバー
ルーディング,ミドルウェア設定,Webサーバー
```go
func main() {
	//ルーディングのグループ
	auth := e.Group("")

	auth.Use(basicAuth())

	// TOPページに一覧を表示
	e.GET("/", handler.Articleindex)

	e.GET("/articles", handler.Articleindex) //一覧画面
	auth.GET("/articles/new", handler.ArticleNew)      // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow) // 詳細画面
	auth.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

	e.GET("/api/articles", handler.ArticleList) // 一覧
	auth.POST("/api/articles", handler.ArticleCreate) // 作成
	auth.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 消去
	auth.PATCH("/api/articles/:articleID", handler.ArticleUpdate) // 更新

	e.Logger.Fatal(e.Start(":8080"))
}
func createMux() *echo.Echo {
	e := echo.New() //アプリケーションインスタンスを生成

	//アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())
	//e.Use(basicAuth())

	e.Static("/css", "src/css")
	e.Static("js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}
```
### DB接続
```go
import (
	os
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
)
func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}
```
### Basic認証
```go
func basicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "SteinsGate" && password == "ElPsyCongroo" {
			return true, nil
		}
		return false, nil
	})
}
```
