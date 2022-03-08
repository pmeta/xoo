package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// 项目元信息
var (
	webpage  = "https://xoo.site/"
	app      = "xoo"
	commitId string
	version  string
	buildAt  string
	branch   string
	platform = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	webListenAddr = kingpin.Flag("meta.listen-addr", "监听地址").Default(":8080").String()
	webPrefix     = kingpin.Flag("meta.web-prefix", "路由地址前缀").Default("/xoo/").String()
	//go:embed static
	static embed.FS
	//go:embed html/*.html
	index embed.FS
)

func init() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	kingpin.Parse()
}

func main() {
	router := gin.Default()
	html, _ := template.ParseFS(index, "html/*.html")
	router.SetHTMLTemplate(html)
	router.Any(path.Join(*webPrefix, ""), func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"staticRoot": path.Join(*webPrefix, "static"),
			"title":      "秀哦工作室",
		})
	})
	router.StaticFS(path.Join(*webPrefix, ""), http.FS(static))
	if err := router.Run(*webListenAddr); err != nil {
		os.Exit(1)
	}
}
