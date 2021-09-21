# miraiLABO-blog
技術ブログGo言語で制作
## 機能
- マークダウン記法
- 新規作成・消去・編集機能
- Basic認証
- DBで管理
## 使い方
記事の作成・編集・保存・消去を利用するにはBasic認証を通す必要があります。
閲覧は問題ありません。

Basic認証パス
user
pass

### 一覧画面
### 記事表示画面
### 記事編集画面

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
