package mgerr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgcache"
	"github.com/maczh/mgconfig"
	"github.com/maczh/mgerr/errcode"
	"github.com/maczh/utils"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var errorCodeMessageI8N = make(map[string]map[int]string)
var errorLangs = []string{"zh-cn", "en-us", "zh-tw", "fr", "it", "de", "ja", "ko", "ru"}

const (
	SUCCESS = 1
)

func LoadErrorCodeMessage(lang, errorFile string) error {
	if !utils.StringArrayContains(errorLangs, lang) {
		logs.Error("{}语言代码暂不支持", lang)
		return errors.New("doesn't support " + lang + " language.")
	}
	f, err := os.Open(errorFile)
	if err != nil {
		logs.Error("加载错误代码json文件错误:{}", err.Error())
		return err
	}

	defer f.Close()
	errCodeMsgMap := map[string]string{}
	err = json.NewDecoder(f).Decode(&errCodeMsgMap)
	if err != nil {
		logs.Error("json文件格式错误:{}", err.Error())
		return err
	}
	errMap := errorCodeMessageI8N[lang]
	if errMap == nil || len(errMap) == 0 {
		errMap = map[int]string{}
	}
	for k, v := range errCodeMsgMap {
		code, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			continue
		}
		errMap[int(code)] = v
	}
	errorCodeMessageI8N[lang] = errMap
	return nil
}

func Error(code int, lang string) (int, string) {
	if !utils.StringArrayContains(errorLangs, lang) {
		logs.Error("{}语言代码暂不支持", lang)
		return code, "未知语言"
	}
	errCodeMsgMap, exists := errorCodeMessageI8N[lang]
	if !exists {
		return code, "未知语言"
	}
	msg, ok := errCodeMsgMap[code]
	if !ok {
		return code, "未知错误代码"
	}
	return code, msg
}

func ErrorWithMsg(code int, lang, msg string) (int, string) {
	_, info := Error(code, lang)
	return code, fmt.Sprintf("%s:%s", info, msg)
}

func LocalError(code int) error {
	_, err := Error(code, GetCurrentLanguage())
	return errors.New(err)
}

func ErrorLangResult(status int, lang string) mgresult.Result {
	return mgresult.Error(Error(status, lang))
}

func ErrorResult(status int) mgresult.Result {
	return mgresult.Error(Error(status, GetCurrentLanguage()))
}

func ErrorResultWithMsg(status int, msg string) mgresult.Result {
	return mgresult.Error(ErrorWithMsg(status, GetCurrentLanguage(), msg))
}

func Init() {
	errorCodeMessageI8N["zh-cn"] = errcode.InitZhCN()
	errorCodeMessageI8N["zh-tw"] = errcode.InitZhTw()
	errorCodeMessageI8N["en-us"] = errcode.InitEnUs()
	errorCodeMessageI8N["fr"] = errcode.InitFr()
	errorCodeMessageI8N["de"] = errcode.InitDe()
	errorCodeMessageI8N["it"] = errcode.InitIt()
	errorCodeMessageI8N["ja"] = errcode.InitJa()
	errorCodeMessageI8N["ko"] = errcode.InitKo()
	languages := mgconfig.GetConfigString("go.i8n.languages")
	if languages == "" {
		return
	}
	langs := strings.Split(languages, ",")
	for _, lang := range langs {
		langFile := mgconfig.GetConfigString("go.i8n.json." + lang)
		_ = LoadErrorCodeMessage(lang, langFile)
	}
}

func RequestLanguage() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("X-Lang")
		if lang == "" {
			lang = "zh-cn"
		}
		if !utils.StringArrayContains(errorLangs, lang) {
			lang = "zh-cn"
		}
		routineId := GetGID()
		mgcache.OnGetCache("Lang").Add(routineId, lang, 5*time.Minute)
	}
}

func GetCurrentLanguage() string {
	lang, found := mgcache.OnGetCache("Lang").Value(GetGID())
	if found {
		return lang.(string)
	} else {
		return "zh-cn"
	}
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
