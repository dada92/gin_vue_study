package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Usre struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(20);not null;unique"`
	Passwd    string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
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
			name = RandomString(10)
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
		newUser := Usre{
			Name:      name,
			Telephone: telephone,
			Passwd:    password,
		}
		db.Create(&newUser)

		// 结果返回
		ctx.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	})
	panic(r.Run(":8080"))
	return
}

func RandomString(n int) string {
	var letter = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIODFGHJKXCVBNM")
	ret := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range ret {
		ret[i] = letter[rand.Intn(len(letter))]
	}
	return string(ret)
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user Usre
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func InitDB() *gorm.DB {
	const (
		user    = "root"
		passwd  = "123"
		host    = "localhost"
		post    = "3306"
		dbname  = "ginvue"
		charset = "utf8mb4"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		passwd,
		host,
		post,
		dbname,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&Usre{})
	return db
}
