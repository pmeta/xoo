package main

import (
	"embed"
	"fmt"
	"net/http"
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
	//go:embed xoo
	uiIndex embed.FS
)

func init() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	kingpin.Parse()
}

func main() {
	router := gin.Default()
	router.StaticFS("/static", http.FS(uiIndex))
	router.Any("", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	if err := router.Run(*webListenAddr); err != nil {
		return
	}
	//if err := http.ListenAndServe(*webListenAddr, http.FileServer(http.FS(uiIndex))); err != nil {
	//	return
	//}
}
