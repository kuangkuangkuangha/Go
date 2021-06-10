package model

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id"`
	Role  int   `json:"role"`
}


var (
	key        = "miniProject" //salt，密钥
	ExpireTime = 604800        //token expire time
)

//生成Token
//首先创建一个自定义的clams结构体，里面放自己需要的字段
//然后通过一个gentoken函数得到一个具有签名的token（signtoken）
	//gentoken函数中的 NewWithClaims 是对 claim 结构体指定加密方式和加密对象， 然后用 SignedString 对结构体用事先声明好的 key 对结构体进行签名
	//返回一个完整的token







func CreateToken(role int) string {
	//生成token的第二部分
	//token := jwt.New(jwt.SigningMethodHS256)
	claims := &jwtClaims{
		
		Role : role,
	}
	//token.Claims = claims
	//打印一遍
	//fmt.Println(claims)
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {


	//token有三部分，这是它的前两部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	//这是token的第三部分，进行签名
	//将token 用已经定义好的 key 进行签名，是为了防止数据被篡改
	// 不同的后端设置的密钥（key）不同，生成的token也不同
	//解析后看生成的签名是否与原来的签名一样
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}


//解析Token
//首先用 ParseWithClaims 对 token 进行解析， 得到的是一个 *jwt.Token
//然后用  TempToken.Claims.(*jwtClaims) 得到一个 claim 的实例结构体对象
//claim 中就有创建结构体中的 role 
//然后就可以 通过 role 进行权限管理 


func VerifyToken(token string) (int, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return 0, errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return 0, errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return 0, errors.New("发生错误")
	}
	return claims.Role, nil

}
