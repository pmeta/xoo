package main

import (
	"fmt"
	"runtime"

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

	webListenAddr = kingpin.Flag("meta.listen-addr", "监听地址").Default(":80").String()
)

func init() {

}

func main() {

}
