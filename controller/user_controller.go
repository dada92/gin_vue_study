package controller

import (
	"gin_vue_study/common"
	"gin_vue_study/model"
	"gin_vue_study/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	// 数据输入
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码必须大于6位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 打印结果
	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "用户已经存在",
		})
		return
	}

	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Passwd:    password,
	}
	db.Create(&newUser)

	// 结果返回
	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
