package main

//利用するパッケージの宣言
import (
	"net/http" //標準パッケージ
	"time"

	//外部パッケージ
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//
)

const tmplPath = "src/template/"

//グローバル変数eにcreateMux()の関数の戻り値を格納
var e = createMux()

func main() {
	//`/`と言うパス(URL)と``articleindex`を結びつける
	e.GET("/", articleindex)

	//wevサーバーをローカルホストで起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New() //アプリケーションインスタンスを生成

	//アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e //アプリケーションインスタンスを返却
}

func articleindex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello, World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
