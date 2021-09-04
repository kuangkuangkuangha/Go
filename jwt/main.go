package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var key = "cyw"

type jwtClaim struct {
	jwt.StandardClaims
	Uid int `json:"uid"`
}

type student struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
}

func creatToken(uid int) (string, error) {
	claim := jwtClaim{
		Uid: uid,
	}

	claim.IssuedAt = time.Now().Unix()

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := unsignedToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func login(c *gin.Context) {
	mike := student{}

	if err := c.ShouldBindJSON(&mike); err != nil {
		c.JSON(404, gin.H{
			"mes": "数据绑定失败",
		})
		return
	}

	if token, err := creatToken(mike.Uid); err != nil {
		c.JSON(404, gin.H{
			"mes": "生成Token失败",
		})
	} else {
		c.JSON(404, gin.H{
			"mes": token,
		})
	}
}

func verity(c *gin.Context) {

	token := c.Request.Header.Get("token")
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		c.JSON(404, gin.H{
			"meg": "token解析失败",
		})
	}

	claims, ok := TempToken.Claims.(*jwtClaim)
	if !ok {
		c.JSON(404, gin.H{
			"meg": "token失败",
		})
	}
	if err := TempToken.Claims.Valid(); err != nil {
		c.JSON(404, gin.H{
			"meg": "失败",
		})
	}

	c.JSON(404, gin.H{
		"meg": claims.Uid,
	})
}

func main() {
	r := gin.Default()

	r.POST("/jwt", login)
	r.POST("/ver", verity)

	r.Run()
}
