package v1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	. "rmonitor/internals/pkg/models/sys"
	. "rmonitor/pkg/common"
)

func RoleList(c *gin.Context) {
	roles, err := Role{}.List()
	if err != nil {
		JsonResult(c,615,nil,StatusText[615])
	} else {
		JsonResult(c,200, roles,"获取列表成功")
	}
}

func RoleAddOne(c *gin.Context) {
	var role = Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	} else {
		err = role.AddOne()
		if err != nil {
			if err.Error() == "role is exist"{
				JsonResult(c,606, nil, StatusText[606])
			}else{
				JsonResult(c,612, nil, StatusText[612])
			}
		}else {
			JsonResult(c,204, nil, role.Name+"添加成功")
		}
	}
}

type RoleUpdate struct {
	Before Role `json:"before"`
	After  Role `json:"after"`
}

func RoleUpdateOneFromName(c *gin.Context) {
	var update = RoleUpdate{}
	if err := c.ShouldBindJSON(&update); err != nil {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	} else {
		if err := update.Before.UpdateOneFromName(update.After); err != nil {
			if err == mongo.ErrNoDocuments {
				JsonResult(c,607, nil, StatusText[607])
			} else {
				JsonResult(c,613, nil, StatusText[613])
			}
		} else {
			JsonResult(c,204, nil, update.Before.Name+"修改成功")
		}
	}
}

func RoleDeleteFromName(c *gin.Context) {
	var role = Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	} else {
		if err = role.DeleteFromName(); err != nil {
			JsonResult(c,614, nil, StatusText[614])
		} else {
			JsonResult(c,204, nil, role.Name+"删除成功")
		}
	}
}