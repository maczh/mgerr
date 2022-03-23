package main

import (
	"github.com/maczh/logs"
	"mgerr"
)

func main() {
	mgerr.LoadErrorCodeMessage("zh-cn","example.zh-cn.json")
	mgerr.LoadErrorCodeMessage("en-us","example.en-us.json")
	logs.Debug("中文:{}",mgerr.ErrorLangResult(1001,"zh-cn"))
	logs.Debug("英文:{}",mgerr.ErrorLangResult(1004,"en-us"))
}
