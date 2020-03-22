package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var StatusText map[int]string = StatusTextEn

func JsonResult(c *gin.Context, code int, data interface{}, msg string)  {
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"data": data,
		"msg": msg,
	})
}

const (
	StatusOK                   = 200
	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6

	//下面是自定义错误------------------------------------------------------------------------------------------------

	StatusDbError                       = 600	//数据库类错误
	StatusDbConnectionError             = 601	//数据库连接失败
	StatusDbPingError                   = 602	//数据库连接后Ping失败
	StatusNotInitDb                     = 603	//没有初始化连接数据库
	StatusShouldBindError               = 604	//请求数据绑定错误
	StatusDecodeError                   = 605	//数据库数据绑定错误
	StatusIsExist                       = 606	//向数据库增加数据重复
	StatusNotExist                      = 607	//数据库未查到数据
	StatusLoginError                    = 608	//登录失败
	StatusCreateTokenError              = 610	//创建token失败
	StatusDbFindError                   = 611	//数据库查找错误
	StatusDbInsertError                 = 612	//数据库增加错误
	StatusDbUpdateError                 = 613	//数据库修改错误
	StatusDbDeleteError                 = 614	//数据库删除错误
	StatusListError                     = 615	//获取列表错误

	StatusBsonMarshalError              = 701	//bson序列化失败
	StatusBsonUnmarshalError            = 702	//bson反序列化失败
	StatusJsonMarshalError              = 703	//json序列化失败
	StatusJsonUnmarshalError            = 704	//json反序列化失败
	StatusParseUrlError                 = 705	//url解析失败
	StatusParseTokenError               = 706	//token解析失败

	StatusFileError                     = 800	//文件类错误
	StatusOpenFileError                 = 801	//打开文件错误
	StatusCloseFileError                = 802	//关闭文件错误
	StatusReadFileError					= 803	//读文件错误
	StatusWriteFileError	     		= 804	//写文件错误
	StatusFileNotExist                  = 805	//文件不存在
)

var StatusTextEn = map[int]string{
	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",

	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthRequired:            "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusPreconditionFailed:           "Precondition Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	StatusExpectationFailed:            "Expectation Failed",
	StatusTeapot:                       "I'm a teapot",
	StatusMisdirectedRequest:           "Misdirected Request",
	StatusUnprocessableEntity:          "Unprocessable Entity",
	StatusLocked:                       "Locked",
	StatusFailedDependency:             "Failed Dependency",
	StatusTooEarly:                     "Too Early",
	StatusUpgradeRequired:              "Upgrade Required",
	StatusPreconditionRequired:         "Precondition Required",
	StatusTooManyRequests:              "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",

	//下面是自定义错误------------------------------------------------------------------------------------------------

	StatusDbError:                       "Database Error",
	StatusDbConnectionError:             "Database Connection Error",
	StatusDbPingError:                   "Database Ping Error",
	StatusNotInitDb:                     "Not Init Database",
	StatusShouldBindError:               "ShouldBind Error",
	StatusDecodeError:                   "Decode Error",
	StatusIsExist:                       "Is Exist",
	StatusNotExist:                      "Not Exist",
	StatusLoginError:                    "Login Error",
	StatusCreateTokenError:              "Create Token Error",
	StatusDbFindError:                   "Database Find Error",
	StatusDbInsertError:                 "Database Insert Error",
	StatusDbUpdateError:                 "Database Update Error",
	StatusDbDeleteError:                 "Database Delete Error",
	StatusListError:                     "Get List Error",

	StatusBsonMarshalError:              "Bson Marshal Error",
	StatusBsonUnmarshalError:            "Bson Unmarshal Error",
	StatusJsonMarshalError:              "Json Marshal Error",
	StatusJsonUnmarshalError:            "Json Unmarshal Error",
	StatusParseUrlError:                 "Parse Url Error",
	StatusParseTokenError:               "Parse Token Error",

	StatusFileError:                     "File Error",
	StatusOpenFileError:                 "Open File Error",
	StatusCloseFileError:                "Close File Error",
	StatusReadFileError:				 "Read File Error",
	StatusWriteFileError:	     		 "Write File Error",
	StatusFileNotExist:                  "File Not Exist",
}

var StatusTextZh = map[int]string{
	StatusOK:                   "成功",
	StatusCreated:              "成功创建",
	StatusNoContent:            "无信息",
	StatusResetContent:         "重置视图",

	StatusBadRequest:                   "错误的请求",
	StatusUnauthorized:                 "未授权",
	StatusForbidden:                    "拒绝处理",
	StatusNotFound:                     "请求内容不存在",
	StatusMethodNotAllowed:             "不被允许通过的URI",
	StatusNotAcceptable:                "不接受的MIME类型",
	StatusProxyAuthRequired:            "需要代理身份验证",
	StatusRequestTimeout:               "请求超时",
	StatusConflict:                     "请求冲突",
	StatusGone:                         "请求已经不存在",

	StatusInternalServerError:           "服务器错误",
	StatusNotImplemented:                "未实现",
	StatusBadGateway:                    "网关错误",
	StatusServiceUnavailable:            "服务暂时不可用",
	StatusGatewayTimeout:                "网关访问超时",
	StatusHTTPVersionNotSupported:       "HTTP版本不受支持",

	//下面是自定义错误------------------------------------------------------------------------------------------------

	StatusDbError:                       "数据库错误",
	StatusDbConnectionError:             "数据库连接失败",
	StatusDbPingError:                   "数据库Ping失败",
	StatusNotInitDb:                     "没有初始化连接数据库",
	StatusShouldBindError:               "请求数据绑定错误",
	StatusDecodeError:                   "数据库数据绑定错误",
	StatusIsExist:                       "向数据库增加数据重复",
	StatusNotExist:                      "数据库未查到数据",
	StatusLoginError:                    "登录失败",
	StatusCreateTokenError:              "创建Token失败",
	StatusDbFindError:                   "数据库查找错误",
	StatusDbInsertError:                 "数据库增加错误",
	StatusDbUpdateError:                 "数据库修改错误",
	StatusDbDeleteError:                 "数据库删除错误",
	StatusListError:                     "获取列表错误",

	StatusBsonMarshalError:              "bson序列化失败",
	StatusBsonUnmarshalError:            "bson反序列化失败",
	StatusJsonMarshalError:              "json序列化失败",
	StatusJsonUnmarshalError:            "json反序列化失败",
	StatusParseUrlError:                 "url解析失败",
	StatusParseTokenError:               "token解析失败",

	StatusFileError:                     "文件类错误",
	StatusOpenFileError:                 "打开文件错误",
	StatusCloseFileError:                "关闭文件错误",
	StatusReadFileError:				 "读文件错误",
	StatusWriteFileError:	     		 "写文件错误",
	StatusFileNotExist:                  "文件不存在",
}

