package v1

import (
	"errors"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	. "rmonitor/internals/pkg/models/sys"
	. "rmonitor/pkg/common"
	"rmonitor/pkg/jwt"
	"time"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	User
}

func LoginCheck(info LoginInfo) (flag bool, u User, err error) {
	var user User
	if len(info.Username) == 0 || len(info.Password) == 0 {
		return false, user, errors.New("用户名或密码为空")
	}
	user.Name = info.Username
	userFind, err := user.FindOne()
	if err != nil {
		return false, userFind, err
	} else {
		if info.Password == userFind.Password {
			return true, userFind, nil
		} else {
			return false, userFind, errors.New("用户名或密码错误")
		}
	}
}

func Login(c *gin.Context) {
	var login LoginInfo
	err := c.ShouldBindJSON(&login)
	if err == nil {
		isPass, user, _ := LoginCheck(login)
		if isPass {
			generateToken(c, user)
		} else {
			JsonResult(c,608, nil, StatusText[608])
		}
	}else {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	}
}

// 生成令牌
func generateToken(c *gin.Context, user User) {
	j := &jwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := jwt.CustomClaims{
		Username: user.Name,
		//RoleName:   user.RoleName,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		JsonResult(c, StatusCreateTokenError,nil, StatusText[StatusCreateTokenError])
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	JsonResult(c,200,data,"登录成功")
	return
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

func UserList(c *gin.Context) {
	users, err := User{}.List()
	if err != nil {
		JsonResult(c,615,nil, StatusText[615])
	} else {
		JsonResult(c,200, users,"获取列表成功")
	}
}

func UserAddOne(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	} else {
		err = user.AddOne()
		if err != nil {
			if err.Error() == "user is exist"{
				JsonResult(c,606, nil, StatusText[606])
			}else {
				JsonResult(c,612, nil, StatusText[612])
			}
		}else {
			JsonResult(c,204, nil, user.Name+"添加成功")
		}
	}
}

type UserUpdate struct {
	Before User `json:"before"`
	After  User `json:"after"`
}

func UserUpdateOneFromName(c *gin.Context) {
	var update = UserUpdate{}
	if err := c.ShouldBindJSON(&update); err != nil {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	} else {
		if err = update.Before.UpdateOneFromName(update.After); err != nil {
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

func UserDeleteFromName(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		JsonResult(c, StatusShouldBindError, nil, StatusText[StatusShouldBindError])
	} else {
		if err = user.DeleteFromName(); err != nil {
			JsonResult(c,614, nil, StatusText[614])
		} else {
			fmt.Println("user=",user)
			JsonResult(c,204, nil, user.Name+"删除成功")
		}
	}
}
