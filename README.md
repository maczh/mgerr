# 通用多国语言错误代码mgerr

mgerr是一个通过配置文件管理同一错误代码多国语言错误信息的错误框架。系统默认定义了zh-cn, zh-tw, en-us, fr, it, de, ja, ko等常用语种。常用的通用错误代码在errcode包中定义。

## 一、安装

```powershell
go get -u github.com/maczh/mgerr
```

## 二、使用

### 2.1 自定义错误代码

工程中自定义错误代码与相应的各语种错误信息的json文件，如工程名为myproject，错误文件放在/opt/go/bin/errcodes/目录下，简体中文错误码定义文件 myproject.errcode.zh-cn.json

```json
{
  "2001": "文件不存在",
  "2002": "请求参数不可为空",
  "2003": "数据库连接失败"
}
```

英文错误码定义文件 myproject.errorcode.en-us.json

```json
{
  "2001": "File not exists",
  "2002": "Request parameter is null or empty",
  "2003": "Mysql connect failed"
}
```

### 2.2 在工程的配置yml文件中添加配置

```yaml
go:
  i8n:
    languages: zh-cn,en-us
    json:
      zh-cn: /opt/go/bin/errcodes/myproject.errcode.zh-cn.json
      en-us: /opt/go/bin/errcodes/myproject.errcode.en-us.json
```

### 2.3 在mgin工程的main中初始化错误代码配置

在main.go中的mgconfig.InitConfig()之后，添加一行：

```go
mgerr.Init()
```

在Router.go中添加一个middleware引用，用于自动处理http接口header中传入的X-Lang定义

```go
	//添加多语种支持
	engine.Use(mgerr.RequestLanguage())
```

### 2.4 修改返回mgresult.Error()

将原来需要用mgresult.Error(status,msg)部分进行如下修改：

#### 2.4.1 使用标准错误信息返回

```go
//code即为status错误码
mgerr.ErrorResult(2002)
```

#### 2.4.2 在标准错误信息之后加上额外错误信息返回

返回的错误信息为 <错误信息>:<附加信息>

```go
mgerr.ErrorResultWithMsg(2003,err.Error())
```

### 2.5 在http接口请求中控制语种

在外部请求时，添加 http header, X-Lang，值为语种定义代码，即可获得相应语种的错误返回信息，若无此参数默认为zh-cn

如：

```
//英文
X-Lang: en-us
//简体中文
X-Lang: zh-cn
//日文
X-Lang: ja
```

### 2.6 其他函数

```go
//根据错误码返回指定的语言的错误信息
func Error(code int, lang string) (int, string)
//根据错误码与指定语言，返回带附加信息的错误信息
func ErrorWithMsg(code int, lang, msg string) (int, string)
//根据错误码返回当前语种的error
func LocalError(code int) error
//取当前请求的语言代码
func GetCurrentLanguage() string
```

### 2.7 系统自带标准错误码

```go
const (
	URI_NOT_FOUND          = 1000
	SYSTEM_ERROR           = 1001
	DB_CONNECT_ERROR       = 1002
	REQUEST_PARAMETER_LOST = 1003
	DATA_NOT_FOUND         = 1004
	USER_NOT_FOUND         = 1005
	PASSWORD_ERROR         = 1006
	VERIFY_CODE_ERROR      = 1007
	TOKEN_ERROR            = 1008
	AUTHENTICATION_FAILURE = 1009
	SERVICE_UNAVAILABLE    = 1010
)
```

