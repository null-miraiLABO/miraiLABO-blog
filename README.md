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

#### Basic認証パス
user: SteinsGate<br>
pass: ElPsyCongroo<br>

### 一覧画面
<img width="1431" alt="スクリーンショット 2021-09-22 17 23 02" src="https://user-images.githubusercontent.com/80031888/134309926-80587d7c-35b5-47ab-bf8b-ae0d1edc608b.png">

<img width="1428" alt="スクリーンショット 2021-09-22 17 23 06" src="https://user-images.githubusercontent.com/80031888/134309950-808c7412-22f0-48c3-922a-ef90b652d0c9.png">

### 記事表示画面
<img width="1430" alt="スクリーンショット 2021-09-22 17 24 35" src="https://user-images.githubusercontent.com/80031888/134309997-7545851d-1677-43b1-bcab-85d9a9a9ebcf.png">

### 記事編集画面
<img width="1432" alt="スクリーンショット 2021-09-22 17 24 44" src="https://user-images.githubusercontent.com/80031888/134310017-9ff4e6fd-2813-40d9-bb2f-bf356e99a258.png">

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
