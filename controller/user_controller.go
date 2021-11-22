package controller

import (
	"gin_vue_study/common"
	"gin_vue_study/model"
	"gin_vue_study/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "加密错误",
		})
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Passwd:    string(hashPassword),
	}
	db.Create(&newUser)

	// 结果返回
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "注册成功",
	})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	// 获取数据
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

	// 判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "电话不存在",
		})
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	// 发放token
	token := "11"

	// 注册成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{
			"token": token,
		},
		"message": "登录成功",
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
