package v1

import (
	"github.com/gin-gonic/gin"
	. "rmonitor/internals/pkg/models/sys"
	. "rmonitor/pkg/common"
)

func HostList(c *gin.Context) {
	roles, err := Host{}.List()
	if err != nil {
		JsonResult(c,615,nil,StatusText[615])
	} else {
		JsonResult(c,200, roles,"获取列表成功")
	}
}
